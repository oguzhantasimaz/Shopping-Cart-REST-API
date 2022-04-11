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

func (s *AuthService) Authenticate(req *LoginRequest, secretKey string) (string, error) {
	if err := LoginRequestValidate(req); err != nil {
		return "", err
	}
	accessToken, err := user.ValidateUser(s.repository, req.Username, req.Password, secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *AuthService) Logout(req *LogoutRequest, secretKey string) error {
	if err := LogoutRequestValidate(req); err != nil {
		return err
	}
	return user.Logout(s.repository, req.AccessToken, secretKey)
}

func (s *AuthService) Signup(req *SignupRequest, secretKey string) error {
	if err := SignupRequestValidate(req); err != nil {
		return err
	}
	return user.Signup(s.repository, req.Username, req.Password)
}
