package domain

import (
	"regexp"
	"strings"
)

var (
	piiEmailRE = regexp.MustCompile(`(?i)[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}`)
	piiPhoneRE = regexp.MustCompile(`(?:\+?55\s*)?(?:\(?\d{2}\)?\s*)?\d{4,5}[-\s]?\d{4}`)
	piiDocRE   = regexp.MustCompile(`\b\d{3}\.?\d{3}\.?\d{3}-?\d{2}\b|\b\d{2}\.?\d{3}\.?\d{3}/?\d{4}-?\d{2}\b`)
	piiCoordRE = regexp.MustCompile(`-?\d{1,3}\.\d{4,}\s*,\s*-?\d{1,3}\.\d{4,}`)
)

// ContainsPII detects contact/document/coordinate patterns in public free text.
// It is deliberately conservative: a false positive is a validation error, not
// a privacy incident.
func ContainsPII(s string) bool {
	return piiEmailRE.MatchString(s) || piiPhoneRE.MatchString(s) ||
		piiDocRE.MatchString(s) || piiCoordRE.MatchString(s)
}

func ValidatePublicText(field, value string, maxRunes int) error {
	value = strings.TrimSpace(value)
	if len([]rune(value)) > maxRunes {
		return &PublicTextError{Field: field, Reason: "too_long"}
	}
	if ContainsPII(value) {
		return &PublicTextError{Field: field, Reason: "pii"}
	}
	return nil
}

type PublicTextError struct {
	Field  string
	Reason string
}

func (e *PublicTextError) Error() string {
	if e.Reason == "pii" {
		return "o campo " + e.Field + " não pode conter e-mail, telefone, documento ou coordenada"
	}
	return "o campo " + e.Field + " está acima do limite"
}
