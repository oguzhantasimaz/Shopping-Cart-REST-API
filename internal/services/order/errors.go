package order_service

import "errors"

var (
	ErrInvalidCustomerId = errors.New("invalid customer id")
	ErrInvalidProducts   = errors.New("invalid products")
	ErrInvalidProductId  = errors.New("invalid product id")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrInvalidUnitPrice  = errors.New("invalid unit price")
	ErrInvalidOrderId    = errors.New("invalid order id")
)
