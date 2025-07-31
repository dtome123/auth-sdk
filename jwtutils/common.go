package jwtutils

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

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

	// Return all claims if no keys specified
	if len(keys) == 0 {
		return claims, nil
	}

	// Filter claims by keys
	extract := make(Claims, len(keys))
	for _, key := range keys {
		if value, exists := claims[key]; exists {
			extract[key] = value
		}
	}

	return extract, nil
}

func RS256KeyFunc(pubKey *rsa.PublicKey) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed with RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, errors.New("unexpected signing method, expected RS256")
		}
		return pubKey, nil
	}
}

func HS256KeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed with HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method, expected HS256")
		}
		return secret, nil
	}
}
