package auth

import "errors"

type User struct {
	Username string
	Password string
}

func NewUser(username, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}

type Auth struct {
	Users []User
}

func NewAuth() Auth {
	return Auth{
		Users: []User{
			NewUser("test", "test"),
			NewUser("user", "user"),
		},
	}
}

func (a *Auth) Login(username, password string) (User, error) {
	for _, u := range a.Users {
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}
	return User{}, errors.New("not found")
}
