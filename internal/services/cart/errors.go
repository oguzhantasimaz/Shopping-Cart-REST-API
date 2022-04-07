package cart_service

import "errors"

var (
	ErrCustomerIdRequired       = errors.New("customer id is required")
	ErrCartIdRequired           = errors.New("cart id is required")
	ErrProductsRequired         = errors.New("products are required")
	ErrProductIdRequired        = errors.New("product id is required")
	ErrProductQuantityRequired  = errors.New("product quantity is required")
	ErrProductUnitPriceRequired = errors.New("product unit price is required")
)
