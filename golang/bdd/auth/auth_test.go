package auth_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mukezhz/learn/golang/bdd/auth"
)

var _ = Describe("Given: Authorization", func() {
	var (
		username    string
		password    string
		authService auth.AuthService
	)

	BeforeEach(func() {
		username = "test"
		password = "test"
		authService = auth.NewAuth()
	})

	Describe("When: Logging in with invalid credentials", func() {
		It("Then: Should not succeed login and Should not return user object", func() {
			user, err := authService.Login(username, "password")
			Ω(err).Should(HaveOccurred())
			Ω(user.Username).Should(Equal(""))
		})
	})

	Describe("When: Registering user with unique username", func() {
		It("Then: Should succeed Auth and Should not have any error", func() {
			user, err := authService.Register(username, password)
			Ω(user.Username).Should(Equal(username))
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("When: Registering user with existing username", func() {
		It("Then: Should Not succeed Auth and Should have error", func() {
			_, err := authService.Register(username, password)
			user, _ := authService.Register(username, password)
			Ω(user.Username).ShouldNot(Equal(username))
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("When: Loggin in with valid credentials", func() {
		var user auth.User
		BeforeEach(func() {
			user, _ = authService.Register(username, password)
		})
		It("Then: Should succeed login and Should return user object", func() {
			u, err := authService.Login(username, password)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(u.Username).Should(Equal(user.Username))
		})
	})

	Describe("When: Loggin in with invalid credentials", func() {
		var user auth.User
		BeforeEach(func() {
			user, _ = authService.Register(username, password)
		})
		It("Then: Should Not Succeed login and Should Not return user object", func() {
			u, err := authService.Login(username, "password")
			fmt.Println(u)
			Ω(u.Password).Should(Equal(user.Password))
			Ω(err).Should(HaveOccurred())
		})
	})
})
