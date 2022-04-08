package auth_service

func LoginRequestValidate(req *LoginRequest) error {
	if req.Username == "" {
		return ErrUsernameRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}
