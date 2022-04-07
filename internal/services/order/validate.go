package order_service

func CreateOrderValidate(r *CreateOrderRequest) error {
	if r.CustomerId <= 0 {
		return ErrInvalidCustomerId
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrInvalidProducts
	}
	if r.Products != nil {
		for _, p := range r.Products {
			if p.Id <= 0 {
				return ErrInvalidProductId
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
	if r.Id <= 0 {
		return ErrInvalidOrderId
	}
	if r.CustomerId <= 0 {
		return ErrInvalidCustomerId
	}
	if r.Products == nil || len(r.Products) == 0 {
		return ErrInvalidProducts
	}
	if r.Products != nil {
		for _, p := range r.Products {
			if p.Id <= 0 {
				return ErrInvalidProductId
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
	if r.Id <= 0 {
		return ErrInvalidOrderId
	}
	return nil
}

func FindByIdValidate(r *FindByIdRequest) error {
	if r.Id <= 0 {
		return ErrInvalidOrderId
	}
	return nil
}

func FindAllByCustomerIdValidate(r *FindAllByCustomerIdRequest) error {
	if r.CustomerId <= 0 {
		return ErrInvalidCustomerId
	}
	return nil
}
