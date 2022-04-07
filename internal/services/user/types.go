package user_service

type CreateUserRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}

type FindUserByIDRequest struct {
	ID int `json:"id"`
}
