package httpx

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/raire/raire/internal/contracts"
)

type correlationKey struct{}

func RequestCorrelation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(r.Header.Get("X-Correlation-ID"))
		if !validCorrelationID(id) {
			id = newCorrelationID()
		}
		ctx := context.WithValue(r.Context(), correlationKey{}, id)
		w.Header().Set("X-Correlation-ID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CorrelationID(ctx context.Context) string {
	v, _ := ctx.Value(correlationKey{}).(string)
	return v
}

func validCorrelationID(id string) bool {
	if len(id) < 8 || len(id) > 64 {
		return false
	}
	for _, r := range id {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_') {
			return false
		}
	}
	return true
}

func newCorrelationID() string {
	var b [12]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "no-referrer")
		h.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; base-uri 'none'")
		if strings.HasPrefix(r.URL.Path, "/v1/") {
			// API responses can contain account data; never let a shared cache
			// retain them.
			h.Set("Cache-Control", "no-store")
		}
		next.ServeHTTP(w, r)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func Logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(rec, r)
			logger.Info("http_request",
				"correlation_id", CorrelationID(r.Context()),
				"method", r.Method,
				"path", r.URL.Path,
				"status", rec.status,
				"duration_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}

func Recover(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panic_recovered",
						"correlation_id", CorrelationID(r.Context()),
						"path", r.URL.Path,
					)
					WriteError(w, r, contracts.New(contracts.CodeInternal, ""))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if v == nil {
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}

type errorBody struct {
	Error struct {
		Code          string `json:"code"`
		Message       string `json:"message"`
		Details       any    `json:"details,omitempty"`
		CorrelationID string `json:"correlation_id"`
	} `json:"error"`
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	appErr := contracts.As(err)
	body := errorBody{}
	body.Error.Code = string(appErr.Code)
	body.Error.Message = appErr.Message
	body.Error.Details = appErr.Details
	body.Error.CorrelationID = CorrelationID(r.Context())
	WriteJSON(w, appErr.HTTPStatus, body)
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any, maxBytes int64) error {
	if maxBytes <= 0 {
		maxBytes = 1 << 20
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			return contracts.Validation("O conteúdo enviado é grande demais.")
		}
		return contracts.Validation("Confira os campos destacados.")
	}
	if dec.More() {
		return contracts.Validation("Envie apenas um objeto JSON.")
	}
	return nil
}

// ClientIP returns the direct peer unless a trusted proxy is configured.
func ClientIP(r *http.Request, trustProxy bool) string {
	if trustProxy {
		if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
			first := strings.TrimSpace(strings.Split(forwarded, ",")[0])
			if net.ParseIP(first) != nil {
				return first
			}
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func IsMutating(method string) bool {
	switch method {
	case http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete:
		return true
	default:
		return false
	}
}
