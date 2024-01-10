package bdd_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given: AuthService and UserService", func() {
	Describe("When: Nothing", func() {
		It("Then: Should pass", func() {
			Ω(true).Should(Equal(true))
		})
	})
})
