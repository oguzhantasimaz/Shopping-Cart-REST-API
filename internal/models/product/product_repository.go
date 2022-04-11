package product

type Repository interface {
	Migration() error
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
	FindByID(id int) (*Product, error)
	FindAll() ([]*Product, error)
}
