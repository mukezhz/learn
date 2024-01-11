package buy_book_bdd_test

import (
	"github.com/mukezhz/learn/golang/buy_book_bdd/book"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Feature: CRUD of Book", func() {
	Describe("Scenario: User try to list, get and add a book.", func() {
		Context("Given: No book exists", func() {
			var (
				bookRepository book.BookRepository
				bookService    book.BookService
			)
			BeforeEach(func() {
				bookRepository = book.NewBookRepository()
				bookService = book.NewBookService(bookRepository)
			})
			When("A user try to list a books", func() {
				It("Then: Should have empty books", func() {
					books := bookService.ListBooks()
					Ω(books).Should(BeEmpty())
				})
			})
			When("A user try to get a books", func() {
				bookName := "ikigai"
				book, err := bookService.GetBook(bookName)
				It("Then: Should have an error", func() {
					Ω(err).Should(HaveOccurred())
				})
				It("Then: Should have an empty data", func() {
					Ω(book.Name).Should(BeEmpty())
				})
			})
			When("A user try to add a book", func() {
				It("Then: Should have one book And Error Should Not have occured", func() {
					book, err := bookService.AddBook(book.NewBook("Power of Now", "Now"))
					books := bookService.ListBooks()
					Ω(book.Name).ShouldNot(BeEmpty())
					Ω(len(books)).Should(Equal(1))
					Ω(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Context("Given: 1 book exists", func() {
			var (
				bookRepository book.BookRepository
				bookService    book.BookService
				bookName       string
				newBook        book.Book
			)
			BeforeEach(func() {
				bookRepository = book.NewBookRepository()
				bookService = book.NewBookService(bookRepository)
				bookName = "ikigai"
				newBook, _ = bookService.AddBook(book.NewBook(bookName, "IDK"))
			})
			When("A user try to list a books", func() {
				It("Then: Should Not have empty books", func() {
					books := bookService.ListBooks()
					Ω(books).ShouldNot(BeEmpty())
				})
			})
			When("A user try to get a book", func() {
				It("Then: Should Not have an error And Should have an book with provided book name", func() {
					b, err := bookService.GetBook(bookName)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(b.Name).Should(Equal(newBook.Name))
				})
			})
		})
	})
})
