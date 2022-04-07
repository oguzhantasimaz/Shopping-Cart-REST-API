package user

import (
	"crypto"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int
	Username  string
	Iat       int
	Exp       int
	Salt      string
	Hash      string
	Roles     []string
	CreatedAt time.Time `gorm:"<-:create"`
}

func GetUserList() []*User {
	return []*User{
		{
			Id:       1,
			Username: "admin",
			Iat:      int(time.Now().Unix()),
			Exp:      int(time.Now().Unix() + 3600),
			Salt:     "206161894b3957ad",
			Hash:     "e9f26691865bd07efcb42f3f4b66589ba175ff0ccad4b03e5655c1129dc1ad7fcdc6b94a83153424993c4377237c1fb910e6a9cfd12a96ee4eff544ae2dabfc4",
			Roles:    []string{"admin", "customer"},
		},
	}
}

func GetUser(username, password string) *User {
	for _, v := range GetUserList() {
		if v.Username == username && v.Hash == string(crypto.SHA512.New().Sum([]byte(v.Salt+password))) {
			return v
		}
	}

	return nil
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

func FindById(r Repository, id int) (*User, error) {
	return r.FindById(id)
}

func FindAll(r Repository) ([]*User, error) {
	return r.FindAll()
}
