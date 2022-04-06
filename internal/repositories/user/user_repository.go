package user_repository

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *user.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) FindById(id int) (*user.User, error) {
	u := new(user.User)
	err := ur.db.First(u, id).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepository) Update(user *user.User) error {
	return ur.db.Save(user).Error
}

func (ur *userRepository) Delete(id int) error {
	return ur.db.Where("id = ?", id).Delete(&user.User{}).Error
}

func (ur *userRepository) FindAll() ([]user.User, error) {
	var users []user.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
