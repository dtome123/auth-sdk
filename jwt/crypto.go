package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dtome123/auth-sdk/pkgs/crypto"
	"github.com/golang-jwt/jwt/v5"
)

// ======================
// HMAC Signer & Verifier
// ======================

// HMACSigner signs JWT tokens using HS256 (HMAC SHA256).
type HMACSigner struct {
	secret []byte
}

// NewHMACSigner creates a new HMACSigner with the provided secret key.
func NewHMACSigner(secret []byte) *HMACSigner {
	return &HMACSigner{secret: secret}
}

// Sign creates a signed JWT token string using HS256 algorithm.
func (h *HMACSigner) Sign(claims Claims, expiry time.Duration) (string, error) {
	rawClaims := claims.ToRaw()
	rawClaims["exp"] = time.Now().Add(expiry).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rawClaims)
	return token.SignedString(h.secret)
}

// Algorithm returns the JWT signing algorithm used (HS256).
func (h *HMACSigner) Algorithm() string {
	return jwt.SigningMethodHS256.Alg()
}

// HMACVerifier verifies JWT tokens signed with HS256 algorithm.
type HMACVerifier struct {
	secret []byte
}

// NewHMACVerifier creates a new HMACVerifier with the provided secret key.
func NewHMACVerifier(secret []byte) *HMACVerifier {
	return &HMACVerifier{secret: secret}
}

// Verify parses and validates the HS256 signed JWT token.
func (h *HMACVerifier) Verify(tokenStr string) (Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method, require HS256")
		}
		return h.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return NewClaimsFromToken(token)
}

// ======================
// RS256 Signer & Verifier
// ======================

// RS256Signer signs JWT tokens using RS256 (RSA SHA256).
type RS256Signer struct {
	privateKey *rsa.PrivateKey
}

// NewRS256SignerFromKey creates a new RS256Signer from a RSA private key.
func NewRS256SignerFromKey(key *rsa.PrivateKey) *RS256Signer {
	return &RS256Signer{privateKey: key}
}

// NewRS256SignerFromString creates a new RS256Signer from a PEM encoded RSA private key string.
func NewRS256SignerFromString(pemStr string) (*RS256Signer, error) {
	key, err := crypto.LoadRSAPrivateKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256SignerFromKey(key), nil
}

// NewRS256SignerFromPath creates a new RS256Signer from a file containing a PEM encoded RSA private key.
func NewRS256SignerFromPath(path string) (*RS256Signer, error) {
	key, err := crypto.LoadRSAPrivateKeyFromPath(path)
	if err != nil {
		return nil, err
	}
	return NewRS256SignerFromKey(key), nil
}

// Sign creates a signed JWT token string using RS256 algorithm.
func (r *RS256Signer) Sign(claims Claims, expiry time.Duration) (string, error) {
	rawClaims := claims.ToRaw()
	rawClaims["exp"] = time.Now().Add(expiry).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, rawClaims)
	return token.SignedString(r.privateKey)
}

// Algorithm returns the JWT signing algorithm used (RS256).
func (r *RS256Signer) Algorithm() string {
	return jwt.SigningMethodRS256.Alg()
}

// RS256Verifier verifies JWT tokens signed with RS256 algorithm.
type RS256Verifier struct {
	publicKey *rsa.PublicKey
}

// NewRS256VerifierFromKey creates a new RS256Verifier from a RSA public key.
func NewRS256VerifierFromKey(key *rsa.PublicKey) *RS256Verifier {
	return &RS256Verifier{publicKey: key}
}

// NewRS256VerifierFromString creates a new RS256Verifier from a PEM encoded RSA public key string.
func NewRS256VerifierFromString(pemStr string) (*RS256Verifier, error) {
	key, err := crypto.LoadRSAPublicKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256VerifierFromKey(key), nil
}

// NewRS256VerifierFromPath creates a new RS256Verifier from a file containing a PEM encoded RSA public key.
func NewRS256VerifierFromPath(path string) (*RS256Verifier, error) {
	key, err := crypto.LoadRSAPublicKeyFromPath(path)
	if err != nil {
		return nil, err
	}
	return NewRS256VerifierFromKey(key), nil
}

// Verify parses and validates the RS256 signed JWT token.
func (r *RS256Verifier) Verify(tokenStr string) (Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, errors.New("unexpected signing method, require RS256")
		}
		return r.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return NewClaimsFromToken(token)
}

// ======================
// Helper Functions & KeyFuncs
// ======================

// Extract extracts claims from a JWT token string.
// If keys are provided, only those claims are returned.
func Extract(tokenString string, keys ...string) (Claims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format: must have 3 parts")
	}

	payloadSegment := parts[1]

	// Ensure padding for base64 decoding
	if m := len(payloadSegment) % 4; m != 0 {
		payloadSegment += strings.Repeat("=", 4-m)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 payload: %w", err)
	}

	var rawClaims jwt.MapClaims
	if err := json.Unmarshal(payloadBytes, &rawClaims); err != nil {
		return nil, fmt.Errorf("failed to parse JSON payload: %w", err)
	}
	claims := NewClaims(rawClaims)

	if len(keys) == 0 {
		return claims, nil
	}

	extract := make(Claims, len(keys))
	for _, key := range keys {
		if value, exists := claims[key]; exists {
			extract[key] = value
		}
	}

	return extract, nil
}

// RS256KeyFunc returns a jwt.Keyfunc for validating RS256 tokens.
func RS256KeyFunc(pubKey *rsa.PublicKey) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, errors.New("unexpected signing method, expected RS256")
		}
		return pubKey, nil
	}
}

// HS256KeyFunc returns a jwt.Keyfunc for validating HS256 tokens.
func HS256KeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method, expected HS256")
		}
		return secret, nil
	}
}
