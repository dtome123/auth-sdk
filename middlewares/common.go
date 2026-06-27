package middlewares

import (
	"fmt"

	"github.com/dtome123/auth-sdk/assertion"
	"github.com/dtome123/auth-sdk/jwt"
)

// verifyToken verifies a token string using the verifier and validates claims based on AuthMode.
func verifyToken(
	verifier jwt.JWTVerifier,
	token string,
	expectedAud string,
	replayChecker assertion.ReplayChecker,
	extraValidateOptions []jwt.ValidateOption,
) (jwt.Claims, error) {

	if verifier == nil {
		return nil, fmt.Errorf("jwt verifier is nil")
	}

	claims, err := verifier.Verify(token)
	if err != nil {
		return nil, err
	}

	var opts []jwt.ValidateOption
	if expectedAud != "" {
		opts = append(opts, jwt.WithExpectedAudience(expectedAud))
	}

	if replayChecker != nil {
		opts = append(opts, jwt.WithReplayChecker(replayChecker))
	}

	opts = append(opts, extraValidateOptions...)

	if len(opts) > 0 {
		oauthClaims := jwt.NewOauthClaims(claims)
		if err := oauthClaims.Validate(opts...); err != nil {
			return nil, err
		}
	}

	return claims, nil
}
