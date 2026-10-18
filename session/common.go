package session

import "fmt"

// DefaultSessionKeyFormat is the default Redis key format.
const DefaultSessionKeyFormat = "session:user:%s:device:%s"

// BuildSessionKey generates a Redis session key from userID and deviceID.
func BuildSessionKey(userID, deviceID string) string {
	return fmt.Sprintf(DefaultSessionKeyFormat, userID, deviceID)
}

// buildSessionKey is an internal helper retained for backward compatibility.
func buildSessionKey(userID, deviceID string) string {
	return BuildSessionKey(userID, deviceID)
}
