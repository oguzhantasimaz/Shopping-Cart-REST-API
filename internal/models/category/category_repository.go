package category

type Repository interface {
	Migration()
	Create(category *Category) error
	CreateBulked(categories []*Category) error
	Update(category *Category) error
	Delete(id int) error
	FindById(id int) (*Category, error)
	FindAll() ([]*Category, error)
}
