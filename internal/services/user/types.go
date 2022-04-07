package user_service

type CreateUserRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type FindUserByIdRequest struct {
	Id int `json:"id"`
}
