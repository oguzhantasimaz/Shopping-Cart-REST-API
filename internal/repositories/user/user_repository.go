package user_repository

import (
	"fmt"
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

func (ur *userRepository) Migration() error {
	err := ur.db.Migrator().DropTable(&user.User{})
	if err != nil {
		return err
	}
	err = ur.db.AutoMigrate(&user.User{})
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) FindByUsername(username string) (*user.User, error) {
	user := new(user.User)
	err := ur.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Create(user *user.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) FindByID(id int) (*user.User, error) {
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

func (ur *userRepository) FindAll() (*[]user.User, error) {
	var users *[]user.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) InsertSampleData() {
	//username = oguzhantasimaz
	//password = password123
	user := user.User{
		Username: "oguzhantasimaz",
		Salt:     "6241fce458b562ab",
		Hash:     "3tP4urZxQZAJegTtTz0c7gbIqe-RrH9HJltr8l5xnDnIXu5RUk6XjSL7mZAYjp_dwpWwc9Q2t1esxb9KVyZBuw==",
		Role:     "admin",
	}
	err := ur.db.Create(&user)
	fmt.Errorf("%v", err)
}
