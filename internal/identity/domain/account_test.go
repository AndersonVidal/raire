package domain

import (
	"testing"
	"time"
)

func TestEmailAndPasswordValidation(t *testing.T) {
	if NormalizeEmail("  Ana@Example.TEST ") != "ana@example.test" {
		t.Fatal("normalization failed")
	}
	for _, email := range []string{"ana@example.test", "a.b+tag@example.test"} {
		if err := ValidateEmail(email); err != nil {
			t.Fatalf("%s rejected: %v", email, err)
		}
	}
	for _, email := range []string{"", "ana", "ana@", "@example.test", "ana@example", "a b@example.test"} {
		if err := ValidateEmail(email); err == nil {
			t.Fatalf("%q accepted", email)
		}
	}
	if err := ValidatePassword("senha-segura-123"); err != nil {
		t.Fatal(err)
	}
	for _, password := range []string{"curta", "somentenumeros", "semnumero", "          "} {
		if err := ValidatePassword(password); err == nil {
			t.Fatalf("%q accepted", password)
		}
	}
	if err := ValidateDisplayName("Ana"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateDisplayName("A"); err == nil {
		t.Fatal("short name accepted")
	}
	if err := ValidatePhone("+55 11 99999-8888"); err != nil {
		t.Fatal(err)
	}
	if err := ValidatePhone("123"); err == nil {
		t.Fatal("short phone accepted")
	}
}

func TestAccountDerivedFields(t *testing.T) {
	now := time.Now()
	a := Account{Role: RoleProfessional, ContactPreference: ContactEmail, ContactPhone: "+5511999998888"}
	if !a.Active() {
		t.Fatal("account should be active")
	}
	if a.EmailVerified() {
		t.Fatal("email should not be verified without timestamp")
	}
	a.EmailVerifiedAt = &now
	if !a.EmailVerified() {
		t.Fatal("email should be verified")
	}
	if a.PhoneVerified() {
		t.Fatal("phone should not be verified without timestamp")
	}
	a.ContactPhoneVerifiedAt = &now
	if !a.PhoneVerified() {
		t.Fatal("phone should be verified")
	}
	if a.EffectiveContact() != ContactEmail {
		t.Fatal("effective contact wrong")
	}
	a.SuspendedAt = &now
	if a.Active() {
		t.Fatal("suspended account reported active")
	}
	if MaskEmail("ana@example.test") != "a***@example.test" {
		t.Fatal("email mask")
	}
	if MaskPhone("+55 11 99999-8888") != "****8888" {
		t.Fatal("phone mask")
	}
}
