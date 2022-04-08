package user

type Repository interface {
	Migration() error
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	FindByID(id int) (*User, error)
	FindAll() (*[]User, error)
	FindByUsername(username string) (*User, error)
}
