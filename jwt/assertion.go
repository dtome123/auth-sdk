package jwt

import (
	"time"
	"uuid"
)

// BuildClientAssertionClaims constructs JWT claims for client assertion using OauthClaims struct.
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

// GenerateClientAssertion creates a signed client assertion JWT.
func GenerateClientAssertion(signer Signer, iss string, sub string, aud string, assertionTTL time.Duration) (string, error) {
	claims := BuildClientAssertionClaims(iss, sub, aud, assertionTTL)
	return signer.Sign(claims, assertionTTL)
}
