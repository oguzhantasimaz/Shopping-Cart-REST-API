package order_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/order"

type OrderService struct {
	repository order.Repository
}

func NewOrderService(repository order.Repository) *OrderService {
	return &OrderService{repository: repository}
}

func (s *OrderService) Create(req *CreateOrderRequest) error {
	err := CreateOrderValidate(req);
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}
	newOrder := &order.Order{
		CustomerId: req.CustomerId,
		Products:   req.Products,
		TotalPrice: totalPrice,
		Active:     true,
	}
	return order.Create(s.repository, newOrder)
}

func (s *OrderService) Update(req *UpdateOrderRequest) error {
	err := UpdateOrderValidate(req);
	if err != nil {
		return err
	}
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}
	updatedOrder := &order.Order{
		Id:         req.Id,
		CustomerId: req.CustomerId,
		Products:   req.Products,
		TotalPrice: totalPrice,
		Active:     true,
	}

	return order.Update(s.repository, updatedOrder)
}

func (s *OrderService) Delete(req *DeleteOrderRequest) error {
	err := DeleteOrderValidate(req);
	if err != nil {
		return err
	}
	return order.Delete(s.repository, req.Id)
}

func (s *OrderService) FindById(req *FindByIdRequest) (*order.Order, error) {
	err := FindByIdValidate(req);
	if err != nil {
		return nil, err
	}
	return order.FindById(s.repository, req.Id)
}

func (s *OrderService) FindAllByCustomerId(req *FindAllByCustomerIdRequest) ([]*order.Order, error) {
	err := FindAllByCustomerIdValidate(req);
	if err != nil {
		return nil, err
	}
	return order.FindAllByCustomerId(s.repository, req.CustomerId)
}
