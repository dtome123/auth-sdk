package jwtutils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ExtractUserID extracts the user ID from a JWT token string by decoding its payload.
// It looks for "sub" or "user_id" claims.
func ExtractUserID(tokenString string, key string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode token payload: %w", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", fmt.Errorf("failed to unmarshal token payload: %w", err)
	}

	if key != "" {
		if userID, ok := claims[key].(string); ok {
			return userID, nil
		}
	}

	if userID, ok := claims["sub"].(string); ok {
		return userID, nil
	}

	return "", errors.New("user ID not found in token claims")
}
