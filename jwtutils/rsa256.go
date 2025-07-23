package jwtutils

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SignJWTWithRS256 signs a JWT token using an RSA private key and includes an expiration duration.
func SignJWTWithRS256(claims map[string]interface{}, privateKey *rsa.PrivateKey, expiry time.Duration) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for key, value := range claims {
		jwtClaims[key] = value
	}
	jwtClaims["exp"] = time.Now().Add(expiry).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)
	return token.SignedString(privateKey)
}

// VerifyJWTWithRS256 verifies a JWT token string using an RSA public key.
// Returns the token claims as a map if valid.
func VerifyJWTWithRS256(tokenStr string, publicKey *rsa.PublicKey) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
