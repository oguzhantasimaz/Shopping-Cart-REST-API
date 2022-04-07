package order_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"

type OrderService struct {
	repository order.Repository
}

func NewOrderService(repository order.Repository) *OrderService {
	return &OrderService{repository: repository}
}

func (s *OrderService) Create(req *CreateOrderRequest) error {
	err := CreateOrderValidate(req)
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}
	newOrder := &order.Order{
		CustomerID: req.CustomerID,
		Products:   req.Products,
		TotalPrice: totalPrice,
		Active:     true,
	}
	return order.Create(s.repository, newOrder)
}

func (s *OrderService) Update(req *UpdateOrderRequest) error {
	err := UpdateOrderValidate(req)
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}
	updatedOrder := &order.Order{
		ID:         req.ID,
		CustomerID: req.CustomerID,
		Products:   req.Products,
		TotalPrice: totalPrice,
		Active:     true,
	}

	return order.Update(s.repository, updatedOrder)
}

func (s *OrderService) Delete(req *DeleteOrderRequest) error {
	err := DeleteOrderValidate(req)
	if err != nil {
		return err
	}
	return order.Delete(s.repository, req.ID)
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
