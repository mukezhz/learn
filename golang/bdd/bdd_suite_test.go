package bdd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bdd Suite")
}
