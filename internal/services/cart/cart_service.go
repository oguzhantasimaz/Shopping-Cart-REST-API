package cart_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/cart"

type CartService struct {
	repository cart.Repository
}

func NewCartService(repository cart.Repository) *CartService {
	return &CartService{repository: repository}
}

func (s *CartService) Create(req *CreateCartRequest) error {
	if err := CreateCartValidate(req); err != nil {
		return err
	}
	//map products and sum their prices and equalize them to cart total price
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}

	newCart := &cart.Cart{
		CustomerId: req.CustomerId,
		Products:   req.Products,
		TotalPrice: totalPrice,
	}
	return cart.CreateCart(s.repository, newCart)
}

func (s *CartService) Update(req *UpdateCartRequest) error {
	if err := UpdateCartValidate(req); err != nil {
		return err
	}
	var totalPrice float64
	for _, product := range req.Products {
		totalPrice += product.UnitPrice * float64(product.Quantity)
	}
	newCart := &cart.Cart{
		Id:         req.Id,
		CustomerId: req.CustomerId,
		Products:   req.Products,
		TotalPrice: totalPrice,
	}
	return cart.UpdateCart(s.repository, newCart)
}

func (s *CartService) Delete(req *DeleteCartRequest) error {
	if err := DeleteCartValidate(req); err != nil {
		return err
	}
	return cart.DeleteCart(s.repository, req.Id)
}

func (s *CartService) FindById(req *FindByIdRequest) (*cart.Cart, error) {
	if err := FindByIdValidate(req); err != nil {
		return nil, err
	}
	return cart.FindByIdCart(s.repository, req.Id)
}
