package category

type Repository interface {
	Migration() error
	Create(category *Category) error
	CreateBulked(categories *[]Category) error
	Update(category *Category) error
	Delete(id int) error
	FindByID(id int) (*Category, error)
	FindAll() (*[]Category, error)
}
