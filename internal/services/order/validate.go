package order_service

func CreateOrderValidate(r *CreateOrderRequest) error {
	if r.CustomerID <= 0 {
		return ErrInvalidCustomerID
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrInvalidProducts
	}
	if r.Products != nil {
		for _, p := range r.Products {
			if p.ID <= 0 {
				return ErrInvalidProductID
			}
			if p.Quantity <= 0 {
				return ErrInvalidQuantity
			}
			if p.UnitPrice <= 0 {
				return ErrInvalidUnitPrice
			}
		}
	}
	return nil
}

func UpdateOrderValidate(r *UpdateOrderRequest) error {
	if r.ID <= 0 {
		return ErrInvalidOrderID
	}
	if r.CustomerID <= 0 {
		return ErrInvalidCustomerID
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrInvalidProducts
	}
	if r.Products != nil {
		for _, p := range r.Products {
			if p.ID <= 0 {
				return ErrInvalidProductID
			}
			if p.Quantity <= 0 {
				return ErrInvalidQuantity
			}
			if p.UnitPrice <= 0 {
				return ErrInvalidUnitPrice
			}
		}
	}
	return nil
}

func DeleteOrderValidate(r *DeleteOrderRequest) error {
	if r.ID <= 0 {
		return ErrInvalidOrderID
	}
	return nil
}

func FindByIDValidate(r *FindByIDRequest) error {
	if r.ID <= 0 {
		return ErrInvalidOrderID
	}
	return nil
}

func FindAllByCustomerIDValidate(r *FindAllByCustomerIDRequest) error {
	if r.CustomerID <= 0 {
		return ErrInvalidCustomerID
	}
	return nil
}
