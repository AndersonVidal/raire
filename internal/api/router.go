package api

import (
	"bytes"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/raire/raire/internal/app"
	catalogapp "github.com/raire/raire/internal/catalog/app"
	"github.com/raire/raire/internal/contracts"
	directoryapp "github.com/raire/raire/internal/directory/app"
	identityapp "github.com/raire/raire/internal/identity/app"
	"github.com/raire/raire/internal/identity/domain"
	"github.com/raire/raire/internal/platform/audit"
	"github.com/raire/raire/internal/platform/httpx"
	"github.com/raire/raire/internal/platform/security"
)

const (
	sessionCookie = "raire_session"
	csrfCookie    = "raire_csrf"
)

type Handler struct {
	app    *app.App
	logger *slog.Logger
}

func NewRouter(a *app.App, logger *slog.Logger) http.Handler {
	h := &Handler{app: a, logger: logger}
	mux := http.NewServeMux()
	h.routes(mux)
	return httpx.RequestCorrelation(
		httpx.SecurityHeaders(
			httpx.Logging(logger)(
				httpx.Recover(logger)(mux),
			),
		),
	)
}

type authContext struct {
	account *domain.Account
	session *domain.Session
	token   string
}

func (h *Handler) routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/healthz", h.health)
	mux.HandleFunc("GET /v1/readyz", h.ready)

	mux.HandleFunc("POST /v1/auth/signup", h.signup)
	mux.HandleFunc("POST /v1/auth/login", h.login)
	mux.HandleFunc("POST /v1/auth/logout", h.protected(true, "", h.logout))
	mux.HandleFunc("POST /v1/auth/confirm-email", h.confirmEmail)
	mux.HandleFunc("POST /v1/auth/resend-confirmation", h.resendConfirmation)
	mux.HandleFunc("POST /v1/auth/password-reset/request", h.passwordResetRequest)
	mux.HandleFunc("POST /v1/auth/password-reset/confirm", h.passwordResetConfirm)

	mux.HandleFunc("GET /v1/me", h.protected(false, "", h.me))
	mux.HandleFunc("PATCH /v1/me", h.protected(true, "", h.updateMe))
	mux.HandleFunc("POST /v1/me/password", h.protected(true, "", h.changePassword))
	mux.HandleFunc("GET /v1/me/consents", h.protected(false, "", h.listConsents))
	mux.HandleFunc("POST /v1/me/consents/{kind}/revoke", h.protected(true, "", h.revokeConsent))

	mux.HandleFunc("GET /v1/categories", h.listCategories)
	mux.HandleFunc("GET /v1/categories/{slug}", h.getCategory)

	mux.HandleFunc("GET /v1/professionals/me", h.protected(false, "", h.ownerProfile))
	mux.HandleFunc("POST /v1/professionals/me", h.protected(true, "", h.createProfile))
	mux.HandleFunc("PATCH /v1/professionals/me", h.protected(true, "", h.updateProfile))
	mux.HandleFunc("POST /v1/professionals/me/publish", h.protected(true, "", h.publishProfile))
	mux.HandleFunc("POST /v1/professionals/me/pause", h.protected(true, "", h.pauseProfile))
	mux.HandleFunc("POST /v1/professionals/me/resume", h.protected(true, "", h.resumeProfile))
	mux.HandleFunc("GET /v1/professionals/{id}", h.publicProfile)
	mux.HandleFunc("GET /v1/professionals/{id}/preview", h.protected(false, "", h.previewProfile))

	mux.HandleFunc("GET /v1/admin/categories", h.protected(false, "operator", h.adminListCategories))
	mux.HandleFunc("POST /v1/admin/categories", h.protected(true, "operator", h.adminCreateCategory))
	mux.HandleFunc("PATCH /v1/admin/categories/{id}", h.protected(true, "operator", h.adminUpdateCategory))
	mux.HandleFunc("POST /v1/admin/categories/{id}/deactivate", h.protected(true, "operator", h.adminDeactivateCategory))

	mux.HandleFunc("GET /v1/admin/professionals", h.protected(false, "operator", h.adminListProfiles))
	mux.HandleFunc("POST /v1/admin/professionals/{id}/suspend", h.protected(true, "operator", h.adminSuspendProfile))
	mux.HandleFunc("POST /v1/admin/professionals/{id}/reactivate", h.protected(true, "operator", h.adminReactivateProfile))
	mux.HandleFunc("GET /v1/admin/audit-events", h.protected(false, "operator", h.adminAuditEvents))

	mux.HandleFunc("GET /v1/admin/feature-flags", h.protected(false, "admin", h.adminListFlags))
	mux.HandleFunc("PATCH /v1/admin/feature-flags/{key}", h.protected(true, "admin", h.adminUpdateFlag))
}

func (h *Handler) protected(mutating bool, minRole string, fn func(http.ResponseWriter, *http.Request, *authContext)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := cookieValue(r, sessionCookie)
		auth, err := h.app.Identity.Authenticate(r.Context(), token)
		if err != nil {
			h.fail(w, r, err)
			return
		}
		if mutating {
			header := strings.TrimSpace(r.Header.Get("X-CSRF-Token"))
			if header == "" || subtle.ConstantTimeCompare([]byte(security.HashToken(header)), []byte(auth.Session.CSRFTokenHash)) != 1 {
				h.fail(w, r, contracts.New(contracts.CodeCSRFInvalid, ""))
				return
			}
		}
		switch minRole {
		case "operator":
			if !auth.Account.Privileged() {
				h.fail(w, r, contracts.New(contracts.CodeForbidden, ""))
				return
			}
		case "admin":
			if auth.Account.Role != domain.RoleAdmin {
				h.fail(w, r, contracts.New(contracts.CodeForbidden, ""))
				return
			}
		}
		fn(w, r, &authContext{account: auth.Account, session: auth.Session, token: token})
	}
}

func (h *Handler) fail(w http.ResponseWriter, r *http.Request, err error) {
	appErr := contracts.As(err)
	if appErr.HTTPStatus >= 500 {
		h.logger.Error("request_failed",
			"correlation_id", httpx.CorrelationID(r.Context()),
			"code", string(appErr.Code),
			"path", r.URL.Path,
		)
	}
	httpx.WriteError(w, r, appErr)
}

func (h *Handler) decode(w http.ResponseWriter, r *http.Request, dst any) ([]byte, error) {
	max := h.app.Cfg.MaxBodyBytes
	if max <= 0 {
		max = 1 << 20
	}
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, max))
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			return nil, contracts.Validation("O conteúdo enviado é grande demais.")
		}
		return nil, contracts.Validation("Não foi possível ler o conteúdo enviado.")
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return body, contracts.Validation("Confira os campos destacados.")
	}
	if dec.More() {
		return body, contracts.Validation("Envie apenas um objeto JSON.")
	}
	return body, nil
}

func (h *Handler) meta(r *http.Request, body []byte) identityapp.RequestMeta {
	return identityapp.RequestMeta{
		CorrelationID:  httpx.CorrelationID(r.Context()),
		IP:             httpx.ClientIP(r, h.app.Cfg.TrustProxy),
		UserAgent:      r.UserAgent(),
		IdempotencyKey: strings.TrimSpace(r.Header.Get("Idempotency-Key")),
		RequestHash:    identityapp.RequestHash(body),
	}
}

func (h *Handler) limit(w http.ResponseWriter, r *http.Request, key string, limit int, window time.Duration) bool {
	if !h.app.Cfg.RateLimitEnabled {
		return true
	}
	allowed, remaining, retry := h.app.Limiter.Allow(key, limit, window)
	w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
	w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(max(remaining, 0)))
	if !allowed {
		seconds := int(retry.Seconds())
		if seconds < 1 {
			seconds = 1
		}
		w.Header().Set("Retry-After", strconv.Itoa(seconds))
		h.app.Metrics.Inc("rate_limited")
		h.fail(w, r, contracts.New(contracts.CodeRateLimited, ""))
		return false
	}
	return true
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ready(w http.ResponseWriter, r *http.Request) {
	if err := h.app.Ready(r.Context()); err != nil {
		httpx.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "unavailable"})
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	if !h.limit(w, r, "signup:"+httpx.ClientIP(r, h.app.Cfg.TrustProxy), 3, time.Hour) {
		return
	}
	var req struct {
		Email         string `json:"email"`
		Password      string `json:"password"`
		DisplayName   string `json:"display_name"`
		AcceptTerms   bool   `json:"accept_terms"`
		AcceptPrivacy bool   `json:"accept_privacy"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	res, err := h.app.Identity.Signup(r.Context(), identityapp.SignupInput{
		Email: req.Email, Password: req.Password, DisplayName: req.DisplayName,
		AcceptTerms: req.AcceptTerms, AcceptPrivacy: req.AcceptPrivacy,
	}, h.meta(r, body))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	if h.app.Cfg.DevMode && res.Created && res.ConfirmToken != "" {
		w.Header().Set("X-Dev-Confirm-Token", res.ConfirmToken)
	}
	// Deliberately identical for new and existing e-mails: the response must not
	// be an account-enumeration oracle.
	httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"status": "ok"})
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		MFACode  string `json:"mfa_code"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	ip := httpx.ClientIP(r, h.app.Cfg.TrustProxy)
	if !h.limit(w, r, "login:ip:"+ip, 5, time.Minute) {
		return
	}
	emailKey := domain.NormalizeEmail(req.Email)
	if !h.limit(w, r, "login:account:"+emailKey, 5, time.Minute) {
		return
	}
	res, err := h.app.Identity.Login(r.Context(), identityapp.LoginInput{
		Email: req.Email, Password: req.Password, MFACode: req.MFACode,
		ExistingToken: cookieValue(r, sessionCookie),
	}, h.meta(r, body))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	h.setSessionCookies(w, res.SessionToken, res.CSRFToken, res.ExpiresAt)
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"account":    accountDTO(res.Account),
		"csrf_token": res.CSRFToken,
	})
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if err := h.app.Identity.Logout(r.Context(), auth.token); err != nil {
		h.fail(w, r, err)
		return
	}
	h.clearSessionCookies(w)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) confirmEmail(w http.ResponseWriter, r *http.Request) {
	if !h.limit(w, r, "confirm:"+httpx.ClientIP(r, h.app.Cfg.TrustProxy), 5, time.Hour) {
		return
	}
	var req struct {
		Token string `json:"token"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	if err := h.app.Identity.ConfirmEmail(r.Context(), req.Token, h.meta(r, body)); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) resendConfirmation(w http.ResponseWriter, r *http.Request) {
	if !h.limit(w, r, "resend:"+httpx.ClientIP(r, h.app.Cfg.TrustProxy), 5, time.Hour) {
		return
	}
	var req struct {
		Email string `json:"email"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	token, _ := h.app.Identity.ResendConfirmation(r.Context(), req.Email, h.meta(r, body))
	if h.app.Cfg.DevMode && token != "" {
		w.Header().Set("X-Dev-Confirm-Token", token)
	}
	httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"status": "ok"})
}

func (h *Handler) passwordResetRequest(w http.ResponseWriter, r *http.Request) {
	if !h.limit(w, r, "reset:"+httpx.ClientIP(r, h.app.Cfg.TrustProxy), 3, time.Hour) {
		return
	}
	var req struct {
		Email string `json:"email"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	token, _ := h.app.Identity.RequestPasswordReset(r.Context(), req.Email, h.meta(r, body))
	if h.app.Cfg.DevMode && token != "" {
		w.Header().Set("X-Dev-Reset-Token", token)
	}
	httpx.WriteJSON(w, http.StatusAccepted, map[string]string{"status": "ok"})
}

func (h *Handler) passwordResetConfirm(w http.ResponseWriter, r *http.Request) {
	if !h.limit(w, r, "reset-confirm:"+httpx.ClientIP(r, h.app.Cfg.TrustProxy), 5, time.Hour) {
		return
	}
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	if err := h.app.Identity.ConfirmPasswordReset(r.Context(), req.Token, req.NewPassword, h.meta(r, body)); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request, auth *authContext) {
	httpx.WriteJSON(w, http.StatusOK, map[string]any{
		"account":    accountDTO(auth.account),
		"csrf_token": cookieValue(r, csrfCookie),
	})
}

func (h *Handler) updateMe(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		DisplayName       *string `json:"display_name"`
		ContactPhone      *string `json:"contact_phone"`
		ContactPreference *string `json:"contact_preference"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	acct, err := h.app.Identity.UpdateMe(r.Context(), auth.account, identityapp.UpdateMeInput{
		DisplayName: req.DisplayName, ContactPhone: req.ContactPhone, ContactPreference: req.ContactPreference,
	}, h.meta(r, body))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"account": accountDTO(acct)})
}

func (h *Handler) changePassword(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
	body, err := h.decode(w, r, &req)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	if err := h.app.Identity.ChangePassword(r.Context(), auth.account, auth.session.ID, req.CurrentPassword, req.NewPassword, h.meta(r, body)); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) listConsents(w http.ResponseWriter, r *http.Request, auth *authContext) {
	consents, err := h.app.Identity.ListConsents(r.Context(), auth.account.ID)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	items := make([]consentDTO, 0, len(consents))
	for _, c := range consents {
		items = append(items, consentDTO{
			ID: c.ID, Kind: string(c.Kind), Version: c.Version, GrantedAt: c.GrantedAt, RevokedAt: c.RevokedAt,
		})
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
}

func (h *Handler) revokeConsent(w http.ResponseWriter, r *http.Request, auth *authContext) {
	kind := r.PathValue("kind")
	body := []byte("{}")
	if r.ContentLength > 0 {
		var err error
		body, err = h.decode(w, r, &struct{}{})
		if err != nil {
			h.fail(w, r, err)
			return
		}
	}
	if err := h.app.Identity.RevokeConsent(r.Context(), auth.account, kind, h.meta(r, body)); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) listCategories(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("q")
	categories, err := h.app.Catalog.Search(r.Context(), term)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	items := make([]categoryDTO, 0, len(categories))
	for _, c := range categories {
		items = append(items, categoryDTOFrom(c))
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
}

func (h *Handler) getCategory(w http.ResponseWriter, r *http.Request) {
	c, err := h.app.Catalog.GetBySlug(r.Context(), r.PathValue("slug"))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	if !c.Active {
		h.fail(w, r, contracts.New(contracts.CodeNotFound, ""))
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"category": categoryDTOFrom(*c)})
}

func (h *Handler) ownerProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	view, err := h.app.Directory.OwnerView(r.Context(), auth.account.ID)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"profile": ownerViewDTO(view)})
}

func (h *Handler) createProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req profileWriteRequest
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	in := directoryapp.CreateInput{
		DisplayName: value(req.DisplayName), Bio: value(req.Bio), ServiceType: value(req.ServiceType),
		City: value(req.City), Neighborhood: value(req.Neighborhood), RadiusKm: req.RadiusKm,
		Availability: windowsFromDTO(req.Availability), AvailabilityNote: value(req.AvailabilityNote),
		ContactPreference: req.ContactPreference, PreciseLat: req.PreciseLat, PreciseLng: req.PreciseLng,
	}
	if req.CategoryIDs != nil {
		in.CategoryIDs = *req.CategoryIDs
	}
	profile, err := h.app.Directory.CreateDraft(r.Context(), auth.account.ID, in, httpx.CorrelationID(r.Context()))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	view, err := h.app.Directory.OwnerView(r.Context(), auth.account.ID)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, map[string]any{"profile": ownerViewDTO(view), "id": profile.ID})
}

func (h *Handler) updateProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req profileWriteRequest
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	in := directoryapp.UpdateInput{
		DisplayName: req.DisplayName, Bio: req.Bio, ServiceType: req.ServiceType,
		CategoryIDs: req.CategoryIDs, City: req.City, Neighborhood: req.Neighborhood,
		RadiusKm: req.RadiusKm, ClearRadius: req.ClearRadius,
		AvailabilityNote: req.AvailabilityNote, ContactPreference: req.ContactPreference,
		PreciseLat: req.PreciseLat, PreciseLng: req.PreciseLng, ClearPreciseLocation: req.ClearPreciseLocation,
	}
	if req.Availability != nil {
		windows := windowsFromDTO(req.Availability)
		in.Availability = &windows
	}
	if _, err := h.app.Directory.UpdateOwner(r.Context(), auth.account.ID, in, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	view, err := h.app.Directory.OwnerView(r.Context(), auth.account.ID)
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"profile": ownerViewDTO(view)})
}

func (h *Handler) publishProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if _, err := h.app.Directory.Publish(r.Context(), auth.account.ID, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	h.ownerProfile(w, r, auth)
}

func (h *Handler) pauseProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if _, err := h.app.Directory.Pause(r.Context(), auth.account.ID, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	h.ownerProfile(w, r, auth)
}

func (h *Handler) resumeProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if _, err := h.app.Directory.Resume(r.Context(), auth.account.ID, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	h.ownerProfile(w, r, auth)
}

func (h *Handler) publicProfile(w http.ResponseWriter, r *http.Request) {
	view, err := h.app.Directory.PublicView(r.Context(), r.PathValue("id"))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"profile": publicViewDTO(view)})
}

func (h *Handler) previewProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	view, err := h.app.Directory.PreviewByID(r.Context(), auth.account.ID, r.PathValue("id"))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"profile": publicViewDTO(view)})
}

func (h *Handler) adminListCategories(w http.ResponseWriter, r *http.Request, _ *authContext) {
	categories, err := h.app.Catalog.ListAll(r.Context())
	if err != nil {
		h.fail(w, r, err)
		return
	}
	items := make([]categoryDTO, 0, len(categories))
	for _, c := range categories {
		items = append(items, categoryDTOFrom(c))
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
}

func (h *Handler) adminCreateCategory(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		Slug        string   `json:"slug"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Synonyms    []string `json:"synonyms"`
		SortOrder   int      `json:"sort_order"`
	}
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	c, err := h.app.Catalog.Create(r.Context(), catalogapp.Actor{ID: auth.account.ID, Role: string(auth.account.Role)},
		catalogapp.CreateInput{Slug: req.Slug, Name: req.Name, Description: req.Description, Synonyms: req.Synonyms, SortOrder: req.SortOrder},
		httpx.CorrelationID(r.Context()))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, map[string]any{"category": categoryDTOFrom(*c)})
}

func (h *Handler) adminUpdateCategory(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		Name        *string   `json:"name"`
		Description *string   `json:"description"`
		Synonyms    *[]string `json:"synonyms"`
		SortOrder   *int      `json:"sort_order"`
	}
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	c, err := h.app.Catalog.Update(r.Context(), catalogapp.Actor{ID: auth.account.ID, Role: string(auth.account.Role)},
		r.PathValue("id"), catalogapp.UpdateInput{Name: req.Name, Description: req.Description, Synonyms: req.Synonyms, SortOrder: req.SortOrder},
		httpx.CorrelationID(r.Context()))
	if err != nil {
		h.fail(w, r, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"category": categoryDTOFrom(*c)})
}

func (h *Handler) adminDeactivateCategory(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if err := h.app.Catalog.Deactivate(r.Context(), catalogapp.Actor{ID: auth.account.ID, Role: string(auth.account.Role)},
		r.PathValue("id"), httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) adminListProfiles(w http.ResponseWriter, r *http.Request, _ *authContext) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, err := h.app.Directory.ListAdmin(r.Context(), directoryapp.ListFilter{
		Status: r.URL.Query().Get("status"), Query: r.URL.Query().Get("q"),
		Cursor: r.URL.Query().Get("cursor"), Limit: limit,
	})
	if err != nil {
		h.fail(w, r, contracts.Validation("Filtro inválido."))
		return
	}
	items := make([]adminProfileDTO, 0, len(page.Items))
	for _, p := range page.Items {
		items = append(items, adminProfileDTO{
			ID: p.ID, AccountID: p.AccountID, DisplayName: p.DisplayName, Status: p.Status,
			City: p.City, Neighborhood: p.Neighborhood, UpdatedAt: p.UpdatedAt,
		})
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items, "next_cursor": page.NextCursor})
}

func (h *Handler) adminSuspendProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		Reason   string `json:"reason"`
		Evidence string `json:"evidence"`
	}
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	if err := h.app.Directory.SuspendAndRevoke(r.Context(), auth.account.ID, string(auth.account.Role),
		r.PathValue("id"), req.Reason, req.Evidence, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) adminReactivateProfile(w http.ResponseWriter, r *http.Request, auth *authContext) {
	if err := h.app.Directory.Reactivate(r.Context(), auth.account.ID, string(auth.account.Role),
		r.PathValue("id"), httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) adminAuditEvents(w http.ResponseWriter, r *http.Request, _ *authContext) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	filter := auditFilterFrom(r)
	filter.Limit = limit
	page, err := h.app.ListAudit(r.Context(), filter)
	if err != nil {
		h.fail(w, r, contracts.Validation("Filtro inválido."))
		return
	}
	items := make([]auditDTO, 0, len(page.Items))
	for _, e := range page.Items {
		items = append(items, auditDTO{
			ID: e.ID, ActorID: e.ActorID, ActorRole: e.ActorRole, Action: e.Action,
			TargetType: e.TargetType, TargetID: e.TargetID, Reason: e.Reason,
			CorrelationID: e.CorrelationID, Metadata: e.Metadata, CreatedAt: e.CreatedAt,
		})
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items, "next_cursor": page.NextCursor})
}

func (h *Handler) adminListFlags(w http.ResponseWriter, r *http.Request, _ *authContext) {
	rows, err := h.app.ListFlags(r.Context())
	if err != nil {
		h.fail(w, r, err)
		return
	}
	items := make([]flagDTO, 0, len(rows))
	for _, f := range rows {
		items = append(items, flagDTO{Key: f.Key, Enabled: f.Enabled, Description: f.Description,
			Owner: f.Owner, ReviewAt: f.ReviewAt, UpdatedAt: f.UpdatedAt})
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
}

func (h *Handler) adminUpdateFlag(w http.ResponseWriter, r *http.Request, auth *authContext) {
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if _, err := h.decode(w, r, &req); err != nil {
		h.fail(w, r, err)
		return
	}
	if err := h.app.SetFeatureFlag(r.Context(), auth.account.ID, string(auth.account.Role),
		r.PathValue("key"), req.Enabled, httpx.CorrelationID(r.Context())); err != nil {
		h.fail(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) setSessionCookies(w http.ResponseWriter, token, csrf string, expires time.Time) {
	secure := h.app.Cfg.SecureCookies
	http.SetCookie(w, &http.Cookie{
		Name: sessionCookie, Value: token, Path: "/", HttpOnly: true, Secure: secure,
		SameSite: http.SameSiteLaxMode, Expires: expires, MaxAge: int(time.Until(expires).Seconds()),
	})
	http.SetCookie(w, &http.Cookie{
		Name: csrfCookie, Value: csrf, Path: "/", HttpOnly: false, Secure: secure,
		SameSite: http.SameSiteLaxMode, Expires: expires, MaxAge: int(time.Until(expires).Seconds()),
	})
}

func (h *Handler) clearSessionCookies(w http.ResponseWriter) {
	for _, name := range []string{sessionCookie, csrfCookie} {
		http.SetCookie(w, &http.Cookie{Name: name, Value: "", Path: "/", MaxAge: -1, HttpOnly: name == sessionCookie,
			Secure: h.app.Cfg.SecureCookies, SameSite: http.SameSiteLaxMode})
	}
}

func cookieValue(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func value(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func auditFilterFrom(r *http.Request) audit.Filter {
	q := r.URL.Query()
	f := audit.Filter{
		ActorID: q.Get("actor"), TargetType: q.Get("target_type"), TargetID: q.Get("target_id"),
		Cursor: q.Get("cursor"),
	}
	if v := q.Get("from"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			f.From = &t
		}
	}
	if v := q.Get("to"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			f.To = &t
		}
	}
	return f
}
