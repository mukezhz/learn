package user

import "errors"

type User struct {
	Username string
	Name     string
	Address  string
	Age      string
}

func NewUser(username, name, address, age string) User {
	return User{
		Username: username,
		Name:     name,
		Address:  address,
		Age:      age,
	}
}

type UserService struct {
	Users []User
}

func NewUserService() UserService {
	return UserService{
		Users: []User{},
	}
}

func (us *UserService) ListUser() []User {
	return us.Users
}

func (us *UserService) AddUser(user User) User {
	for _, u := range us.Users {
		if u.Username == user.Username {
			return u
		}
	}
	us.Users = append(us.Users, user)
	return user
}

func (us *UserService) GetUser(username string) (User, error) {
	user := User{}
	for _, u := range us.Users {
		if u.Username == username {
			user = u
			break
		}
	}
	if user.Username != "" {
		return user, nil
	}

	return user, errors.New("user not found")
}

type UserController struct {
	userService UserService
}

func NewUserController(us UserService) *UserController {
	uc := UserController{
		userService: us,
	}
	return &uc
}
