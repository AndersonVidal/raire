package api_test

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOpenAPIContract(t *testing.T) {
	raw, err := os.ReadFile("../../api/openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		OpenAPI    string                    `yaml:"openapi"`
		Paths      map[string]map[string]any `yaml:"paths"`
		Components struct {
			Schemas map[string]any `yaml:"schemas"`
		} `yaml:"components"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("invalid openapi yaml: %v", err)
	}
	if doc.OpenAPI != "3.1.0" {
		t.Fatalf("openapi version = %q", doc.OpenAPI)
	}
	required := map[string][]string{
		"/v1/healthz":                             {"get"},
		"/v1/readyz":                              {"get"},
		"/v1/auth/signup":                         {"post"},
		"/v1/auth/login":                          {"post"},
		"/v1/auth/logout":                         {"post"},
		"/v1/auth/confirm-email":                  {"post"},
		"/v1/auth/resend-confirmation":            {"post"},
		"/v1/auth/password-reset/request":         {"post"},
		"/v1/auth/password-reset/confirm":         {"post"},
		"/v1/me":                                  {"get", "patch"},
		"/v1/me/password":                         {"post"},
		"/v1/me/consents":                         {"get"},
		"/v1/me/consents/{kind}/revoke":           {"post"},
		"/v1/categories":                          {"get"},
		"/v1/categories/{slug}":                   {"get"},
		"/v1/professionals/me":                    {"get", "post", "patch"},
		"/v1/professionals/me/publish":            {"post"},
		"/v1/professionals/me/pause":              {"post"},
		"/v1/professionals/me/resume":             {"post"},
		"/v1/professionals/{id}":                  {"get"},
		"/v1/professionals/{id}/preview":          {"get"},
		"/v1/admin/categories":                    {"get", "post"},
		"/v1/admin/categories/{id}":               {"patch"},
		"/v1/admin/categories/{id}/deactivate":    {"post"},
		"/v1/admin/professionals":                 {"get"},
		"/v1/admin/professionals/{id}/suspend":    {"post"},
		"/v1/admin/professionals/{id}/reactivate": {"post"},
		"/v1/admin/audit-events":                  {"get"},
		"/v1/admin/feature-flags":                 {"get"},
		"/v1/admin/feature-flags/{key}":           {"patch"},
	}
	for path, methods := range required {
		item, ok := doc.Paths[path]
		if !ok {
			t.Fatalf("missing path %s", path)
		}
		for _, method := range methods {
			if _, ok := item[method]; !ok {
				t.Fatalf("missing %s %s", method, path)
			}
		}
	}
	for _, schema := range []string{"Error", "Account", "Category", "Window", "ProfileWrite", "PublicProfile"} {
		if _, ok := doc.Components.Schemas[schema]; !ok {
			t.Fatalf("missing schema %s", schema)
		}
	}
	protected := map[string][]string{
		"/v1/me":                                  {"get", "patch"},
		"/v1/me/password":                         {"post"},
		"/v1/me/consents":                         {"get"},
		"/v1/me/consents/{kind}/revoke":           {"post"},
		"/v1/professionals/me":                    {"get", "post", "patch"},
		"/v1/professionals/me/publish":            {"post"},
		"/v1/professionals/me/pause":              {"post"},
		"/v1/professionals/me/resume":             {"post"},
		"/v1/professionals/{id}/preview":          {"get"},
		"/v1/admin/categories":                    {"get", "post"},
		"/v1/admin/categories/{id}":               {"patch"},
		"/v1/admin/categories/{id}/deactivate":    {"post"},
		"/v1/admin/professionals":                 {"get"},
		"/v1/admin/professionals/{id}/suspend":    {"post"},
		"/v1/admin/professionals/{id}/reactivate": {"post"},
		"/v1/admin/audit-events":                  {"get"},
		"/v1/admin/feature-flags":                 {"get"},
		"/v1/admin/feature-flags/{key}":           {"patch"},
	}
	for path, methods := range protected {
		for _, method := range methods {
			op, ok := doc.Paths[path][method].(map[string]any)
			if !ok {
				t.Fatalf("operation %s %s has unexpected shape", method, path)
			}
			if _, ok := op["security"]; !ok {
				t.Fatalf("protected operation %s %s has no security requirement", method, path)
			}
		}
	}
	if _, ok := doc.Paths["/v1/auth/signup"]["post"].(map[string]any)["responses"].(map[string]any)["201"]; ok {
		t.Fatal("signup must not document a distinguishable 201 response")
	}
}
