package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/raire/raire/internal/api"
	"github.com/raire/raire/internal/app"
	"github.com/raire/raire/internal/platform/config"
	"github.com/raire/raire/internal/platform/security"
)

type testEnv struct {
	t      *testing.T
	app    *app.App
	server *httptest.Server
	client *http.Client
	logs   *bytes.Buffer
}

type apiResponse struct {
	Status int
	Header http.Header
	Body   []byte
}

func newTestEnv(t *testing.T, mutate ...func(*config.Config)) *testEnv {
	t.Helper()
	cfg := config.Config{
		Env:                "test",
		DBPath:             filepath.Join(t.TempDir(), "raire-test.db"),
		SecureCookies:      false,
		DevMode:            true,
		SessionTTL:         time.Hour,
		ConfirmTokenTTL:    time.Hour,
		ResetTokenTTL:      time.Hour,
		TermsVersion:       "2026-11-20",
		PrivacyVersion:     "2026-11-20",
		EncryptionKey:      "MDEyMzQ1Njc4OWFiY2RlZjAxMjM0NTY3ODlhYmNkZWY=",
		MaxBodyBytes:       1 << 20,
		RateLimitEnabled:   false,
		PreciseLocationTTL: 90 * 24 * time.Hour,
		Argon2:             config.Argon2{Memory: 8 * 1024, Iterations: 1, Parallelism: 1, SaltLength: 16, KeyLength: 32},
	}
	for _, m := range mutate {
		m(&cfg)
	}
	a, err := app.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = a.Close() })

	logs := &bytes.Buffer{}
	logger := slog.New(slog.NewJSONHandler(logs, nil))
	srv := httptest.NewServer(api.NewRouter(a, logger))
	t.Cleanup(srv.Close)

	jar, _ := cookiejar.New(nil)
	return &testEnv{
		t: t, app: a, server: srv, logs: logs,
		client: &http.Client{Jar: jar, Timeout: 10 * time.Second},
	}
}

func (e *testEnv) do(method, path string, body any, headers map[string]string) apiResponse {
	return e.doWith(e.client, method, path, body, headers)
}

func (e *testEnv) doWith(client *http.Client, method, path string, body any, headers map[string]string) apiResponse {
	e.t.Helper()
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			e.t.Fatal(err)
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequest(method, e.server.URL+path, reader)
	if err != nil {
		e.t.Fatal(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		e.t.Fatal(err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	return apiResponse{Status: resp.StatusCode, Header: resp.Header, Body: raw}
}

func (e *testEnv) signup(email, password string) string {
	e.t.Helper()
	res := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": email, "password": password, "display_name": "Ana Souza",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if res.Status != http.StatusAccepted {
		e.t.Fatalf("signup status=%d body=%s", res.Status, res.Body)
	}
	token := res.Header.Get("X-Dev-Confirm-Token")
	if token == "" {
		e.t.Fatal("missing dev confirm token")
	}
	return token
}

func (e *testEnv) confirm(token string) apiResponse {
	e.t.Helper()
	return e.do("POST", "/v1/auth/confirm-email", map[string]any{"token": token}, nil)
}

func (e *testEnv) signupConfirmed(email, password string) {
	e.t.Helper()
	token := e.signup(email, password)
	res := e.confirm(token)
	if res.Status != http.StatusNoContent {
		e.t.Fatalf("confirm status=%d body=%s", res.Status, res.Body)
	}
}

func (e *testEnv) login(email, password, mfa string) (string, apiResponse) {
	return e.loginWith(e.client, email, password, mfa)
}

func (e *testEnv) loginWith(client *http.Client, email, password, mfa string) (string, apiResponse) {
	e.t.Helper()
	res := e.doWith(client, "POST", "/v1/auth/login", map[string]any{
		"email": email, "password": password, "mfa_code": mfa,
	}, nil)
	var out struct {
		CSRFToken string `json:"csrf_token"`
	}
	_ = json.Unmarshal(res.Body, &out)
	return out.CSRFToken, res
}

func (e *testEnv) loginOK(email, password string) string {
	e.t.Helper()
	csrf, res := e.login(email, password, "")
	if res.Status != http.StatusOK {
		e.t.Fatalf("login status=%d body=%s", res.Status, res.Body)
	}
	if csrf == "" {
		e.t.Fatal("missing csrf token")
	}
	return csrf
}

func authHeaders(csrf string) map[string]string {
	return map[string]string{"X-CSRF-Token": csrf}
}

func errorCode(t *testing.T, body []byte) string {
	t.Helper()
	var out struct {
		Error struct {
			Code string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("decode error body %s: %v", body, err)
	}
	return out.Error.Code
}

func decode[T any](t *testing.T, body []byte) T {
	t.Helper()
	var out T
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("decode body %s: %v", body, err)
	}
	return out
}

func (e *testEnv) operator(email, password, role string) (csrf, secret string, recovery []string) {
	return e.operatorWith(e.client, email, password, role)
}

func (e *testEnv) operatorWith(client *http.Client, email, password, role string) (csrf, secret string, recovery []string) {
	e.t.Helper()
	acct, secret, codes, err := e.app.CreateOperator(context.Background(), email, password)
	if err != nil {
		e.t.Fatal(err)
	}
	if role == "admin" {
		if _, err := e.app.DB.Exec(`UPDATE accounts SET role = 'admin' WHERE id = ?`, acct.ID); err != nil {
			e.t.Fatal(err)
		}
	}
	code, err := security.GenerateTOTP(secret, time.Now())
	if err != nil {
		e.t.Fatal(err)
	}
	csrf, res := e.loginWith(client, email, password, code)
	if res.Status != http.StatusOK {
		e.t.Fatalf("operator login status=%d body=%s", res.Status, res.Body)
	}
	return csrf, secret, codes
}

func (e *testEnv) publishProfile(client *http.Client, csrf string) string {
	e.t.Helper()
	body := map[string]any{
		"display_name": "Ana Souza",
		"bio":          "Eletricista com 10 anos de experiência.",
		"service_type": "in_person",
		"category_ids": []string{"cat_eletricista"},
		"city":         "Cidade Piloto",
		"neighborhood": "Centro",
		"radius_km":    10,
		"availability": []map[string]any{
			{"day": 1, "start": "08:00", "end": "12:00", "timezone": "America/Sao_Paulo"},
		},
	}
	res := e.doWith(client, "POST", "/v1/professionals/me", body, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		e.t.Fatalf("create profile status=%d body=%s", res.Status, res.Body)
	}
	created := decode[struct {
		Profile struct {
			ID string `json:"id"`
		} `json:"profile"`
	}](e.t, res.Body)
	res = e.doWith(client, "POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusOK {
		e.t.Fatalf("publish status=%d body=%s", res.Status, res.Body)
	}
	if created.Profile.ID == "" {
		e.t.Fatal("missing profile id")
	}
	return created.Profile.ID
}
