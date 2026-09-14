package jwt

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/dtome123/auth-sdk/pkgs/crypto"
	"github.com/golang-jwt/jwt/v5"
)

// RS256Signer signs JWT tokens using RS256 (RSA SHA256).
type RS256Signer struct {
	privateKey *rsa.PrivateKey
}

// NewRS256SignerFromKey creates a new RS256Signer from a RSA private key.
func NewRS256SignerFromKey(key *rsa.PrivateKey) *RS256Signer {
	return &RS256Signer{privateKey: key}
}

// NewRS256SignerFromString creates a new RS256Signer from a PEM encoded RSA private key string.
func NewRS256SignerFromString(pemStr string) (*RS256Signer, error) {
	key, err := crypto.LoadRSAPrivateKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256SignerFromKey(key), nil
}

// NewRS256SignerFromPath creates a new RS256Signer from a file containing a PEM encoded RSA private key.
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

// RS256Verifier verifies JWT tokens signed with RS256 algorithm.
type RS256Verifier struct {
	publicKey *rsa.PublicKey
}

// NewRS256VerifierFromKey creates a new RS256Verifier from a RSA public key.
func NewRS256VerifierFromKey(key *rsa.PublicKey) *RS256Verifier {
	return &RS256Verifier{publicKey: key}
}

// NewRS256VerifierFromString creates a new RS256Verifier from a PEM encoded RSA public key string.
func NewRS256VerifierFromString(pemStr string) (*RS256Verifier, error) {
	key, err := crypto.LoadRSAPublicKeyFromString(pemStr)
	if err != nil {
		return nil, err
	}
	return NewRS256VerifierFromKey(key), nil
}

// NewRS256VerifierFromPath creates a new RS256Verifier from a file containing a PEM encoded RSA public key.
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

// RS256KeyFunc returns a jwt.Keyfunc for validating RS256 tokens.
func RS256KeyFunc(pubKey *rsa.PublicKey) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, errors.New("unexpected signing method, expected RS256")
		}
		return pubKey, nil
	}
}
