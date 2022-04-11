package product_service

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
	if r.Stock <= 0 {
		return ErrProductQuantityInvalid
	}
	if r.CategoryID <= 0 {
		return ErrProductCategoryInvalid
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
	if r.Stock <= 0 {
		return ErrProductQuantityInvalid
	}
	if r.CategoryID <= 0 {
		return ErrProductCategoryInvalid
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
