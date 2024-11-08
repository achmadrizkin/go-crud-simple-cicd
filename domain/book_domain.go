package domain

import "go-crud-simple-cicd/model"

type BookRepo interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBookByID(id uint) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
	UpdateBook(book model.Book) (model.Book, error)
	DeleteBook(id uint) error
}

type BookUseCase interface {
	CreateBook(book model.Book) (model.Book, error)
	GetBookByID(id uint) (model.Book, error)
	GetAllBooks() ([]model.Book, error)
	UpdateBook(book model.Book) (model.Book, error)
	DeleteBook(id uint) error
}
