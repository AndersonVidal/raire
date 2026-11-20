package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// Cipher encrypts small sensitive values (precise coordinates) with
// AES-256-GCM. The key lives outside the database.
type Cipher struct {
	aead cipher.AEAD
}

func NewCipher(keyBase64 string) (*Cipher, error) {
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, fmt.Errorf("decode encryption key: %w", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes, got %d", len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &Cipher{aead: aead}, nil
}

// EncryptString returns base64(nonce || ciphertext || tag).
func (c *Cipher) EncryptString(plain string) (string, error) {
	if c == nil || c.aead == nil {
		return "", errors.New("cipher not configured")
	}
	nonce := make([]byte, c.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	out := c.aead.Seal(nonce, nonce, []byte(plain), nil)
	return base64.StdEncoding.EncodeToString(out), nil
}

func (c *Cipher) DecryptString(encoded string) (string, error) {
	if c == nil || c.aead == nil {
		return "", errors.New("cipher not configured")
	}
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	ns := c.aead.NonceSize()
	if len(raw) < ns+c.aead.Overhead() {
		return "", errors.New("ciphertext too short")
	}
	nonce, ciphertext := raw[:ns], raw[ns:]
	plain, err := c.aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}

func sha256Sum(b []byte) []byte {
	sum := sha256.Sum256(b)
	return sum[:]
}
