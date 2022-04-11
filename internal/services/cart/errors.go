package cart_service

import "errors"

var (
	ErrCustomerIDRequired       = errors.New("customer id is required")
	ErrCartIDRequired           = errors.New("cart id is required")
	ErrProductsRequired         = errors.New("products are required")
	ErrProductIDRequired        = errors.New("product id is required")
	ErrProductQuantityRequired  = errors.New("product quantity is required")
	ErrProductUnitPriceRequired = errors.New("product unit price is required")
	ErrInvalidProductID         = errors.New("invalid product id")
	ErrInvalidQuantity          = errors.New("invalid quantity")
)
