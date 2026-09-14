package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

// HS256KeyFunc returns a jwt.Keyfunc for validating HS256 tokens.
func HS256KeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method, expected HS256")
		}
		return secret, nil
	}
}
