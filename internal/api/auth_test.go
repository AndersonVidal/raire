package api_test

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/raire/raire/internal/platform/security"
)

func TestSignupConfirmLoginLogout(t *testing.T) {
	e := newTestEnv(t)
	token := e.signup("ana@example.test", "senha-segura-123")
	if res := e.confirm(token); res.Status != http.StatusNoContent {
		t.Fatalf("confirm status=%d body=%s", res.Status, res.Body)
	}
	csrf, res := e.login("ana@example.test", "senha-segura-123", "")
	if res.Status != http.StatusOK {
		t.Fatalf("login status=%d body=%s", res.Status, res.Body)
	}
	if csrf == "" {
		t.Fatal("missing csrf token")
	}
	me := e.do("GET", "/v1/me", nil, nil)
	if me.Status != http.StatusOK {
		t.Fatalf("me status=%d body=%s", me.Status, me.Body)
	}
	if body := string(me.Body); strings.Contains(body, "password_hash") || strings.Contains(body, "senha-segura-123") {
		t.Fatalf("me leaked password material: %s", body)
	}
	out := decode[struct {
		Account struct {
			EmailVerified bool `json:"email_verified"`
		} `json:"account"`
	}](t, me.Body)
	if !out.Account.EmailVerified {
		t.Fatal("email_verified should be true")
	}
	if res := e.do("POST", "/v1/auth/logout", nil, authHeaders(csrf)); res.Status != http.StatusNoContent {
		t.Fatalf("logout status=%d body=%s", res.Status, res.Body)
	}
	me = e.do("GET", "/v1/me", nil, nil)
	if me.Status != http.StatusUnauthorized {
		t.Fatalf("me after logout status=%d", me.Status)
	}
	if code := errorCode(t, me.Body); code != "UNAUTHENTICATED" {
		t.Fatalf("error code = %s", code)
	}
}

func TestDuplicateSignupIsGeneric(t *testing.T) {
	e := newTestEnv(t)
	first := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": "ana@example.test", "password": "senha-segura-123", "display_name": "Ana",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if first.Status != http.StatusAccepted {
		t.Fatalf("first status=%d body=%s", first.Status, first.Body)
	}
	if first.Header.Get("X-Dev-Confirm-Token") == "" {
		t.Fatal("new signup did not return a dev confirmation token")
	}
	second := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": "ana@example.test", "password": "outra-senha-456", "display_name": "Ana",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if second.Status != first.Status {
		t.Fatalf("status differs: new=%d duplicate=%d", first.Status, second.Status)
	}
	if string(second.Body) != string(first.Body) {
		t.Fatalf("body differs: new=%s duplicate=%s", first.Body, second.Body)
	}
	if second.Header.Get("X-Dev-Confirm-Token") != "" {
		t.Fatal("duplicate signup returned a confirmation token")
	}
	var count int
	if err := e.app.DB.QueryRow(`SELECT count(*) FROM accounts WHERE email_normalized = 'ana@example.test'`).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("account count = %d", count)
	}
}

func TestConfirmTokenSingleUseAndExpiry(t *testing.T) {
	e := newTestEnv(t)
	token := e.signup("ana@example.test", "senha-segura-123")
	if res := e.confirm(token); res.Status != http.StatusNoContent {
		t.Fatalf("first confirm status=%d body=%s", res.Status, res.Body)
	}
	res := e.confirm(token)
	if res.Status != http.StatusUnprocessableEntity {
		t.Fatalf("reuse status=%d body=%s", res.Status, res.Body)
	}
	if code := errorCode(t, res.Body); code != "TOKEN_INVALID" {
		t.Fatalf("reuse code = %s", code)
	}

	token2 := e.signup("bia@example.test", "senha-segura-123")
	past := time.Now().Add(-time.Hour).UTC().Format(time.RFC3339Nano)
	if _, err := e.app.DB.Exec(`UPDATE email_tokens SET expires_at = ? WHERE token_hash = ?`, past, security.HashToken(token2)); err != nil {
		t.Fatal(err)
	}
	res = e.confirm(token2)
	if res.Status != http.StatusUnprocessableEntity {
		t.Fatalf("expired status=%d body=%s", res.Status, res.Body)
	}
	if code := errorCode(t, res.Body); code != "TOKEN_EXPIRED" {
		t.Fatalf("expired code = %s", code)
	}
}

func TestLoginRotatesSession(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	e.loginOK("ana@example.test", "senha-segura-123")

	u, _ := url.Parse(e.server.URL)
	var oldToken string
	for _, c := range e.client.Jar.Cookies(u) {
		if c.Name == "raire_session" {
			oldToken = c.Value
		}
	}
	if oldToken == "" {
		t.Fatal("missing session cookie")
	}
	e.loginOK("ana@example.test", "senha-segura-123")

	req, _ := http.NewRequest("GET", e.server.URL+"/v1/me", nil)
	req.AddCookie(&http.Cookie{Name: "raire_session", Value: oldToken})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("old session status=%d", resp.StatusCode)
	}
}

func TestPasswordReset(t *testing.T) {
	e := newTestEnv(t)
	res := e.do("POST", "/v1/auth/password-reset/request", map[string]any{"email": "ghost@example.test"}, nil)
	if res.Status != http.StatusAccepted || res.Header.Get("X-Dev-Reset-Token") != "" {
		t.Fatalf("unknown reset status=%d token=%q", res.Status, res.Header.Get("X-Dev-Reset-Token"))
	}
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	res = e.do("POST", "/v1/auth/password-reset/request", map[string]any{"email": "ana@example.test"}, nil)
	if res.Status != http.StatusAccepted {
		t.Fatalf("reset request status=%d body=%s", res.Status, res.Body)
	}
	token := res.Header.Get("X-Dev-Reset-Token")
	if token == "" {
		t.Fatal("missing dev reset token")
	}
	res = e.do("POST", "/v1/auth/password-reset/confirm", map[string]any{"token": token, "new_password": "nova-senha-456"}, nil)
	if res.Status != http.StatusNoContent {
		t.Fatalf("reset confirm status=%d body=%s", res.Status, res.Body)
	}
	_, res = e.login("ana@example.test", "senha-segura-123", "")
	if res.Status != http.StatusUnauthorized {
		t.Fatalf("old password status=%d", res.Status)
	}
	e.loginOK("ana@example.test", "nova-senha-456")
	res = e.do("POST", "/v1/auth/password-reset/confirm", map[string]any{"token": token, "new_password": "outra-senha-789"}, nil)
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "TOKEN_INVALID" {
		t.Fatalf("token reuse status=%d body=%s", res.Status, res.Body)
	}
}

func TestPasswordChangeRevokesOtherSessions(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf1 := e.loginOK("ana@example.test", "senha-segura-123")

	jar2, _ := cookiejar.New(nil)
	client2 := &http.Client{Jar: jar2, Timeout: 10 * time.Second}
	res := e.doWith(client2, "POST", "/v1/auth/login", map[string]any{
		"email": "ana@example.test", "password": "senha-segura-123",
	}, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("second login status=%d body=%s", res.Status, res.Body)
	}

	res = e.doWith(e.client, "POST", "/v1/me/password", map[string]any{
		"current_password": "senha-segura-123", "new_password": "nova-senha-456",
	}, authHeaders(csrf1))
	if res.Status != http.StatusNoContent {
		t.Fatalf("password change status=%d body=%s", res.Status, res.Body)
	}
	if res = e.doWith(client2, "GET", "/v1/me", nil, nil); res.Status != http.StatusUnauthorized {
		t.Fatalf("other session status=%d", res.Status)
	}
	if res = e.doWith(e.client, "GET", "/v1/me", nil, nil); res.Status != http.StatusOK {
		t.Fatalf("current session status=%d", res.Status)
	}
	_, res = e.login("ana@example.test", "senha-segura-123", "")
	if res.Status != http.StatusUnauthorized {
		t.Fatalf("old password status=%d", res.Status)
	}
	e.loginOK("ana@example.test", "nova-senha-456")
}

func TestOperatorRequiresMFAAndRecovery(t *testing.T) {
	e := newTestEnv(t)
	_, secret, codes, err := e.app.CreateOperator(t.Context(), "ops@example.test", "senha-segura-123")
	if err != nil {
		t.Fatal(err)
	}
	var stored string
	if err := e.app.DB.QueryRow(`SELECT mfa_secret FROM accounts WHERE email_normalized = 'ops@example.test'`).Scan(&stored); err != nil {
		t.Fatal(err)
	}
	if stored == secret || !strings.HasPrefix(stored, "enc:v1:") {
		t.Fatalf("MFA secret stored without encryption: %q", stored)
	}
	_, res := e.login("ops@example.test", "senha-segura-123", "")
	if res.Status != http.StatusForbidden || errorCode(t, res.Body) != "MFA_REQUIRED" {
		t.Fatalf("no MFA status=%d body=%s", res.Status, res.Body)
	}
	_, res = e.login("ops@example.test", "senha-segura-123", "000000")
	if res.Status != http.StatusUnauthorized {
		t.Fatalf("wrong MFA status=%d body=%s", res.Status, res.Body)
	}
	code, _ := security.GenerateTOTP(secret, time.Now())
	_, res = e.login("ops@example.test", "senha-segura-123", code)
	if res.Status != http.StatusOK {
		t.Fatalf("valid MFA status=%d body=%s", res.Status, res.Body)
	}
	_, res = e.login("ops@example.test", "senha-segura-123", codes[0])
	if res.Status != http.StatusOK {
		t.Fatalf("recovery status=%d body=%s", res.Status, res.Body)
	}
	_, res = e.login("ops@example.test", "senha-segura-123", codes[0])
	if res.Status != http.StatusUnauthorized {
		t.Fatalf("reused recovery status=%d body=%s", res.Status, res.Body)
	}
}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}
