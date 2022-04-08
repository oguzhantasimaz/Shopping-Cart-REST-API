package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jwtHelper "github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/jwt"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey, os.Getenv("ENV"))
			if decodedClaims != nil {
				for _, role := range decodedClaims.Role {
					if role == "admin" {
						c.Next()
						c.Abort()
						return
					}
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
