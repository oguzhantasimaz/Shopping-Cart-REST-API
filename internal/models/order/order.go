package order

import (
	"time"

	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id         int
	CustomerId int
	Products   []*product.Product
	TotalPrice float64
	Active     bool
	CreatedAt  time.Time `gorm:"<-:create"`
}
