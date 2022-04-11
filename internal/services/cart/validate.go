package cart_service

func CreateCartValidate(r *CreateCartRequest) error {
	if r.CustomerID <= 0 {
		return ErrCustomerIDRequired
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrProductsRequired
	}
	if r.Products != nil {
		for _, product := range r.Products {
			if product.ID <= 0 {
				return ErrProductIDRequired
			}
			if product.Stock <= 0 {
				return ErrProductQuantityRequired
			}
			if product.UnitPrice <= 0 {
				return ErrProductUnitPriceRequired
			}
		}
	}
	return nil
}

func UpdateCartValidate(r *UpdateCartRequest) error {
	if r.ID <= 0 {
		return ErrCartIDRequired
	}
	if r.CustomerID <= 0 {
		return ErrCustomerIDRequired
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrProductsRequired
	}
	if r.Products != nil {
		for _, product := range r.Products {
			if product.ID <= 0 {
				return ErrProductIDRequired
			}
			if product.Stock <= 0 {
				return ErrProductQuantityRequired
			}
			if product.UnitPrice <= 0 {
				return ErrProductUnitPriceRequired
			}
		}
	}
	return nil
}

func DeleteCartValidate(r *DeleteCartRequest) error {
	if r.ID <= 0 {
		return ErrCartIDRequired
	}
	return nil
}

func FindByIDValidate(r *FindByIDRequest) error {
	if r.ID <= 0 {
		return ErrCartIDRequired
	}
	return nil
}
