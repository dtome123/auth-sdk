package middlewares

import (
	"github.com/dtome123/auth-sdk/jwtutils"
)

// verifyServiceToken verifies a JWT token and returns claims
// and applies replay checks for service tokens.
func verifyServiceToken(
	verifier jwtutils.JWTVerifier,
	token string,
	expectedAud string,
	replayChecker jwtutils.ReplayChecker,
) (jwtutils.Claims, error) {
	claims, err := verifier.Verify(token)
	if err != nil {
		return nil, err
	}

	oauthClaims := jwtutils.NewOauthClaims(claims)

	err = oauthClaims.Validate(
		jwtutils.WithExpectedAudience(expectedAud),
		jwtutils.WithReplayChecker(replayChecker),
	)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func verifyUserToken(
	verifier jwtutils.JWTVerifier,
	token string,
	expectedAud string,
) (jwtutils.Claims, error) {
	claims, err := verifier.Verify(token)
	if err != nil {
		return nil, err
	}

	oauthClaims := jwtutils.NewOauthClaims(claims)

	err = oauthClaims.Validate(
		jwtutils.WithExpectedAudience(expectedAud),
	)

	if err != nil {
		return nil, err
	}

	return claims, nil
}
