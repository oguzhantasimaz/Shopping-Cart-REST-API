package order_repository

import (
	"fmt"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Migration() error {
	err := r.db.Migrator().DropTable(&order.Order{})
	if err != nil {
		return err
	}
	err = r.db.AutoMigrate(&order.Order{})
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) Create(o *order.Order) error {
	fmt.Println(o)
	return r.db.Create(o).Error
}

func (r *orderRepository) FindAllByCustomerID(customerID int) (*[]order.Order, error) {
	var orders []order.Order
	err := r.db.Preload("Products").Where("customer_id = ?", customerID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *orderRepository) FindByID(id int) (*order.Order, error) {
	o := new(order.Order)
	// get order by id and associated order products with it
	err := r.db.Preload("Products").Where("id = ?", id).First(o).Error

	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *orderRepository) Update(o *order.Order) error {
	return r.db.Save(o).Error
}

func (r *orderRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&order.Order{}).Error
}
