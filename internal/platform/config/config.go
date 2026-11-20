package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Argon2 holds the password hashing parameters. Production defaults follow
// OWASP guidance for Argon2id; tests may lower them for speed.
type Argon2 struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// Config is the runtime configuration of the foundation binary.
type Config struct {
	Env                string
	Addr               string
	DBPath             string
	SecureCookies      bool
	DevMode            bool
	TrustProxy         bool
	SessionTTL         time.Duration
	ConfirmTokenTTL    time.Duration
	ResetTokenTTL      time.Duration
	TermsVersion       string
	PrivacyVersion     string
	EncryptionKey      string
	MaxBodyBytes       int64
	Argon2             Argon2
	RateLimitEnabled   bool
	PreciseLocationTTL time.Duration
}

const devEncryptionKey = "MDEyMzQ1Njc4OWFiY2RlZjAxMjM0NTY3ODlhYmNkZWY="

// Load reads configuration from environment variables, applying safe defaults
// for local development. Production requires an explicit encryption key.
func Load() (Config, error) {
	env := strings.ToLower(strings.TrimSpace(getenv("RAIRE_ENV", "prod")))
	dev := env == "dev" || env == "test"

	cfg := Config{
		Env:                env,
		Addr:               getenv("RAIRE_ADDR", ":8080"),
		DBPath:             getenv("RAIRE_DB", "raire.db"),
		SecureCookies:      getbool("RAIRE_SECURE_COOKIES", !dev),
		DevMode:            dev,
		TrustProxy:         getbool("RAIRE_TRUST_PROXY", false),
		SessionTTL:         24 * time.Hour,
		ConfirmTokenTTL:    24 * time.Hour,
		ResetTokenTTL:      time.Hour,
		TermsVersion:       getenv("RAIRE_TERMS_VERSION", "2026-11-20"),
		PrivacyVersion:     getenv("RAIRE_PRIVACY_VERSION", "2026-11-20"),
		EncryptionKey:      getenv("RAIRE_ENCRYPTION_KEY", ""),
		MaxBodyBytes:       1 << 20,
		RateLimitEnabled:   getbool("RAIRE_RATE_LIMIT_ENABLED", true),
		PreciseLocationTTL: 90 * 24 * time.Hour,
		Argon2: Argon2{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
	}

	if cfg.EncryptionKey == "" {
		if !dev {
			return Config{}, fmt.Errorf("RAIRE_ENCRYPTION_KEY is required outside dev/test")
		}
		cfg.EncryptionKey = devEncryptionKey
	}
	key, err := base64.RawStdEncoding.DecodeString(strings.TrimRight(cfg.EncryptionKey, "="))
	if err != nil {
		key, err = base64.StdEncoding.DecodeString(cfg.EncryptionKey)
	}
	if err != nil {
		return Config{}, fmt.Errorf("RAIRE_ENCRYPTION_KEY must be base64: %w", err)
	}
	if len(key) != 32 {
		return Config{}, fmt.Errorf("RAIRE_ENCRYPTION_KEY must decode to 32 bytes, got %d", len(key))
	}
	cfg.EncryptionKey = base64.StdEncoding.EncodeToString(key)

	if cfg.MaxBodyBytes <= 0 {
		cfg.MaxBodyBytes = 1 << 20
	}
	return cfg, nil
}

func getenv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func getbool(key string, fallback bool) bool {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}
