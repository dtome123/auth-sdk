package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// OauthClaims represents the standard OAuth 2.0 client assertion claims.
type OauthClaims struct {
	Iss       string   // Issuer - client identifier
	Sub       string   // Subject - same as issuer in client assertions
	Aud       string   // Audience - single primary audience
	Audiences []string // Audiences - list of intended recipients
	Exp       int64    // Expiration time (unix timestamp)
	Iat       int64    // Issued at (unix timestamp)
	Jti       string   // Unique token ID (to prevent replay attacks)
}

// ToClaims converts OauthClaims to the generic Claims type.
func (c OauthClaims) ToClaims() Claims {
	var audVal interface{} = c.Aud
	if len(c.Audiences) > 1 {
		audVal = c.Audiences
	} else if len(c.Audiences) == 1 {
		audVal = c.Audiences[0]
	}

	return Claims{
		"iss": Value{c.Iss},
		"sub": Value{c.Sub},
		"aud": Value{audVal},
		"exp": Value{c.Exp},
		"iat": Value{c.Iat},
		"jti": Value{c.Jti},
	}
}

// NewOauthClaimsFromToken parses a JWT token string and extracts OauthClaims.
func NewOauthClaimsFromToken(tokenStr string, keyFunc jwt.Keyfunc) (OauthClaims, error) {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return OauthClaims{}, err
	}

	if !token.Valid {
		return OauthClaims{}, errors.New("invalid token")
	}

	claims := NewClaims(token.Claims.(jwt.MapClaims))
	return NewOauthClaims(claims), nil
}

// NewOauthClaims takes a map of claims and returns an OauthClaims struct.
func NewOauthClaims(mapClaims Claims) OauthClaims {
	var claims OauthClaims

	claims.Iss = mapClaims.Get("iss").AsString()
	claims.Sub = mapClaims.Get("sub").AsString()
	claims.Exp = mapClaims.Get("exp").AsInt64()
	claims.Iat = mapClaims.Get("iat").AsInt64()
	claims.Jti = mapClaims.Get("jti").AsString()

	audiences := mapClaims.Get("aud").AsStringSlice()
	claims.Audiences = audiences
	if len(audiences) > 0 {
		claims.Aud = audiences[0]
	} else {
		claims.Aud = mapClaims.Get("aud").AsString()
	}

	return claims
}
