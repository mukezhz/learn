package book

import (
	"errors"
)

type Book struct {
	Name   string
	Author string
}

func NewBook(name, author string) Book {
	return Book{
		Name:   name,
		Author: author,
	}
}

type BookRepository struct {
	Books []Book
}

func NewBookRepository() BookRepository {
	return BookRepository{
		Books: []Book{},
	}
}

func (br *BookRepository) FetchBooks() []Book {
	return br.Books
}

func (br *BookRepository) FetchBook(name string) (Book, error) {
	for _, b := range br.Books {
		if b.Name == name {
			return b, nil
		}
	}
	return Book{}, errors.New("book not found")
}

func (bs *BookRepository) CreateBook(book Book) (Book, error) {
	b, err := bs.FetchBook(book.Name)
	if err != nil {
		bs.Books = append(bs.Books, book)
		return book, nil
	}
	return b, errors.New("already exists")
}

type BookService struct {
	repository BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return BookService{
		repository: repo,
	}
}

func (bs *BookService) ListBooks() []Book {
	return bs.repository.FetchBooks()
}

func (bs *BookService) GetBook(name string) (Book, error) {
	return bs.repository.FetchBook(name)
}

func (bs *BookService) AddBook(book Book) (Book, error) {
	return bs.repository.CreateBook(book)
}
