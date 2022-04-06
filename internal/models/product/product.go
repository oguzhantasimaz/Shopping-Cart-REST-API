package product

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id        int
	Name      string
	SKU       string
	UnitPrice float64
	Quantity  int
	Category  category.Category
	CreatedAt time.Time `gorm:"<-:create"`
}
