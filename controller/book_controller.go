package controller

import (
	"go-crud-simple-cicd/domain"
	"go-crud-simple-cicd/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	bookUseCase domain.BookUseCase
}

func NewBookController(bookUseCase domain.BookUseCase) *bookController {
	return &bookController{bookUseCase: bookUseCase}
}

func (d *bookController) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Hello world",
	})
}

func (d *bookController) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	bookResponse, err := d.bookUseCase.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Created book success",
		Data:       bookResponse,
	})
}

// GetBookByID handles fetching a book by ID
func (d *bookController) GetBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid book ID",
		})
		return
	}

	book, err := d.bookUseCase.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Fetched book successfully",
		Data:       book,
	})
}

// GetAllBooks handles fetching all books
func (d *bookController) GetAllBooks(c *gin.Context) {
	books, err := d.bookUseCase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Fetched all books successfully",
		Data:       books,
	})
}

// UpdateBook handles updating an existing book
func (d *bookController) UpdateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid book ID",
		})
		return
	}
	book.Id = uint(id)

	updatedBook, err := d.bookUseCase.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated book successfully",
		Data:       updatedBook,
	})
}

// DeleteBook handles deleting a book by ID
func (d *bookController) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid book ID",
		})
		return
	}

	if err := d.bookUseCase.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Deleted book successfully",
	})
}
