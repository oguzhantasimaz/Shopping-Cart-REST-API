package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/hash"
	jwtHelper "github.com/oguzhantasimaz/Shopping-Cart-REST-API/pkg/jwt"
	"os"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int
	Username  string `gorm:"unique"`
	Iat       int
	Exp       int
	Salt      string
	Hash      string
	Role      string    `gorm:"type:enum('admin', 'user');default:'user'"`
	CreatedAt time.Time `gorm:"<-:create"`
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)

func GetUser(r Repository, username, password string) (*User, error) {
	// find user by username if exists in db check password hash
	user := &User{}
	var err error
	user, err = r.FindByUsername(username)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func ValidateUser(r Repository, username string, password string) (string, error) {
	// find user by username if exists in db check password hash then return accessToken
	user, err := r.FindByUsername(username)
	if err != nil {
		return "", ErrUserNotFound
	}
	if user.Hash != hash.GenerateHash(password, user.Salt) {
		return "", ErrPasswordIncorrect
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"iss":      os.Getenv("ENV"),
		"exp": time.Now().Add(24 *
			time.Hour).Unix(),
		"role": user.Role,
	})

	//TODO: Read secret key from config
	accessToken := jwtHelper.GenerateToken(jwtClaims, "OGUZHANTASIMAZSECRETKEY")
	return accessToken, nil
}

func Create(r Repository, u *User) error {
	return r.Create(u)
}

func Update(r Repository, u *User) error {
	return r.Update(u)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}

func FindByID(r Repository, id int) (*User, error) {
	return r.FindByID(id)
}

func FindAll(r Repository) (*[]User, error) {
	return r.FindAll()
}
