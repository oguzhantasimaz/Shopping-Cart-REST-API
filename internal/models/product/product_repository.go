package product

type Repository interface {
	Migration()
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
	FindById(id int) (*Product, error)
}
