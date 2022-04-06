package product

type ProductRepository interface {
	Migration()
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
	FindByID(id int) (*Product, error)
}
