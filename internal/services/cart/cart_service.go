package cart_service

import (
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/cart"
	"github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/product"
)

type CartService struct {
	repository cart.Repository
	proRepo    product.Repository
}

func NewCartService(repository cart.Repository, proRepo product.Repository) *CartService {
	return &CartService{repository: repository, proRepo: proRepo}
}

func (s *CartService) Create(req *CreateCartRequest) error {
	if err := CreateCartValidate(req); err != nil {
		return err
	}
	//map products and sum their prices and equalize them to cart total price
	var totalPrice float64
	for _, p := range req.CartProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		totalPrice += product.UnitPrice * float64(p.Quantity)
	}

	newCart := &cart.Cart{
		CustomerID:   req.CustomerID,
		CartProducts: req.CartProducts,
		TotalPrice:   totalPrice,
	}
	return cart.Create(s.repository, newCart)
}

func (s *CartService) Update(req *UpdateCartRequest) error {
	if err := UpdateCartValidate(req); err != nil {
		return err
	}
	var totalPrice float64
	for _, p := range req.CartProducts {
		product, err := product.FindByID(s.proRepo, p.ProductID)
		if err != nil {
			return err
		}
		totalPrice += product.UnitPrice * float64(p.Quantity)
	}
	newCart := &cart.Cart{
		ID:           req.ID,
		CustomerID:   req.CustomerID,
		CartProducts: req.CartProducts,
		TotalPrice:   totalPrice,
	}
	return cart.Update(s.repository, newCart)
}

func (s *CartService) Delete(req *DeleteCartRequest) error {
	if err := DeleteCartValidate(req); err != nil {
		return err
	}
	return cart.Delete(s.repository, req.ID)
}

func (s *CartService) FindByID(req *FindByIDRequest) (*cart.Cart, error) {
	if err := FindByIDValidate(req); err != nil {
		return nil, err
	}
	return cart.FindByID(s.repository, req.ID)
}
