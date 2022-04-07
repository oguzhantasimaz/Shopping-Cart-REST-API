package order

type Repository interface {
	Migration()
	Create(order *Order) error
	Update(order *Order) error
	Delete(id int) error
	FindById(id int) (*Order, error)
	FindAllByCustomerID(customerID int) ([]*Order, error)
}
