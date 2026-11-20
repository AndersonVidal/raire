package api_test

import (
	"net/http"
	"strings"
	"testing"
)

func TestPublicTextRejectsPII(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
		"bio": "Fale comigo em ana@example.test",
	}, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "PII_NOT_ALLOWED" {
		t.Fatalf("bio email status=%d body=%s", res.Status, res.Body)
	}
	var count int
	_ = e.app.DB.QueryRow(`SELECT count(*) FROM professional_profiles`).Scan(&count)
	if count != 0 {
		t.Fatal("profile persisted with PII in bio")
	}
	res = e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
		"bio": "Atendo em Cidade Piloto.", "availability_note": "Ligue 11 99999-8888",
	}, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "PII_NOT_ALLOWED" {
		t.Fatalf("availability phone status=%d body=%s", res.Status, res.Body)
	}
}

func TestPhoneVerificationSurvivesSameNumber(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	if res := e.do("PATCH", "/v1/me", map[string]any{"contact_phone": "+55 11 99999-8888"}, authHeaders(csrf)); res.Status != http.StatusOK {
		t.Fatalf("set phone status=%d body=%s", res.Status, res.Body)
	}
	if _, err := e.app.DB.Exec(`UPDATE accounts SET contact_phone_verified_at = ? WHERE email_normalized = 'ana@example.test'`,
		"2026-11-20T00:00:00Z"); err != nil {
		t.Fatal(err)
	}
	// Re-saving the same number must not destroy verification.
	res := e.do("PATCH", "/v1/me", map[string]any{"contact_phone": "+55 11 99999-8888"}, authHeaders(csrf))
	if res.Status != http.StatusOK || !strings.Contains(string(res.Body), `"contact_phone_verified":true`) {
		t.Fatalf("same phone status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("PATCH", "/v1/me", map[string]any{"contact_phone": "+55 11 98888-7777"}, authHeaders(csrf))
	if res.Status != http.StatusOK || !strings.Contains(string(res.Body), `"contact_phone_verified":false`) {
		t.Fatalf("changed phone status=%d body=%s", res.Status, res.Body)
	}
}

func TestResetIdempotency(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	requestHeaders := map[string]string{"Idempotency-Key": "reset-request-1"}
	res := e.do("POST", "/v1/auth/password-reset/request", map[string]any{"email": "ana@example.test"}, requestHeaders)
	if res.Status != http.StatusAccepted {
		t.Fatalf("request status=%d body=%s", res.Status, res.Body)
	}
	token := res.Header.Get("X-Dev-Reset-Token")
	if token == "" {
		t.Fatal("missing reset token")
	}
	confirmHeaders := map[string]string{"Idempotency-Key": "reset-confirm-1"}
	body := map[string]any{"token": token, "new_password": "nova-senha-456"}
	if res := e.do("POST", "/v1/auth/password-reset/confirm", body, confirmHeaders); res.Status != http.StatusNoContent {
		t.Fatalf("confirm status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("POST", "/v1/auth/password-reset/confirm", body, confirmHeaders); res.Status != http.StatusNoContent {
		t.Fatalf("idempotent confirm status=%d body=%s", res.Status, res.Body)
	}
	e.loginOK("ana@example.test", "nova-senha-456")
}
