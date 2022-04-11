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

func LogoutRequestValidate(req *LogoutRequest) error {
	if req.AccessToken == "" {
		return ErrTokenRequired
	}
	return nil
}

func SignupRequestValidate(req *SignupRequest) error {
	if req.Username == "" {
		return ErrUsernameRequired
	}
	if req.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}
