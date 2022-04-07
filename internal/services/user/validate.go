package user_service

func CreateUserValidate(r CreateUserRequest) error {
	if r.Username == "" {
		return ErrUserNameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func UpdateUserValidate(r UpdateUserRequest) error {
	if r.Id <= 0 {
		return ErrUserIdRequired
	}
	if r.Username == "" {
		return ErrUserNameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func DeleteUserValidate(r DeleteUserRequest) error {
	if r.Id <= 0 {
		return ErrUserIdRequired
	}
	return nil
}

func FindUserByIdValidate(r FindUserByIdRequest) error {
	if r.Id <= 0 {
		return ErrUserIdRequired
	}
	return nil
}
