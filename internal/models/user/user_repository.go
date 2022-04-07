package user

type Repository interface {
	Migration()
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	FindById(id int) (*User, error)
	FindAll() ([]*User, error)
}
