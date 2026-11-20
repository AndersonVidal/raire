package config

import (
	"encoding/base64"
	"testing"
)

func TestLoadDevDefaults(t *testing.T) {
	t.Setenv("RAIRE_ENV", "dev")
	t.Setenv("RAIRE_ENCRYPTION_KEY", "")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.DevMode || cfg.EncryptionKey == "" {
		t.Fatalf("cfg = %+v", cfg)
	}
	if _, err := base64.StdEncoding.DecodeString(cfg.EncryptionKey); err != nil {
		t.Fatal(err)
	}
}

func TestLoadProductionRequiresKey(t *testing.T) {
	t.Setenv("RAIRE_ENV", "prod")
	t.Setenv("RAIRE_ENCRYPTION_KEY", "")
	if _, err := Load(); err == nil {
		t.Fatal("production without key accepted")
	}
	t.Setenv("RAIRE_ENCRYPTION_KEY", "not-base64")
	if _, err := Load(); err == nil {
		t.Fatal("invalid key accepted")
	}
}

func TestLoadDefaultsToProductionSafe(t *testing.T) {
	t.Setenv("RAIRE_ENV", "")
	t.Setenv("RAIRE_ENCRYPTION_KEY", "")
	if _, err := Load(); err == nil {
		t.Fatal("missing RAIRE_ENV should default to production-safe behavior and fail without key")
	}
}
