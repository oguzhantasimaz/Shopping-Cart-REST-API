package user_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/user"

type UserService struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) *UserService {
	return &UserService{repository: repository}
}

//func Create(req CreateUserRequest) error {
//	user := user.User{
//		Username: req.Username,
//		Password: req.Password,
//	}
//	return repository.Create(user)
//}
