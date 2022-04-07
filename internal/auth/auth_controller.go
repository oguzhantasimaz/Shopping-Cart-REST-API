package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/config"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/user"
	jwtHelper "github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/jwt"
)

type AuthController struct {
	appConfig *config.Configuration
}

// @BasePath /auth

func NewAuthController(appConfig *config.Configuration) *AuthController {
	return &AuthController{
		appConfig: appConfig,
	}
}

// Login godoc
// @Summary Login into system with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param loginRequest body LoginRequest true "login informations"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (c *AuthController) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
	}
	user := user.GetUser(req.Username, req.Password)
	if user == nil {
		g.JSON(http.StatusNotFound, gin.H{
			"error_message": "User not found!",
		})
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ENV"),
		"exp": time.Now().Add(24 *
			time.Hour).Unix(),
		"roles": user.Roles,
	})
	token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.JwtSettings.SecretKey)
	g.JSON(http.StatusOK, token)
}

func (c *AuthController) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.JwtSettings.SecretKey, os.Getenv("ENV"))

	g.JSON(http.StatusOK, decodedClaims)

}
