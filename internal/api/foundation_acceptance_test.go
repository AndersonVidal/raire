package api_test

import (
	"database/sql"
	"net/http"
	"testing"
)

func TestDraftIsNotPublic(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	created := decode[struct {
		Profile struct {
			ID string `json:"id"`
		} `json:"profile"`
	}](t, res.Body)
	if res := e.do("GET", "/v1/professionals/"+created.Profile.ID, nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("draft public status=%d", res.Status)
	}
}

func TestConsentVersioningAndRevocation(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("GET", "/v1/me/consents", nil, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("consents status=%d body=%s", res.Status, res.Body)
	}
	items := decode[struct {
		Items []struct {
			Kind    string `json:"kind"`
			Version string `json:"version"`
		} `json:"items"`
	}](t, res.Body)
	if len(items.Items) != 2 {
		t.Fatalf("consents = %+v", items.Items)
	}
	for _, c := range items.Items {
		if c.Version != "2026-11-20" {
			t.Fatalf("version = %s", c.Version)
		}
	}
	if _, err := e.app.DB.Exec(`INSERT INTO consents (id, account_id, kind, version, granted_at)
		SELECT 'cns_research', id, 'research', '2026-11-20', '2026-11-20T00:00:00Z' FROM accounts LIMIT 1`); err != nil {
		t.Fatal(err)
	}
	res = e.do("POST", "/v1/me/consents/research/revoke", nil, authHeaders(csrf))
	if res.Status != http.StatusNoContent {
		t.Fatalf("revoke status=%d body=%s", res.Status, res.Body)
	}
	var revoked sql.NullString
	if err := e.app.DB.QueryRow(`SELECT revoked_at FROM consents WHERE id = 'cns_research'`).Scan(&revoked); err != nil {
		t.Fatal(err)
	}
	if !revoked.Valid {
		t.Fatal("research consent not revoked")
	}
	if res := e.do("POST", "/v1/me/consents/terms/revoke", nil, authHeaders(csrf)); res.Status != http.StatusConflict {
		t.Fatalf("terms revoke status=%d body=%s", res.Status, res.Body)
	}
}

func TestAccountCanBeClientAndProfessional(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	me := e.do("GET", "/v1/me", nil, nil)
	if me.Status != http.StatusOK {
		t.Fatalf("me status=%d body=%s", me.Status, me.Body)
	}
	account := decode[struct {
		Account struct {
			Roles []string `json:"roles"`
		} `json:"account"`
	}](t, me.Body)
	hasClient, hasProfessional := false, false
	for _, role := range account.Account.Roles {
		if role == "client" {
			hasClient = true
		}
		if role == "professional" {
			hasProfessional = true
		}
	}
	if !hasClient || !hasProfessional {
		t.Fatalf("roles = %v", account.Account.Roles)
	}
}

func TestVerificationBadgeIsHonest(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	id := e.publishProfile(e.client, csrf)
	res := e.do("GET", "/v1/professionals/"+id, nil, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("public status=%d body=%s", res.Status, res.Body)
	}
	view := decode[struct {
		Profile struct {
			Verification string `json:"verification"`
		} `json:"profile"`
	}](t, res.Body)
	if view.Profile.Verification != "contact_verified" {
		t.Fatalf("verification = %s", view.Profile.Verification)
	}
}

func TestFunnelMetrics(t *testing.T) {
	e := newTestEnv(t)
	token := e.signup("ana@example.test", "senha-segura-123")
	if res := e.confirm(token); res.Status != http.StatusNoContent {
		t.Fatalf("confirm status=%d", res.Status)
	}
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	e.publishProfile(e.client, csrf)
	snap := e.app.Metrics.Snapshot()
	for _, metric := range []string{"signup_completed", "email_confirmed", "login_succeeded", "profile_created", "profile_published"} {
		if snap[metric] == 0 {
			t.Fatalf("metric %s missing: %+v", metric, snap)
		}
	}
}
