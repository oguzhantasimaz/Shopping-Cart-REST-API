package cart_service

func CreateCartValidate(r *CreateCartRequest) error {
	if r.CustomerId <= 0 {
		return ErrCustomerIdRequired
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrProductsRequired
	}
	if r.Products != nil {
		for _, product := range r.Products {
			if product.Id <= 0 {
				return ErrProductIdRequired
			}
			if product.Quantity <= 0 {
				return ErrProductQuantityRequired
			}
		}
	}
	return nil
}

func UpdateCartValidate(r *UpdateCartRequest) error {
	if r.Id <= 0 {
		return ErrCartIdRequired
	}
	if r.CustomerId <= 0 {
		return ErrCustomerIdRequired
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrProductsRequired
	}
	if r.Products != nil {
		for _, product := range r.Products {
			if product.Id <= 0 {
				return ErrProductIdRequired
			}
			if product.Quantity <= 0 {
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
	if r.Id <= 0 {
		return ErrCartIdRequired
	}
	return nil
}

func FindByIdValidate(r *FindByIdRequest) error {
	if r.Id <= 0 {
		return ErrCartIdRequired
	}
	return nil
}
