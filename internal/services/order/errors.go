package order_service

import "errors"

var (
	ErrInvalidCustomerID = errors.New("invalid customer id")
	ErrInvalidProducts   = errors.New("invalid products")
	ErrInvalidProductID  = errors.New("invalid product id")
	ErrInvalidQuantity   = errors.New("invalid quantity")
	ErrInvalidUnitPrice  = errors.New("invalid unit price")
	ErrInvalidOrderID    = errors.New("invalid order id")
)
