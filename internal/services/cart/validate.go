package cart_service

func CreateCartValidate(r *CreateCartRequest) error {
	if r.CustomerID <= 0 {
		return ErrCustomerIDRequired
	}
	if r.CartProducts == nil || len(r.CartProducts) == 0 {
		return ErrProductsRequired
	}
	if r.CartProducts != nil {
		for _, p := range r.CartProducts {
			if p.Quantity <= 0 {
				return ErrInvalidQuantity
			}
			if p.ProductID <= 0 {
				return ErrInvalidProductID
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
	if r.CartProducts == nil || len(r.CartProducts) == 0 {
		return ErrProductsRequired
	}
	if r.CartProducts != nil {
		for _, p := range r.CartProducts {
			if p.Quantity <= 0 {
				return ErrInvalidQuantity
			}
			if p.ProductID <= 0 {
				return ErrInvalidProductID
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
