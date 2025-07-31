package jwtutils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// OauthClaims represents the standard OAuth 2.0 client assertion claims.
type OauthClaims struct {
	Iss string // Issuer - client identifier
	Sub string // Subject - same as issuer in client assertions
	Aud string // Audience - the intended recipient (token endpoint)
	Exp int64  // Expiration time (unix timestamp)
	Iat int64  // Issued at (unix timestamp)
	Jti string // Unique token ID (to prevent replay attacks)
}

// Validate runs all validators including custom ones with the provided options.
func (c OauthClaims) Validate(opts ...ValidateOption) error {
	options := &validateOptions{
		timeNow: time.Now,
	}
	for _, opt := range opts {
		opt(options)
	}

	validators := []validateFunc{
		validateRequiredFields,
		validateIssuerSubject,
		validateExpiration,
		validateAudience,
		validateReplay,
	}

	validators = append(validators, options.customValidators...)

	for _, v := range validators {
		if err := v(c, options); err != nil {
			return err
		}
	}
	return nil
}

// ToClaims converts OauthClaims to the generic Claims type,
// which is used to sign a JWT.
func (c OauthClaims) ToClaims() Claims {
	return Claims{
		"iss": Value{c.Iss},
		"sub": Value{c.Sub},
		"aud": Value{c.Aud},
		"exp": Value{c.Exp},
		"iat": Value{c.Iat},
		"jti": Value{c.Jti},
	}
}

// NewOauthClaimsFromToken parses a JWT token string and extracts OAuthClaims.
// The keyFunc is provided for token validation (if nil, only parsing is performed).
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
	claims.Aud = mapClaims.Get("aud").AsString()

	return claims
}
