package jwtutils

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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

// buildClientAssertionClaims constructs JWT claims for client assertion using OauthClaims struct.
func BuildClientAssertionClaims(iss string, sub string, aud string, assertionTTL time.Duration) Claims {
	now := time.Now()
	oauthClaims := OauthClaims{
		Iss: iss,
		Sub: sub,
		Aud: aud,
		Exp: now.Add(assertionTTL).Unix(),
		Iat: now.Unix(),
		Jti: uuid.New().String(),
	}
	return oauthClaims.ToClaims()
}

func GenerateClientAssertion(signer Signer, iss string, sub string, aud string, assertionTTL time.Duration) (string, error) {
	claims := BuildClientAssertionClaims(iss, sub, aud, assertionTTL)
	return signer.Sign(claims, assertionTTL)
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
