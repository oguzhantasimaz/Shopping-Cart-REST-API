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
	if r.Id == 0 {
		return ErrCategoryIdEmpty
	}
	if r.Name == "" {
		return ErrCategoryNameEmpty
	}
	return nil
}

func DeleteCategoryValidate(r *DeleteCategoryRequest) error {
	if r.Id == 0 {
		return ErrCategoryIdEmpty
	}
	return nil
}

func FindByIdCategoryValidate(r *FindByIdCategoryRequest) error {
	if r.Id == 0 {
		return ErrCategoryIdEmpty
	}
	return nil
}
