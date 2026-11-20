package api_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
	"time"
)

func TestCatalogSynonymsAndInactive(t *testing.T) {
	e := newTestEnv(t)
	res := e.do("GET", "/v1/categories?q=instalacao", nil, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("search status=%d body=%s", res.Status, res.Body)
	}
	items := decode[struct {
		Items []struct {
			ID string `json:"id"`
		} `json:"items"`
	}](t, res.Body)
	if len(items.Items) == 0 || items.Items[0].ID != "cat_eletricista" {
		t.Fatalf("synonym search failed: %s", res.Body)
	}

	csrf, _, _ := e.operator("ops@example.test", "senha-segura-123", "operator")
	if res := e.do("POST", "/v1/admin/categories/cat_eletricista/deactivate", nil, authHeaders(csrf)); res.Status != http.StatusNoContent {
		t.Fatalf("deactivate status=%d body=%s", res.Status, res.Body)
	}
	res = e.do("GET", "/v1/categories?q=instalacao", nil, nil)
	if strings.Contains(string(res.Body), "cat_eletricista") {
		t.Fatalf("inactive category still searchable: %s", res.Body)
	}
	if res := e.do("GET", "/v1/categories/eletricista", nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("inactive category detail status=%d", res.Status)
	}
	admin := e.do("GET", "/v1/admin/categories", nil, nil)
	if admin.Status != http.StatusOK || !strings.Contains(string(admin.Body), `"active":false`) {
		t.Fatalf("admin list missing inactive: status=%d body=%s", admin.Status, admin.Body)
	}
}

func TestAdminCreatesAndUpdatesCategoryWithAudit(t *testing.T) {
	e := newTestEnv(t)
	csrf, _, _ := e.operator("ops@example.test", "senha-segura-123", "operator")
	res := e.do("POST", "/v1/admin/categories", map[string]any{
		"slug": "jardineiro", "name": "Jardineiro", "description": "Jardinagem residencial.",
		"synonyms": []string{"jardinagem", "poda"}, "sort_order": 60,
	}, authHeaders(csrf))
	if res.Status != http.StatusCreated {
		t.Fatalf("create status=%d body=%s", res.Status, res.Body)
	}
	created := decode[struct {
		Category struct {
			ID string `json:"id"`
		} `json:"category"`
	}](t, res.Body)
	res = e.do("PATCH", "/v1/admin/categories/"+created.Category.ID, map[string]any{
		"name": "Jardineiro e paisagista",
	}, authHeaders(csrf))
	if res.Status != http.StatusOK {
		t.Fatalf("update status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/categories/jardineiro", nil, nil); res.Status != http.StatusOK {
		t.Fatalf("public category status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("POST", "/v1/admin/categories/"+created.Category.ID+"/deactivate", nil, authHeaders(csrf)); res.Status != http.StatusNoContent {
		t.Fatalf("deactivate status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/categories/jardineiro", nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("inactive public status=%d", res.Status)
	}
	for _, action := range []string{"catalog.category_created", "catalog.category_updated", "catalog.category_deactivated"} {
		var count int
		if err := e.app.DB.QueryRow(`SELECT count(*) FROM audit_events WHERE action = ?`, action).Scan(&count); err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Fatalf("audit %s count=%d", action, count)
		}
	}
}

func TestAdminSuspendsReactivatesAndAudits(t *testing.T) {
	e := newTestEnv(t)
	anaJar, _ := cookiejar.New(nil)
	ana := &http.Client{Jar: anaJar, Timeout: 10 * time.Second}
	e.signupConfirmed("ana@example.test", "senha-segura-123")
	anaCSRF, res := e.loginWith(ana, "ana@example.test", "senha-segura-123", "")
	if res.Status != http.StatusOK {
		t.Fatalf("ana login status=%d body=%s", res.Status, res.Body)
	}
	id := e.publishProfile(ana, anaCSRF)

	opJar, _ := cookiejar.New(nil)
	op := &http.Client{Jar: opJar, Timeout: 10 * time.Second}
	opCSRF, _, _ := e.operatorWith(op, "ops@example.test", "senha-segura-123", "operator")

	res = e.doWith(op, "POST", "/v1/admin/professionals/"+id+"/suspend", map[string]any{
		"reason": "Denúncia de fraude em análise.", "evidence": "case-123",
	}, authHeaders(opCSRF))
	if res.Status != http.StatusNoContent {
		t.Fatalf("suspend status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/professionals/"+id, nil, nil); res.Status != http.StatusNotFound {
		t.Fatalf("suspended public status=%d", res.Status)
	}
	if res := e.doWith(ana, "GET", "/v1/me", nil, nil); res.Status != http.StatusUnauthorized {
		t.Fatalf("suspended session status=%d", res.Status)
	}

	audit := e.doWith(op, "GET", "/v1/admin/audit-events?target_type=profile&target_id="+id, nil, nil)
	if audit.Status != http.StatusOK {
		t.Fatalf("audit status=%d body=%s", audit.Status, audit.Body)
	}
	if !strings.Contains(string(audit.Body), "directory.profile_suspended") ||
		!strings.Contains(string(audit.Body), "Denúncia de fraude") ||
		!strings.Contains(string(audit.Body), "correlation_id") {
		t.Fatalf("audit missing fields: %s", audit.Body)
	}

	res = e.doWith(op, "POST", "/v1/admin/professionals/"+id+"/reactivate", nil, authHeaders(opCSRF))
	if res.Status != http.StatusNoContent {
		t.Fatalf("reactivate status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("GET", "/v1/professionals/"+id, nil, nil); res.Status != http.StatusOK {
		t.Fatalf("reactivated public status=%d", res.Status)
	}
	if _, res := e.loginWith(ana, "ana@example.test", "senha-segura-123", ""); res.Status != http.StatusOK {
		t.Fatalf("ana login after reactivation status=%d body=%s", res.Status, res.Body)
	}
}

func TestAdminPaginationAndFilters(t *testing.T) {
	e := newTestEnv(t)
	for i := 0; i < 3; i++ {
		email := fmt.Sprintf("prof%d@example.test", i)
		e.signupConfirmed(email, "senha-segura-123")
		csrf := e.loginOK(email, "senha-segura-123")
		res := e.do("POST", "/v1/professionals/me", map[string]any{
			"display_name": fmt.Sprintf("Prof Teste %d", i), "service_type": "remote",
			"category_ids": []string{"cat_consultor"},
		}, authHeaders(csrf))
		if res.Status != http.StatusCreated {
			t.Fatalf("create %d status=%d body=%s", i, res.Status, res.Body)
		}
		if res := e.do("POST", "/v1/professionals/me/publish", nil, authHeaders(csrf)); res.Status != http.StatusOK {
			t.Fatalf("publish %d status=%d body=%s", i, res.Status, res.Body)
		}
	}
	csrf, _, _ := e.operator("ops@example.test", "senha-segura-123", "operator")
	res := e.do("GET", "/v1/admin/professionals?limit=2&status=published", nil, nil)
	if res.Status != http.StatusOK {
		t.Fatalf("list status=%d body=%s", res.Status, res.Body)
	}
	page1 := decode[struct {
		Items      []map[string]any `json:"items"`
		NextCursor string           `json:"next_cursor"`
	}](t, res.Body)
	if len(page1.Items) != 2 || page1.NextCursor == "" {
		t.Fatalf("page1 = %+v", page1)
	}
	res = e.do("GET", "/v1/admin/professionals?limit=2&status=published&cursor="+page1.NextCursor, nil, nil)
	page2 := decode[struct {
		Items      []map[string]any `json:"items"`
		NextCursor string           `json:"next_cursor"`
	}](t, res.Body)
	if len(page2.Items) != 1 {
		t.Fatalf("page2 items=%d body=%s", len(page2.Items), res.Body)
	}
	res = e.do("GET", "/v1/admin/professionals?q=Prof+Teste+1&status=published", nil, nil)
	filtered := decode[struct {
		Items []map[string]any `json:"items"`
	}](t, res.Body)
	if len(filtered.Items) != 1 {
		t.Fatalf("filtered items=%d body=%s", len(filtered.Items), res.Body)
	}
	_ = csrf
}

func TestFeatureFlags(t *testing.T) {
	e := newTestEnv(t)
	if e.app.Flags.IsEnabled(context.Background(), "unknown.flag") {
		t.Fatal("unknown flag enabled")
	}
	opCSRF, _, _ := e.operator("ops@example.test", "senha-segura-123", "operator")
	if res := e.do("GET", "/v1/admin/feature-flags", nil, nil); res.Status != http.StatusForbidden {
		t.Fatalf("operator list flags status=%d", res.Status)
	}
	if res := e.do("PATCH", "/v1/admin/feature-flags/foundation.enabled", map[string]any{"enabled": true}, authHeaders(opCSRF)); res.Status != http.StatusForbidden {
		t.Fatalf("operator patch flag status=%d body=%s", res.Status, res.Body)
	}

	adminCSRF, _, _ := e.operator("admin@example.test", "senha-segura-123", "admin")
	if res := e.do("GET", "/v1/admin/feature-flags", nil, nil); res.Status != http.StatusOK {
		t.Fatalf("admin list flags status=%d body=%s", res.Status, res.Body)
	}
	if res := e.do("PATCH", "/v1/admin/feature-flags/foundation.enabled", map[string]any{"enabled": true}, authHeaders(adminCSRF)); res.Status != http.StatusNoContent {
		t.Fatalf("admin patch flag status=%d body=%s", res.Status, res.Body)
	}
	if !e.app.Flags.IsEnabled(context.Background(), "foundation.enabled") {
		t.Fatal("flag not enabled")
	}
	if res := e.do("PATCH", "/v1/admin/feature-flags/does.not.exist", map[string]any{"enabled": true}, authHeaders(adminCSRF)); res.Status != http.StatusNotFound {
		t.Fatalf("unknown flag patch status=%d body=%s", res.Status, res.Body)
	}
	var audits int
	if err := e.app.DB.QueryRow(`SELECT count(*) FROM audit_events WHERE action = 'platform.feature_flag_changed'`).Scan(&audits); err != nil {
		t.Fatal(err)
	}
	if audits != 1 {
		t.Fatalf("flag audits = %d", audits)
	}
}

func TestHealthAndReadiness(t *testing.T) {
	e := newTestEnv(t)
	if res := e.do("GET", "/v1/healthz", nil, nil); res.Status != http.StatusOK {
		t.Fatalf("health status=%d", res.Status)
	}
	if res := e.do("GET", "/v1/readyz", nil, nil); res.Status != http.StatusOK {
		t.Fatalf("ready status=%d body=%s", res.Status, res.Body)
	}
	if err := e.app.DB.Close(); err != nil {
		t.Fatal(err)
	}
	if res := e.do("GET", "/v1/readyz", nil, nil); res.Status != http.StatusServiceUnavailable {
		t.Fatalf("ready after close status=%d body=%s", res.Status, res.Body)
	}
}
