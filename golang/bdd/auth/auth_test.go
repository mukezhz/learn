package auth_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mukezhz/learn/golang/bdd/auth"
)

var _ = Describe("Given: Authorization", func() {
	var (
		username string
		password string
		authObj  auth.Auth
	)

	BeforeEach(func() {
		username = "test"
		password = "test"
		authObj = auth.NewAuth()
	})

	Describe("When: Logging in with invalid credentials", func() {
		It("Then: Should not succeed login and Should not return user object", func() {
			user, err := authObj.Login(username, "password")
			Ω(err).Should(HaveOccurred())
			Ω(user.Username).Should(Equal(""))
		})
	})

	Describe("When: Loggin in with valid credentials", func() {
		It("Then: Should succeed login and Should return user object", func() {
			user, err := authObj.Login(username, password)
			Ω(err).Should(Not(HaveOccurred()))
			Ω(user.Username).Should(Equal(username))
		})
	})
})
