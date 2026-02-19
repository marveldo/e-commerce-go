package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	apperrors "github.com/marveldo/gogin/internal/application/errors"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/validator"
)

type Bookhandler struct {
	r *services.BookService
}

func (h *Bookhandler) CreateBook(g *gin.Context) {
	book_input := dto.BookInputdto{}
	book_domain := domain.BookInput{}
	result := validator.Validate(g, &book_input, &book_domain)
	if result == nil {
		return
	}
	data, err := h.r.CreateBook(result)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(201, gin.H{
		"code":    201,
		"message": "Book created",
		"data":    data,
	})
}

func (h *Bookhandler) GetAllBooks(g *gin.Context) {
	query  := dto.BookQueryDto{}
	query_domain := domain.GetBookQuery{}
	result := validator.ValidateQuery(g, &query , &query_domain)
	if result == nil {
		return
	}
	books, err := h.r.FindAllBooks(result)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(200, gin.H{
		"code":    200,
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

func (h *Bookhandler) DeleteBook(g *gin.Context) {
	id := g.Param("id")
	err := h.r.DeleteBook(id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(204, nil)


}

func (h *Bookhandler) GetBookById(g *gin.Context) {
	id := g.Param("id")
	book, err := h.r.GetBookById(id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(200, gin.H{
		"code":    200,
		"message": "Book retrieved successfully",
		"data":    book,
	})
}
func (b *Bookhandler) Initialize(r *gin.Engine) {
	h := r.Group("/api/v1/book")
	h.POST("", b.CreateBook)
	h.GET("", b.GetAllBooks)
	h.DELETE("/:id", b.DeleteBook)
	h.GET("/:id", b.GetBookById)
}

func NewBookHandler(r *gin.Engine, s *services.BookService) {
	handler := &Bookhandler{
		r: s,
	}
	handler.Initialize(r)
}
