package category_service

import "errors"

var (
	ErrCategoryIDEmpty    = errors.New("category id is empty")
	ErrCategoryNameEmpty  = errors.New("category name is empty")
	ErrCategoryEmptyArray = errors.New("category array is empty")
)
