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
				if decodedClaims.Exp < jwtHelper.GetCurrentTime() {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
					c.Abort()
					return
				}
				//if path is cart or order check if user is logged in
				if c.Request.URL.Path == "/cart" || c.Request.URL.Path == "/order" {
					if decodedClaims.Role == "admin" || decodedClaims.Role == "user" {
						c.Next()
						c.Abort()
						return
					} else {
						c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
						c.Abort()
						return
					}
				}
				if decodedClaims.Role == "admin" {
					c.Next()
					c.Abort()
					return
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
