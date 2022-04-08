package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/config"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/user"
	auth_service "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/services/auth"
	"net/http"
)

type AuthController struct {
	appConfig *config.Configuration
	service   auth_service.AuthService
}

// @BasePath /auth

func NewAuthController(appConfig *config.Configuration, repository user.Repository) *AuthController {
	return &AuthController{
		appConfig: appConfig,
		service:   *auth_service.NewAuthService(repository),
	}
}

func (c *AuthController) Login(g *gin.Context) {
	request := new(auth_service.LoginRequest)
	if err := g.BindJSON(request); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := c.service.Authenticate(request, c.appConfig.SecretKey)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
