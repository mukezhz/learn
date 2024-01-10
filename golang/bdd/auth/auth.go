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

type AuthService struct {
	Users []User
}

func NewAuth() AuthService {
	return AuthService{
		Users: []User{},
	}
}

func (a *AuthService) Login(username, password string) (User, error) {
	for _, u := range a.Users {
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}
	return User{}, errors.New("not found")
}

func (a *AuthService) Register(username, password string) (User, error) {
	user := NewUser(username, password)
	for _, u := range a.Users {
		if u.Username == username {
			return User{}, errors.New("user already exists")
		}
	}
	a.Users = append(a.Users, user)
	return user, nil
}
