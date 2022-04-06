package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id        int
	Name      string
	CreatedAt time.Time `gorm:"<-:create"`
}
