package jwtutils

import (
	"crypto/rsa"
	"time"

	"github.com/dtome123/auth-sdk/pkgs/crypto"
	"github.com/golang-jwt/jwt/v5"
)

// JWTSigner defines an interface for signing JWT tokens.
type JWTSigner interface {
	// Sign generates a signed JWT string with the given claims and expiry duration.
	Sign(claims Claims, expiry time.Duration) (string, error)
	// Algorithm returns the signing algorithm name, e.g. "HS256", "RS256".
	Algorithm() string
}

// ======================
// HMAC Signer
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

// ======================
// RS256 Signer
// ======================

// RS256Signer signs JWT tokens using RS256 (RSA SHA256).
type RS256Signer struct {
	privateKey *rsa.PrivateKey
}

// NewRS256SignerFromKey creates a new RS256Signer from a RSA private key.
func NewRS256SignerFromKey(key *rsa.PrivateKey) *RS256Signer {
	return &RS256Signer{privateKey: key}
}

func NewRS256SignerFromString(pemStr string) (*RS256Signer, error) {
	key, err := crypto.LoadRSAPrivateKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256SignerFromKey(key), nil
}

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
