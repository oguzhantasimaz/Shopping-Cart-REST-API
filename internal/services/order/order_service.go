package order_service

import (
	"errors"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
	"time"
)

type OrderService struct {
	repository order.Repository
	proRepo    product.Repository
}

func NewOrderService(repository order.Repository, proRepo product.Repository) *OrderService {
	return &OrderService{repository: repository, proRepo: proRepo}
}

func (s *OrderService) Create(req *CreateOrderRequest) error {
	err := CreateOrderValidate(req)
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, p := range req.OrderProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		totalPrice += product.UnitPrice * float64(p.Quantity)
	}
	newOrder := &order.Order{
		CustomerID:    req.CustomerID,
		OrderProducts: req.OrderProducts,
		TotalPrice:    totalPrice,
		Active:        true,
	}
	return order.Create(s.repository, newOrder)
}

func (s *OrderService) Update(req *UpdateOrderRequest) error {
	err := UpdateOrderValidate(req)
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, p := range req.OrderProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		totalPrice += product.UnitPrice * float64(p.Quantity)
	}
	updatedOrder := &order.Order{
		ID:            req.ID,
		CustomerID:    req.CustomerID,
		OrderProducts: req.OrderProducts,
		TotalPrice:    totalPrice,
		Active:        true,
	}

	return order.Update(s.repository, updatedOrder)
}

func (s *OrderService) Delete(req *DeleteOrderRequest) error {
	err := DeleteOrderValidate(req)
	if err != nil {
		return err
	}
	newOrder, err := order.FindByID(s.repository, req.ID)
	if err != nil {
		return err
	}
	//if its been 14 days since the order was placed, we can delete it
	if newOrder.CreatedAt.AddDate(0, 0, 14).Before(time.Now()) {
		newOrder.Active = false
		return order.Delete(s.repository, req.ID)
	} else {
		return errors.New("cannot delete order that is not older than 14 days")
	}
}

func (s *OrderService) FindByID(req *FindByIDRequest) (*order.Order, error) {
	err := FindByIDValidate(req)
	if err != nil {
		return nil, err
	}
	return order.FindByID(s.repository, req.ID)
}

func (s *OrderService) FindAllByCustomerID(req *FindAllByCustomerIDRequest) (*[]order.Order, error) {
	err := FindAllByCustomerIDValidate(req)
	if err != nil {
		return nil, err
	}
	return order.FindAllByCustomerID(s.repository, req.CustomerID)
}
