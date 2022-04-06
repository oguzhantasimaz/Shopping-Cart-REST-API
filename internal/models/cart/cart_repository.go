package cart

type CartRepository interface {
	Migration()
	Create(cart *Cart) error
	Update(cart *Cart) error
	Delete(id int) error
	FindByID(id int) (*Cart, error)
}
