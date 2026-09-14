package token

import (
	"context"
	"errors"
	"strings"

	"github.com/dtome123/auth-sdk/jwt"
)

func IdentityToken(ctx context.Context) (*jwt.OauthClaims, error) {

	authorizationHeader := ctx.Value("authorization").(string)
	parts := strings.SplitN(authorizationHeader, " ", 2)
	if len(parts) < 2 {
		return nil, errors.New("invalid authorization header")
	}
	clientAssertion := parts[1]

	claims, err := jwt.Extract(clientAssertion)
	if err != nil {
		return nil, errors.New("invalid client assertion: " + err.Error())
	}
	oauthClaims := claims.ToOauthClaims()

	return &oauthClaims, nil
}
