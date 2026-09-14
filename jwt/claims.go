package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// ======================
// Claims Map Type
// ======================

// Claims represents a set of JWT claims as a map of keys to Values.
type Claims map[string]Value

// NewClaims creates a Claims instance from jwt.MapClaims.
func NewClaims(raw jwt.MapClaims) Claims {
	claims := make(Claims, len(raw))
	for k, v := range raw {
		claims[k] = Value{v}
	}
	return claims
}

// Get returns the Value associated with the given key.
// If the key does not exist, returns a zero Value.
func (c Claims) Get(key string) Value {
	if val, ok := c[key]; ok {
		return val
	}
	return Value{}
}

// Set assigns a Value to the given key.
func (c Claims) Set(key string, val interface{}) {
	c[key] = Value{val}
}

// Remove deletes a key-value pair from the claims.
func (c Claims) Remove(key string) {
	delete(c, key)
}

// Keys returns all keys present in the claims.
func (c Claims) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}

// ToRaw converts Claims back into jwt.MapClaims format.
func (c Claims) ToRaw() jwt.MapClaims {
	raw := make(jwt.MapClaims, len(c))
	for k, v := range c {
		raw[k] = v.Raw()
	}
	return raw
}

// ToString returns JSON string representation of Claims.
func (c Claims) ToString() string {
	raw := c.ToRaw()

	str, err := json.Marshal(raw)
	if err != nil {
		return ""
	}
	return string(str)
}

// ToOauthClaims converts Claims into a typed OauthClaims struct.
func (c Claims) ToOauthClaims() OauthClaims {
	return NewOauthClaims(c)
}

// NewClaimsFromTokenString parses a JWT token string and extracts claims without verification.
// Caller should verify token separately.
func NewClaimsFromTokenString(tokenStr string) (Claims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	return NewClaims(token.Claims.(jwt.MapClaims)), nil
}

// NewClaimsFromToken extracts claims from a parsed JWT token.
func NewClaimsFromToken(token *jwt.Token) (Claims, error) {
	if token == nil || token.Claims == nil {
		return nil, nil
	}
	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		return NewClaims(mapClaims), nil
	}
	return nil, nil
}

// ======================
// Token Parsing & Extraction
// ======================

// Extract extracts claims from a JWT token string without signature verification.
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
