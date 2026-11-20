package domain

import (
	"errors"
	"net/mail"
	"strings"
	"time"
	"unicode"
)

type Role string

const (
	RoleClient       Role = "client"
	RoleProfessional Role = "professional"
	RoleOperator     Role = "operator"
	RoleAdmin        Role = "admin"
)

func (r Role) Valid() bool {
	switch r {
	case RoleClient, RoleProfessional, RoleOperator, RoleAdmin:
		return true
	default:
		return false
	}
}

func (r Role) Privileged() bool { return r == RoleOperator || r == RoleAdmin }

type ContactPreference string

const (
	ContactPlatform ContactPreference = "platform"
	ContactEmail    ContactPreference = "email"
	ContactPhone    ContactPreference = "phone"
)

func (c ContactPreference) Valid() bool {
	switch c {
	case ContactPlatform, ContactEmail, ContactPhone:
		return true
	default:
		return false
	}
}

type ConsentKind string

const (
	ConsentTerms     ConsentKind = "terms"
	ConsentPrivacy   ConsentKind = "privacy"
	ConsentResearch  ConsentKind = "research"
	ConsentMarketing ConsentKind = "marketing"
)

func (c ConsentKind) Valid() bool {
	switch c {
	case ConsentTerms, ConsentPrivacy, ConsentResearch, ConsentMarketing:
		return true
	default:
		return false
	}
}

type Account struct {
	ID                     string
	EmailNormalized        string
	PasswordHash           string
	DisplayName            string
	Role                   Role
	EmailVerifiedAt        *time.Time
	ContactEmail           string
	ContactPhone           string
	ContactPhoneVerifiedAt *time.Time
	ContactPreference      ContactPreference
	MFASecret              string
	MFARecoveryCodesHash   []string
	SuspendedAt            *time.Time
	DeletedAt              *time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func (a *Account) EmailVerified() bool { return a != nil && a.EmailVerifiedAt != nil }
func (a *Account) Active() bool        { return a != nil && a.SuspendedAt == nil && a.DeletedAt == nil }
func (a *Account) Privileged() bool    { return a != nil && a.Role.Privileged() }
func (a *Account) PhoneVerified() bool {
	return a != nil && a.ContactPhone != "" && a.ContactPhoneVerifiedAt != nil
}

func (a *Account) EffectiveContact() ContactPreference {
	if a == nil || !a.ContactPreference.Valid() {
		return ContactPlatform
	}
	return a.ContactPreference
}

type Consent struct {
	ID        string
	AccountID string
	Kind      ConsentKind
	Version   string
	GrantedAt time.Time
	RevokedAt *time.Time
}

type Session struct {
	ID            string
	AccountID     string
	TokenHash     string
	CSRFTokenHash string
	MFAVerifiedAt *time.Time
	CreatedAt     time.Time
	ExpiresAt     time.Time
	RevokedAt     *time.Time
	LastSeenAt    *time.Time
	IPHash        string
	UserAgentHash string
}

func (s *Session) Valid(now time.Time) bool {
	return s != nil && s.RevokedAt == nil && now.Before(s.ExpiresAt)
}

type EmailToken struct {
	ID        string
	AccountID string
	Kind      string // confirm|reset
	TokenHash string
	ExpiresAt time.Time
	UsedAt    *time.Time
	CreatedAt time.Time
}

func (t *EmailToken) Usable(now time.Time) bool {
	return t != nil && t.UsedAt == nil && now.Before(t.ExpiresAt)
}

// NormalizeEmail returns the canonical lookup key. The original value is kept
// for display/contact only.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" || len(email) > 254 || strings.ContainsAny(email, "\r\n\t ") {
		return errors.New("e-mail inválido")
	}
	addr, err := mail.ParseAddress(email)
	if err != nil || NormalizeEmail(addr.Address) != NormalizeEmail(email) {
		return errors.New("e-mail inválido")
	}
	at := strings.LastIndex(email, "@")
	if at <= 0 || at == len(email)-1 || !strings.Contains(email[at+1:], ".") {
		return errors.New("e-mail inválido")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 10 {
		return errors.New("a senha deve ter ao menos 10 caracteres")
	}
	if len(password) > 128 {
		return errors.New("a senha deve ter no máximo 128 caracteres")
	}
	if strings.TrimSpace(password) == "" {
		return errors.New("a senha não pode ser só espaços")
	}
	hasLetter, hasDigit := false, false
	for _, r := range password {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return errors.New("a senha deve combinar letras e números")
	}
	return nil
}

func ValidateDisplayName(name string) error {
	name = strings.TrimSpace(name)
	if len([]rune(name)) < 2 {
		return errors.New("o nome deve ter ao menos 2 caracteres")
	}
	if len([]rune(name)) > 80 {
		return errors.New("o nome deve ter no máximo 80 caracteres")
	}
	if strings.ContainsAny(name, "\r\n\t") {
		return errors.New("o nome contém caracteres inválidos")
	}
	return nil
}

func NormalizePhone(phone string) string {
	var b strings.Builder
	for _, r := range strings.TrimSpace(phone) {
		if r >= '0' && r <= '9' || r == '+' && b.Len() == 0 {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func ValidatePhone(phone string) error {
	digits := 0
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			digits++
		}
	}
	if digits < 10 || digits > 15 {
		return errors.New("telefone inválido")
	}
	return nil
}

func MaskEmail(email string) string {
	at := strings.LastIndex(email, "@")
	if at <= 0 {
		return "***"
	}
	return email[:1] + "***" + email[at:]
}

func MaskPhone(phone string) string {
	digits := NormalizePhone(phone)
	digits = strings.TrimPrefix(digits, "+")
	if len(digits) < 4 {
		return "****"
	}
	return "****" + digits[len(digits)-4:]
}
