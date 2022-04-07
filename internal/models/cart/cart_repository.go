package cart

type Repository interface {
	Migration() error
	Create(cart *Cart) error
	Update(cart *Cart) error
	Delete(id int) error
	FindByID(id int) (*Cart, error)
}
