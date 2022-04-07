package category_service

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateCategoryBulkedRequest struct {
	Categories []CreateCategoryRequest `json:"categories"`
}

type UpdateCategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	ID int `json:"id"`
}

type FindByIDCategoryRequest struct {
	ID int `json:"id"`
}
