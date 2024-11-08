package usecase

import (
	"go-crud-simple-cicd/domain"
	"go-crud-simple-cicd/model"
)

type bookUseCase struct {
	bookRepo domain.BookRepo
}

func NewBookUseCase(bookRepo domain.BookRepo) domain.BookUseCase {
	return &bookUseCase{
		bookRepo: bookRepo,
	}
}

// CreateBook implements domain.BookUseCase.
func (b *bookUseCase) CreateBook(book model.Book) (model.Book, error) {
	return b.bookRepo.CreateBook(book)
}

// DeleteBook implements domain.BookUseCase.
func (b *bookUseCase) DeleteBook(id uint) error {
	return b.bookRepo.DeleteBook(id)
}

// GetAllBooks implements domain.BookUseCase.
func (b *bookUseCase) GetAllBooks() ([]model.Book, error) {
	return b.bookRepo.GetAllBooks()
}

// GetBookByID implements domain.BookUseCase.
func (b *bookUseCase) GetBookByID(id uint) (model.Book, error) {
	return b.bookRepo.GetBookByID(id)
}

// UpdateBook implements domain.BookUseCase.
func (b *bookUseCase) UpdateBook(book model.Book) (model.Book, error) {
	return b.bookRepo.UpdateBook(book)
}
