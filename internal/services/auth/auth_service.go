package auth_service

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/user"
)

type AuthService struct {
	repository user.Repository
}

func NewAuthService(repository user.Repository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Authenticate(req *LoginRequest) (string, error) {
	if err := LoginRequestValidate(req); err != nil {
		return "", err
	}
	accessToken, err := user.ValidateUser(s.repository, req.Username, req.Password)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
