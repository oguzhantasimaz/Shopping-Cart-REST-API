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

func (ur *userRepository) FindAll() ([]user.User, error) {
	var users []user.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) InsertSampleData() {
	roles := []string{"admin", "user"}
	//username = oguzhantasimaz
	//password = password123
	user := user.User{
		Username: "oguzhantasimaz",
		Salt:     "6241fce458b562ab",
		Hash:     "921bbc0931d1bd48d87c8ac267cf7dd4fe52589396b0aded3cf0c0cd4cc4319d6f460b331fce7466ca9f8eee55ba42894943e19da713b9973f0e5127a5cfede1",
		Roles:    roles,
	}
	ur.db.Create(&user)
}
