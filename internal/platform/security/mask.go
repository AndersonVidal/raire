package security

import (
	"crypto/rand"
	"regexp"
	"strings"
)

func randRead(b []byte) (int, error) { return rand.Read(b) }

// MaskEmail exposes only the first character of the local part.
func MaskEmail(email string) string {
	email = strings.TrimSpace(email)
	at := strings.LastIndex(email, "@")
	if at <= 0 || at == len(email)-1 {
		return "***"
	}
	return email[:1] + "***" + email[at:]
}

// MaskPhone keeps only the last four digits.
func MaskPhone(phone string) string {
	digits := onlyDigits(phone)
	if len(digits) < 4 {
		return "****"
	}
	return "****" + digits[len(digits)-4:]
}

func onlyDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var (
	emailRE  = regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	phoneRE  = regexp.MustCompile(`(?:\+?55\s*)?(?:\(?\d{2}\)?\s*)?\d{4,5}[-\s]?\d{4}`)
	cpfRE    = regexp.MustCompile(`\b\d{3}\.?\d{3}\.?\d{3}-?\d{2}\b`)
	coordsRE = regexp.MustCompile(`-?\d{1,3}\.\d{4,},\s*-?\d{1,3}\.\d{4,}`)
)

// Redact removes common PII patterns from free-form log messages. Structured
// logs should not receive raw PII in the first place; this is defense in depth.
func Redact(s string) string {
	s = emailRE.ReplaceAllString(s, "[email]")
	s = cpfRE.ReplaceAllString(s, "[document]")
	s = coordsRE.ReplaceAllString(s, "[coordinates]")
	s = phoneRE.ReplaceAllString(s, "[phone]")
	return s
}

// Pseudonymize produces a stable, non-reversible actor reference for logs.
func Pseudonymize(value string) string {
	if value == "" {
		return ""
	}
	h := sha256Hex(value)
	return h[:16]
}
