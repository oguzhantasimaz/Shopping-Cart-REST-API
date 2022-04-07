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
	if r.ID <= 0 {
		return ErrUserIDRequired
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
	if r.ID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}

func FindUserByIDValidate(r FindUserByIDRequest) error {
	if r.ID <= 0 {
		return ErrUserIDRequired
	}
	return nil
}
