package user

type UserRepository interface {
	Migration()
	Create(user *User) UserRepository
	Update(user *User) error
	Delete(id int) error
	FindByID(id int) (*User, error)
	FindAll() ([]*User, error)
}
