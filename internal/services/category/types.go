package category_service

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateCategoryBulkedRequest struct {
	Categories []CreateCategoryRequest `json:"categories"`
}

type UpdateCategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	Id int `json:"id"`
}

type FindByIdCategoryRequest struct {
	Id int `json:"id"`
}
