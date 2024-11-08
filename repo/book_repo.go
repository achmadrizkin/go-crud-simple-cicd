package repo

import (
	"go-crud-simple-cicd/domain"
	"go-crud-simple-cicd/model"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) domain.BookRepo {
	return &bookRepo{
		db: db,
	}
}

// CreateBook inserts a new book into the database
func (b *bookRepo) CreateBook(book model.Book) (model.Book, error) {
	if err := b.db.Create(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// GetBookByID retrieves a book by its ID
func (b *bookRepo) GetBookByID(id uint) (model.Book, error) {
	var book model.Book
	if err := b.db.First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

// GetAllBooks retrieves all books
func (b *bookRepo) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

// UpdateBook updates an existing book
func (b *bookRepo) UpdateBook(book model.Book) (model.Book, error) {
	if err := b.db.Save(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

// DeleteBook deletes a book by its ID
func (b *bookRepo) DeleteBook(id uint) error {
	if err := b.db.Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
