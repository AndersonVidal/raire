package security

import (
	"strings"
	"testing"
	"time"
)

func testParams() Params {
	return Params{Memory: 8 * 1024, Iterations: 1, Parallelism: 1, SaltLength: 16, KeyLength: 32}
}

func TestHashPasswordRoundTrip(t *testing.T) {
	hash, err := HashPassword("senha-segura-123", testParams())
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(hash, "senha-segura-123") {
		t.Fatal("hash contains plaintext password")
	}
	ok, err := VerifyPassword(hash, "senha-segura-123")
	if err != nil || !ok {
		t.Fatalf("expected match, ok=%v err=%v", ok, err)
	}
	ok, err = VerifyPassword(hash, "senha-errada-123")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("wrong password matched")
	}
	if _, err := VerifyPassword("not-a-hash", "x"); err == nil {
		t.Fatal("expected malformed hash error")
	}
}

func TestTOTP(t *testing.T) {
	secret, err := GenerateTOTPSecret()
	if err != nil {
		t.Fatal(err)
	}
	now := time.Unix(1_700_000_000, 0).UTC()
	code, err := GenerateTOTP(secret, now)
	if err != nil {
		t.Fatal(err)
	}
	if len(code) != 6 {
		t.Fatalf("code length = %d", len(code))
	}
	if !ValidateTOTP(secret, code, now) {
		t.Fatal("current code should validate")
	}
	if !ValidateTOTP(secret, code, now.Add(30*time.Second)) {
		t.Fatal("one-step skew should validate")
	}
	if ValidateTOTP(secret, "000000", now) {
		t.Fatal("unexpected code accepted")
	}
	if ValidateTOTP(secret, code, now.Add(5*time.Minute)) {
		t.Fatal("stale code accepted")
	}
}

func TestCipher(t *testing.T) {
	c, err := NewCipher("MDEyMzQ1Njc4OWFiY2RlZjAxMjM0NTY3ODlhYmNkZWY=")
	if err != nil {
		t.Fatal(err)
	}
	enc, err := c.EncryptString("-23.550520,-46.633308")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(enc, "23.55") {
		t.Fatal("ciphertext contains plaintext")
	}
	dec, err := c.DecryptString(enc)
	if err != nil || dec != "-23.550520,-46.633308" {
		t.Fatalf("decrypt = %q, err=%v", dec, err)
	}
	if _, err := c.DecryptString(enc + "AAAA"); err == nil {
		t.Fatal("tampered ciphertext accepted")
	}
}

func TestRedact(t *testing.T) {
	in := "contato ana@example.test tel 11 99999-8888 cpf 123.456.789-00 coord -23.5505,-46.6333"
	out := Redact(in)
	for _, forbidden := range []string{"ana@example.test", "99999-8888", "123.456.789-00", "-23.5505,-46.6333"} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("redaction left %q in %q", forbidden, out)
		}
	}
}

func TestMask(t *testing.T) {
	if got := MaskEmail("ana@example.test"); got != "a***@example.test" {
		t.Fatalf("MaskEmail = %q", got)
	}
	if got := MaskPhone("+55 11 99999-8888"); got != "****8888" {
		t.Fatalf("MaskPhone = %q", got)
	}
	if got := MaskPhone("123"); got != "****" {
		t.Fatalf("MaskPhone short = %q", got)
	}
}
