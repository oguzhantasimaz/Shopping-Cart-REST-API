package category

type CategoryRepository interface {
	Migration()
	Create(category *Category) error
	CreateBulked(categories []*Category) error
	Update(category *Category) error
	Delete(id int) error
	FindByID(id int) (*Category, error)
	FindAll() ([]*Category, error)
}
