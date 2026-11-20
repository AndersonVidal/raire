package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

var b32 = base32.StdEncoding.WithPadding(base32.NoPadding)

// GenerateTOTPSecret returns a 160-bit base32 secret suitable for RFC 6238.
func GenerateTOTPSecret() (string, error) {
	b := make([]byte, 20)
	if _, err := randRead(b); err != nil {
		return "", err
	}
	return b32.EncodeToString(b), nil
}

// GenerateTOTP returns the current 6-digit code for a secret.
func GenerateTOTP(secret string, at time.Time) (string, error) {
	key, err := decodeSecret(secret)
	if err != nil {
		return "", err
	}
	counter := uint64(at.Unix() / 30)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], counter)
	mac := hmac.New(sha1.New, key)
	_, _ = mac.Write(buf[:])
	sum := mac.Sum(nil)
	offset := sum[len(sum)-1] & 0x0f
	value := (uint32(sum[offset])&0x7f)<<24 |
		uint32(sum[offset+1])<<16 |
		uint32(sum[offset+2])<<8 |
		uint32(sum[offset+3])
	return fmt.Sprintf("%06d", value%1_000_000), nil
}

// ValidateTOTP accepts the current step and one step of clock skew in each
// direction, reducing lockouts without allowing replay across a large window.
func ValidateTOTP(secret, code string, at time.Time) bool {
	code = strings.TrimSpace(code)
	if len(code) != 6 {
		return false
	}
	for _, delta := range []int64{-1, 0, 1} {
		want, err := GenerateTOTP(secret, at.Add(time.Duration(delta)*30*time.Second))
		if err != nil {
			return false
		}
		if subtle.ConstantTimeCompare([]byte(want), []byte(code)) == 1 {
			return true
		}
	}
	return false
}

func decodeSecret(secret string) ([]byte, error) {
	s := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(secret), " ", ""))
	s = strings.TrimRight(s, "=")
	key, err := b32.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid TOTP secret: %w", err)
	}
	if len(key) < 10 {
		return nil, fmt.Errorf("TOTP secret too short")
	}
	return key, nil
}
