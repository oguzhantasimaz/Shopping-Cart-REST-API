package user

type User struct {
	Id       int
	Username string
	Password string
	Roles    []string
}

func GetUserList() []*User {
	return []*User{
		{
			Id:       1,
			Username: "admin",
			Password: "1234",
			Roles:    []string{"admin", "customer"},
		},
		{
			Id:       2,
			Username: "customer",
			Password: "12345",
			Roles:    []string{"customer"},
		},
	}
}

func GetUser(username, password string) *User {
	for _, v := range GetUserList() {
		if v.Username == username && v.Password == password {
			return v
		}
	}

	return nil
}
