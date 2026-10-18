package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckSession returns a Gin middleware that validates an active session for the authenticated user.
func CheckSession[T any](cfg SessionConfig[T]) gin.HandlerFunc {
	userClaim := cfg.UserIDClaim
	if userClaim == "" {
		userClaim = "sub"
	}
	deviceClaim := cfg.DeviceIDClaim
	if deviceClaim == "" {
		deviceClaim = "device_id"
	}
	jtiClaim := cfg.JTIClaim
	if jtiClaim == "" {
		jtiClaim = "jti"
	}

	return func(c *gin.Context) {
		claims, ok := GetGinClaims(c)
		if !ok || claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated: missing claims"})
			return
		}

		userID := claims.Get(userClaim).AsString()
		deviceID := claims.Get(deviceClaim).AsString()
		jti := claims.Get(jtiClaim).AsString()

		if userID == "" || deviceID == "" || jti == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: missing user_id, device_id or jti in claims"})
			return
		}

		sess, err := cfg.Repo.GetSession(c.Request.Context(), userID, deviceID)
		if err != nil || sess == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: session not found or expired"})
			return
		}

		if sess.JTI != jti {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: session revoked or replaced"})
			return
		}

		SetGinSession(c, sess)
		c.Next()
	}
}
