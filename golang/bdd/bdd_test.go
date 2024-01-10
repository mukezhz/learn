package bdd_test

import (
	"github.com/mukezhz/learn/golang/bdd/auth"
	"github.com/mukezhz/learn/golang/bdd/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given: AuthService and UserService", func() {
	var (
		authService  auth.AuthService
		userService  user.UserService
		registerUser auth.User
		userDetail   user.User
		username     string
		password     string
	)

	BeforeEach(func() {
		username = "test"
		password = "test"
		registerUser, _ = authService.Register(username, password)
		registerUser, _ = authService.Login(username, password)
		_ = registerUser
		_ = userDetail
	})

	Describe("When: Nothing", func() {
		It("Then: Should pass", func() {
			Ω(true).Should(Equal(true))
		})
	})

	Describe("When: User is logged User with valid credential", func() {
		It("Then: Should add user detail", func() {
			u := userService.AddUser(user.NewUser(username, "Test User", "test address", "100"))
			Ω(u.Username).Should(Equal(username))
		})
	})
})
