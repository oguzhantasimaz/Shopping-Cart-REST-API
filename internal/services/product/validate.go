package product_service

import "github.com/oguzhantasimaz/Shopping-Cart-REST-API/internal/models/category"

func CreateProductValidate(r *CreateProductRequest) error {
	if r.Name == "" {
		return ErrProductNameEmpty
	}
	if r.SKU == "" {
		return ErrProductSKUEmpty
	}
	if r.UnitPrice <= 0 {
		return ErrProductPriceInvalid
	}
	if r.Quantity <= 0 {
		return ErrProductQuantityInvalid
	}
	if r.Category == (category.Category{}) {
		return ErrProductCategoryEmpty
	}
	return nil
}

func UpdateProductValidate(r *UpdateProductRequest) error {
	if r.ID == 0 {
		return ErrProductIDEmpty
	}
	if r.Name == "" {
		return ErrProductNameEmpty
	}
	if r.SKU == "" {
		return ErrProductSKUEmpty
	}
	if r.UnitPrice <= 0 {
		return ErrProductPriceInvalid
	}
	if r.Quantity <= 0 {
		return ErrProductQuantityInvalid
	}
	if r.Category == (category.Category{}) {
		return ErrProductCategoryEmpty
	}
	return nil
}

func DeleteProductValidate(r *DeleteProductRequest) error {
	if r.ID == 0 {
		return ErrProductIDEmpty
	}
	return nil
}

func FindProductByIDValidate(r *FindProductByIDRequest) error {
	if r.ID == 0 {
		return ErrProductIDEmpty
	}
	return nil
}
