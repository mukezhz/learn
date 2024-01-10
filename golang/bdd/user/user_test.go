package user_test

import (
	"github.com/mukezhz/learn/golang/bdd/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given: User", func() {
	var userService user.UserService
	BeforeEach(func() {
		userService = user.NewUserService()
	})

	Context("When: Empty Users", func() {
		var users []user.User
		BeforeEach(func() {
			users = userService.ListUser()
		})
		It("Then: Should have users length 0", func() {
			users = userService.Users
			Ω(len(users)).Should(Equal(0))
		})
	})

	Context("When: Add user", func() {
		var (
			users        []user.User
			newUser      user.User
			returnedUser user.User
		)
		BeforeEach(func() {
			newUser = user.NewUser("mukezhz", "Mukesh Kumar Chaudhary", "Nepal", "25")
			returnedUser = userService.AddUser(newUser)
			users = userService.ListUser()
		})
		It("Then: Should have users length 1", func() {
			Ω(len(users)).Should(Equal(1))
		})
		It("Then: Should have user equal new user", func() {
			Ω(newUser.Username).Should(Equal(returnedUser.Username))
		})
	})

	Context("When: Get user if invalid username is passed", func() {
		It("Then: Should not fetch user whose username is test", func() {
			returnedUser, err := userService.GetUser("test")
			Ω(returnedUser.Name).Should(Equal(""))
			Ω(err).Should(HaveOccurred())
		})
	})

	Context("When: Get user if valid username is passed", func() {
		It("Then: Should not fetch user whose username is test", func() {
			u := user.NewUser("mukezhz", "Mukesh Kumar Chaudhary", "Nepal", "25")
			_ = userService.AddUser(u)
			returnedUser, err := userService.GetUser(u.Username)
			Ω(returnedUser.Name).Should(Equal(u.Name))
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
