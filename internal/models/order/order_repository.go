package order

type OrderRepository interface {
	Migration()
	Create(order *Order) error
	Update(order *Order) error
	Delete(id int) error
	FindByID(id int) (*Order, error)
	FindAllByCustomerID(customerID int) ([]*Order, error)
}
