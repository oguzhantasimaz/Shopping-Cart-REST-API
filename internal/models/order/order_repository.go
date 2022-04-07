package order

type Repository interface {
	Migration() error
	Create(order *Order) error
	Update(order *Order) error
	Delete(id int) error
	FindByID(id int) (*Order, error)
	FindAllByCustomerID(customerID int) (*[]Order, error)
}
