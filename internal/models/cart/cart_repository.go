package cart

type Repository interface {
	Migration()
	Create(cart *Cart) error
	Update(cart *Cart) error
	Delete(id int) error
	FindById(id int) (*Cart, error)
}
