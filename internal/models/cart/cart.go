package cart

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id         int
	CustomerId int
	Products   []*product.Product
	TotalPrice float64
	CreatedAt  time.Time `gorm:"<-:create"`
}
