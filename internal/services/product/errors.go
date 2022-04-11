package product_service

import "errors"

var (
	ErrProductNameEmpty       = errors.New("product name is empty")
	ErrProductPriceInvalid    = errors.New("product price is invalid")
	ErrProductSKUEmpty        = errors.New("product sku is empty")
	ErrProductQuantityInvalid = errors.New("product quantity is invalid")
	ErrProductCategoryEmpty   = errors.New("product category is empty")
	ErrProductIDEmpty         = errors.New("product id is empty")
	ErrProductCategoryInvalid = errors.New("product category is invalid")
)
