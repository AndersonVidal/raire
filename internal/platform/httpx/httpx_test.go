package httpx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/raire/raire/internal/contracts"
)

func TestCorrelationAndError(t *testing.T) {
	h := RequestCorrelation(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if CorrelationID(r.Context()) == "" {
			t.Fatal("missing correlation id")
		}
		WriteError(w, r, contracts.New(contracts.CodeForbidden, ""))
	}))
	req := httptest.NewRequest("GET", "/x", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("status = %d", rec.Code)
	}
	if rec.Header().Get("X-Correlation-ID") == "" {
		t.Fatal("missing response correlation header")
	}
	var body errorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Error.Code != "FORBIDDEN" || body.Error.CorrelationID == "" {
		t.Fatalf("body = %+v", body)
	}
}

func TestDecodeJSONRejectsUnknownAndTrailing(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}
	req := httptest.NewRequest("POST", "/x", strings.NewReader(`{"name":"a","extra":1}`))
	rec := httptest.NewRecorder()
	var p payload
	if err := DecodeJSON(rec, req, &p, 1024); err == nil {
		t.Fatal("unknown field accepted")
	}
	req = httptest.NewRequest("POST", "/x", strings.NewReader(`{"name":"a"} {"name":"b"}`))
	rec = httptest.NewRecorder()
	if err := DecodeJSON(rec, req, &p, 1024); err == nil {
		t.Fatal("trailing object accepted")
	}
}

func TestClientIP(t *testing.T) {
	req := httptest.NewRequest("GET", "/x", nil)
	req.RemoteAddr = "10.0.0.1:1234"
	if got := ClientIP(req, false); got != "10.0.0.1" {
		t.Fatalf("ip = %q", got)
	}
	req.Header.Set("X-Forwarded-For", "203.0.113.9, 10.0.0.1")
	if got := ClientIP(req, true); got != "203.0.113.9" {
		t.Fatalf("forwarded ip = %q", got)
	}
	if got := ClientIP(req, false); got != "10.0.0.1" {
		t.Fatalf("untrusted proxy ip = %q", got)
	}
}

func TestCorrelationID(t *testing.T) {
	if got := CorrelationID(context.Background()); got != "" {
		t.Fatalf("empty context id = %q", got)
	}
}
