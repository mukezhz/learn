package buy_book_bdd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuyBookBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuyBookBdd Suite")
}
