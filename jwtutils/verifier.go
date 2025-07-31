package jwtutils

import (
	"crypto/rsa"
	"errors"

	"github.com/dtome123/auth-sdk/pkgs/crypto"
	"github.com/golang-jwt/jwt/v5"
)

// JWTVerifier defines an interface for verifying JWT tokens.
type JWTVerifier interface {
	// Verify parses and validates the JWT token string, returning claims if valid.
	Verify(tokenStr string) (Claims, error)
}

// ======================
// HMAC Verifier
// ======================

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
		// Accept only HS256 signing method
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
// RS256 Verifier
// ======================

// RS256Verifier verifies JWT tokens signed with RS256 algorithm.
type RS256Verifier struct {
	publicKey *rsa.PublicKey
}

// NewRS256VerifierFromKey creates a new RS256Verifier from a RSA public key.
func NewRS256VerifierFromKey(key *rsa.PublicKey) *RS256Verifier {
	return &RS256Verifier{publicKey: key}
}

func NewRS256VerifierFromString(pemStr string) (*RS256Verifier, error) {
	key, err := crypto.LoadRSAPublicKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256VerifierFromKey(key), nil
}

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
		// Accept only RS256 signing method
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
