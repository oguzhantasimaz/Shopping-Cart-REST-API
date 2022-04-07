package category_service

func CreateCategoryValidate(r *CreateCategoryRequest) error {
	if r.Name == "" {
		return ErrCategoryNameEmpty
	}
	return nil
}

func CreateCategoryBulkedValidate(r *CreateCategoryBulkedRequest) error {
	if len(r.Categories) == 0 {
		return ErrCategoryEmptyArray
	}
	return nil
}

func UpdateCategoryValidate(r *UpdateCategoryRequest) error {
	if r.ID == 0 {
		return ErrCategoryIDEmpty
	}
	if r.Name == "" {
		return ErrCategoryNameEmpty
	}
	return nil
}

func DeleteCategoryValidate(r *DeleteCategoryRequest) error {
	if r.ID == 0 {
		return ErrCategoryIDEmpty
	}
	return nil
}

func FindByIDCategoryValidate(r *FindByIDCategoryRequest) error {
	if r.ID == 0 {
		return ErrCategoryIDEmpty
	}
	return nil
}
