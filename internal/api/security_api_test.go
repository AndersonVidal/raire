package api_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/raire/raire/internal/platform/config"
)

func TestCSRFRequired(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")

	res := e.do("POST", "/v1/me/password", map[string]any{
		"current_password": "senha-segura-123", "new_password": "nova-senha-456",
	}, nil)
	if res.Status != http.StatusForbidden || errorCode(t, res.Body) != "CSRF_INVALID" {
		t.Fatalf("missing csrf status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
	}, nil)
	if res.Status != http.StatusForbidden || errorCode(t, res.Body) != "CSRF_INVALID" {
		t.Fatalf("missing csrf profile status=%d body=%s", res.Status, res.Body)
	}
	var count int
	_ = e.app.DB.QueryRow(`SELECT count(*) FROM professional_profiles`).Scan(&count)
	if count != 0 {
		t.Fatal("profile persisted without CSRF")
	}
	res = e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("valid csrf status=%d body=%s", res.Status, res.Body)
	}
}

func TestSignupRateLimit(t *testing.T) {
	e := newTestEnv(t, func(c *config.Config) { c.RateLimitEnabled = true })
	for i := 0; i < 3; i++ {
		token := e.signupEmailOnly("user"+string(rune('a'+i))+"@example.test", "senha-segura-123")
		if token == "" {
			t.Fatal("missing token")
		}
	}
	res := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": "fourth@example.test", "password": "senha-segura-123", "display_name": "Quarto",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if res.Status != http.StatusTooManyRequests {
		t.Fatalf("fourth signup status=%d body=%s", res.Status, res.Body)
	}
	if res.Header.Get("Retry-After") == "" {
		t.Fatal("missing Retry-After")
	}
}

func TestLoginRateLimit(t *testing.T) {
	e := newTestEnv(t, func(c *config.Config) { c.RateLimitEnabled = true })
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	for i := 0; i < 5; i++ {
		_, res := e.login("ana@example.test", "senha-errada-000", "")
		if res.Status != http.StatusUnauthorized {
			t.Fatalf("attempt %d status=%d body=%s", i+1, res.Status, res.Body)
		}
	}
	_, res := e.login("ana@example.test", "senha-errada-000", "")
	if res.Status != http.StatusTooManyRequests {
		t.Fatalf("sixth attempt status=%d body=%s", res.Status, res.Body)
	}
}

func TestResetRateLimit(t *testing.T) {
	e := newTestEnv(t, func(c *config.Config) { c.RateLimitEnabled = true })
	for i := 0; i < 3; i++ {
		res := e.do("POST", "/v1/auth/password-reset/request", map[string]any{"email": "ghost@example.test"}, nil)
		if res.Status != http.StatusAccepted {
			t.Fatalf("reset %d status=%d body=%s", i+1, res.Status, res.Body)
		}
	}
	res := e.do("POST", "/v1/auth/password-reset/request", map[string]any{"email": "ghost@example.test"}, nil)
	if res.Status != http.StatusTooManyRequests {
		t.Fatalf("fourth reset status=%d body=%s", res.Status, res.Body)
	}
}

func TestResendRateLimit(t *testing.T) {
	e := newTestEnv(t, func(c *config.Config) { c.RateLimitEnabled = true })
	for i := 0; i < 5; i++ {
		res := e.do("POST", "/v1/auth/resend-confirmation", map[string]any{"email": "ghost@example.test"}, nil)
		if res.Status != http.StatusAccepted {
			t.Fatalf("resend %d status=%d body=%s", i+1, res.Status, res.Body)
		}
	}
	res := e.do("POST", "/v1/auth/resend-confirmation", map[string]any{"email": "ghost@example.test"}, nil)
	if res.Status != http.StatusTooManyRequests {
		t.Fatalf("sixth resend status=%d body=%s", res.Status, res.Body)
	}
}

func TestSignupIdempotency(t *testing.T) {
	e := newTestEnv(t)
	body := map[string]any{
		"email": "ana@example.test", "password": "senha-segura-123", "display_name": "Ana",
		"accept_terms": true, "accept_privacy": true,
	}
	headers := map[string]string{"Idempotency-Key": "signup-key-1"}
	res := e.do("POST", "/v1/auth/signup", body, headers)
	if res.Status != http.StatusAccepted {
		t.Fatalf("first status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/auth/signup", body, headers)
	if res.Status != http.StatusAccepted {
		t.Fatalf("replay status=%d body=%s", res.Status, res.Body)
	}
	body["display_name"] = "Outra"
	res = e.do("POST", "/v1/auth/signup", body, headers)
	if res.Status != http.StatusConflict {
		t.Fatalf("conflict status=%d body=%s", res.Status, res.Body)
	}
	var count int
	_ = e.app.DB.QueryRow(`SELECT count(*) FROM accounts`).Scan(&count)
	if count != 1 {
		t.Fatalf("accounts = %d", count)
	}
}

func TestConfirmIdempotency(t *testing.T) {
	e := newTestEnv(t)
	token := e.signup("ana@example.test", "senha-segura-123")
	headers := map[string]string{"Idempotency-Key": "confirm-key-1"}
	if res := e.do("POST", "/v1/auth/confirm-email", map[string]any{"token": token}, headers); res.Status != http.StatusNoContent {
		t.Fatalf("first confirm status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("POST", "/v1/auth/confirm-email", map[string]any{"token": token}, headers); res.Status != http.StatusNoContent {
		t.Fatalf("idempotent confirm status=%d body=%s", res.Status, res.Body)
	}
}

func TestLogsHaveCorrelationAndNoPII(t *testing.T) {
	e := newTestEnv(t)
	res := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": "log@example.test", "password": "senha-segura-123", "display_name": "Log",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if res.Header.Get("X-Correlation-ID") == "" {
		t.Fatal("missing correlation header")
	}
	logs := e.logs.String()
	if !strings.Contains(logs, `"correlation_id"`) {
		t.Fatalf("logs missing correlation_id: %s", logs)
	}
	if strings.Contains(logs, "log@example.test") || strings.Contains(logs, "senha-segura-123") {
		t.Fatalf("logs contain PII: %s", logs)
	}
}

func (e *testEnv) signupEmailOnly(email, password string) string {
	e.t.Helper()
	res := e.do("POST", "/v1/auth/signup", map[string]any{
		"email": email, "password": password, "display_name": "Usuário",
		"accept_terms": true, "accept_privacy": true,
	}, nil)
	if res.Status != http.StatusAccepted {
		e.t.Fatalf("signup status=%d body=%s", res.Status, res.Body)
	}
	return res.Header.Get("X-Dev-Confirm-Token")
}
