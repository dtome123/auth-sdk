package oauthx

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func randString(n int) string {
	b := make([]byte, n)
	_, _ = crand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

// PKCEPair returns (verifier, challenge)
func PKCEPair() (string, string) {
	v := randString(32)
	h := sha256.Sum256([]byte(v))
	c := base64.RawURLEncoding.EncodeToString(h[:])
	return v, c
}

func GenerateState() string { return randString(16) }
func GenerateNonce() string { return randString(16) }

func truncate(s string) string {
	if s == "" {
		return ""
	}
	if len(s) <= 8 {
		return s
	}
	return s[:4] + "…" + s[len(s)-4:]
}
