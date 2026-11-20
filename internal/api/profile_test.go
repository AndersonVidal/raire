package api_test

import (
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
	"time"
)

func TestPublishedProfileHasNoPII(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("PATCH", "/v1/me", map[string]any{
		"contact_phone": "+55 11 99999-8888", "contact_preference": "email",
	}, authHeaders(csrf))
	if res.Status != http.StatusOK {
		t.Fatalf("update me status=%d body=%s", res.Status, res.Body)
	}
	body := map[string]any{
		"display_name": "Ana Souza", "bio": "Eletricista", "service_type": "in_person",
		"category_ids": []string{"cat_eletricista"}, "city": "Cidade Piloto", "neighborhood": "Centro",
		"radius_km": 10, "precise_lat": -23.550520, "precise_lng": -46.633308,
		"availability": []map[string]any{{"day": 1, "start": "08:00", "end": "12:00", "timezone": "America/Sao_Paulo"}},
	}
	res = e.do("POST", "/v1/professionals/me", body, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	created := decode[struct {
		Profile struct {
			ID string `json:"id"`
		} `json:"profile"`
	}](t, res.Body)
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusOK {
		t.Fatalf("publish status=%d body=%s", res.Status, res.Body)
	}
	pub := e.do("GET", "/v1/professionals/"+created.Profile.ID, nil, nil)
	if pub.Status != http.StatusOK {
		t.Fatalf("public status=%d body=%s", pub.Status, pub.Body)
	}
	raw := string(pub.Body)
	for _, forbidden := range []string{"ana@example.test", "99999", "999998888", "-23.55", "precise_lat", "precise_lng"} {
		if strings.Contains(raw, forbidden) {
			t.Fatalf("public profile leaked %q: %s", forbidden, raw)
		}
	}
	owner := e.do("GET", "/v1/professionals/me", nil, nil)
	if owner.Status != http.StatusOK {
		t.Fatalf("owner status=%d body=%s", owner.Status, owner.Body)
	}
	ownerBody := string(owner.Body)
	if strings.Contains(ownerBody, "99999-8888") || strings.Contains(ownerBody, "-23.550520") {
		t.Fatalf("owner view leaked raw PII: %s", ownerBody)
	}
	view := decode[struct {
		Profile struct {
			Contact struct {
				EffectiveChannel string `json:"effective_channel"`
				EmailMasked      string `json:"email_masked"`
				PhoneMasked      string `json:"phone_masked"`
				PhoneVerified    bool   `json:"phone_verified"`
			} `json:"contact"`
			PreciseLocationSet bool `json:"precise_location_set"`
		} `json:"profile"`
	}](t, owner.Body)
	if view.Profile.Contact.EffectiveChannel != "email" {
		t.Fatalf("effective channel = %s", view.Profile.Contact.EffectiveChannel)
	}
	if view.Profile.Contact.EmailMasked != "a***@example.test" {
		t.Fatalf("email masked = %q", view.Profile.Contact.EmailMasked)
	}
	if view.Profile.Contact.PhoneMasked != "" || view.Profile.Contact.PhoneVerified {
		t.Fatal("unverified phone was exposed")
	}
	if !view.Profile.PreciseLocationSet {
		t.Fatal("precise location flag missing")
	}
	preview := e.do("GET", "/v1/professionals/"+created.Profile.ID+"/preview", nil, nil)
	if preview.Status != http.StatusOK {
		t.Fatalf("preview status=%d body=%s", preview.Status, preview.Body)
	}
	if strings.Contains(string(preview.Body), "ana@example.test") || strings.Contains(string(preview.Body), "99999") {
		t.Fatalf("preview leaked PII: %s", preview.Body)
	}
}

func TestRemoteProfileWithoutArea(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana Remota", "service_type": "remote",
		"category_ids": []string{"cat_professor_particular"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create remote status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusOK {
		t.Fatalf("publish remote status=%d body=%s", res.Status, res.Body)
	}
}

func TestHybridWithoutAreaIsRejected(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana Híbrida", "service_type": "hybrid",
		"category_ids": []string{"cat_eletricista"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create hybrid status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "INVALID_SERVICE_AREA" {
		t.Fatalf("publish hybrid status=%d body=%s", res.Status, res.Body)
	}
	owner := e.do("GET", "/v1/professionals/me", nil, nil)
	view := decode[struct {
		Profile struct {
			Status string `json:"status"`
		} `json:"profile"`
	}](t, owner.Body)
	if view.Profile.Status != "draft" {
		t.Fatalf("status after failed publish = %s", view.Profile.Status)
	}
}

func TestEmailNotConfirmedCannotPublish(t *testing.T) {
	e := newTestEnv(t)
	e.signup("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusForbidden || errorCode(t, res.Body) != "EMAIL_NOT_CONFIRMED" {
		t.Fatalf("publish unconfirmed status=%d body=%s", res.Status, res.Body)
	}
}

func TestPublishRequiresCategory(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote",
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "PROFILE_INCOMPLETE" {
		t.Fatalf("publish no category status=%d body=%s", res.Status, res.Body)
	}
}

func TestRadiusOutOfRange(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "in_person", "category_ids": []string{"cat_eletricista"},
		"city": "Cidade Piloto", "neighborhood": "Centro", "radius_km": 60,
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "INVALID_SERVICE_AREA" {
		t.Fatalf("publish radius status=%d body=%s", res.Status, res.Body)
	}
}

func TestInvalidAvailabilityIsRejected(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
		"availability": []map[string]any{
			{"day": 1, "start": "08:00", "end": "12:00", "timezone": "America/Sao_Paulo"},
			{"day": 1, "start": "11:00", "end": "13:00", "timezone": "America/Sao_Paulo"},
		},
	}, authHeaders(csrf))
	if res.Status != http.StatusUnprocessableEntity || errorCode(t, res.Body) != "INVALID_AVAILABILITY" {
		t.Fatalf("overlap status=%d body=%s", res.Status, res.Body)
	}
	var count int
	_ = e.app.DB.QueryRow(`SELECT count(*) FROM professional_profiles`).Scan(&count)
	if count != 0 {
		t.Fatalf("profile persisted after invalid availability: %d", count)
	}
}

func TestPauseResumeProfile(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	id := e.publishProfile(e.client, csrf)
	if res := e.do("POST", "/v1/professionals/me/pause", nil, authHeaders(csrf)); res.Status != http.StatusOK {
		t.Fatalf("pause status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/professionals/"+id, nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("paused public status=%d", res.Status)
	}
	if res := e.do("POST", "/v1/professionals/me/resume", nil, authHeaders(csrf)); res.Status != http.StatusOK {
		t.Fatalf("resume status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/professionals/"+id, nil, nil); res.Status != http.StatusOK {
		t.Fatalf("resumed public status=%d", res.Status)
	}
}

func TestCategoryDeactivatedHidesProfile(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	id := e.publishProfile(e.client, csrf)

	opCSRF, _, _ := e.operator("ops@example.test", "senha-segura-123", "operator")
	res := e.do("POST", "/v1/admin/categories/cat_eletricista/deactivate", nil, authHeaders(opCSRF))
	if res.Status != http.StatusNoContent {
		t.Fatalf("deactivate status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/professionals/"+id, nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("inactive category public status=%d", res.Status)
	}
}

func TestContactPreferenceResolution(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("PATCH", "/v1/me", map[string]any{
		"contact_phone": "+55 11 99999-8888", "contact_preference": "email",
	}, authHeaders(csrf))
	if res.Status != http.StatusOK {
		t.Fatalf("update me status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "remote", "category_ids": []string{"cat_consultor"},
		"contact_preference": "phone",
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}

	owner := e.do("GET", "/v1/professionals/me", nil, nil)
	view := decode[struct {
		Profile struct {
			Contact struct {
				EffectiveChannel string `json:"effective_channel"`
				PhoneMasked      string `json:"phone_masked"`
				PhoneVerified    bool   `json:"phone_verified"`
			} `json:"contact"`
		} `json:"profile"`
	}](t, owner.Body)
	if view.Profile.Contact.EffectiveChannel != "email" {
		t.Fatalf("unverified phone effective = %s", view.Profile.Contact.EffectiveChannel)
	}
	if view.Profile.Contact.PhoneMasked != "" {
		t.Fatalf("unverified phone masked = %q", view.Profile.Contact.PhoneMasked)
	}
	if _, err := e.app.DB.Exec(`UPDATE accounts SET contact_phone_verified_at = ? WHERE email_normalized = 'ana@example.test'`,
		time.Now().UTC().Format(time.RFC3339Nano)); err != nil {
		t.Fatal(err)
	}
	owner = e.do("GET", "/v1/professionals/me", nil, nil)
	view = decode[struct {
		Profile struct {
			Contact struct {
				EffectiveChannel string `json:"effective_channel"`
				PhoneMasked      string `json:"phone_masked"`
				PhoneVerified    bool   `json:"phone_verified"`
			} `json:"contact"`
		} `json:"profile"`
	}](t, owner.Body)
	if view.Profile.Contact.EffectiveChannel != "phone" || !view.Profile.Contact.PhoneVerified {
		t.Fatalf("verified phone not effective: %+v", view.Profile.Contact)
	}
	if view.Profile.Contact.PhoneMasked != "****8888" {
		t.Fatalf("phone masked = %q", view.Profile.Contact.PhoneMasked)
	}
}

func TestPreciseLocationEncryptedAndPurged(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrf := e.loginOK("ana@example.test", "senha-segura-123")
	res := e.do("POST", "/v1/professionals/me", map[string]any{
		"display_name": "Ana", "service_type": "in_person", "category_ids": []string{"cat_eletricista"},
		"city": "Cidade Piloto", "neighborhood": "Centro", "radius_km": 10,
		"precise_lat": -23.550520, "precise_lng": -46.633308,
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	var encLat, encLng string
	if err := e.app.DB.QueryRow(`SELECT precise_lat_enc, precise_lng_enc FROM professional_profiles LIMIT 1`).Scan(&encLat, &encLng); err != nil {
		t.Fatal(err)
	}
	if encLat == "" || encLng == "" {
		t.Fatal("precise location not stored")
	}
	if strings.Contains(encLat, "23.55") || strings.Contains(encLng, "46.63") {
		t.Fatal("precise location stored in plaintext")
	}
	n, err := e.app.Directory.PurgeExpiredLocations(t.Context(), time.Now().Add(91*24*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("purged = %d", n)
	}
	var nullLat, nullLng *string
	if err := e.app.DB.QueryRow(`SELECT precise_lat_enc, precise_lng_enc FROM professional_profiles LIMIT 1`).Scan(&nullLat, &nullLng); err != nil {
		t.Fatal(err)
	}
	if nullLat != nil || nullLng != nil {
		t.Fatal("precise location was not purged")
	}
	var audits int
	if err := e.app.DB.QueryRow(`SELECT count(*) FROM audit_events WHERE action = 'directory.location_purged'`).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 1 {
		t.Fatalf("purge audits = %d", audits)
	}
}

func TestPreviewOwnership(t *testing.T) {
	e := newTestEnv(t)
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	csrfA := e.loginOK("ana@example.test", "senha-segura-123")
	id := e.publishProfile(e.client, csrfA)

	jar, _ := cookiejar.New(nil)
	clientB := &http.Client{Jar: jar, Timeout: 10 * time.Second}
	jarB, _ := cookiejar.New(nil)
	_ = jarB
	// Signup is session-less; use a fresh client for B's login.
	e.signupConfirmed("bia@example.test", "senha-segura-123")
	res := e.doWith(clientB, "POST", "/v1/auth/login", map[string]any{
		"email": "bia@example.test", "password": "senha-segura-123",
	}, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("B login status=%d body=%s", res.Status, res.Body)
	}
	res = e.doWith(clientB, "GET", "/v1/professionals/"+id+"/preview", nil, nil)
	if res.Status != http.StatusNotFound {
		t.Fatalf("foreign preview status=%d body=%s", res.Status, res.Body)
	}
}
