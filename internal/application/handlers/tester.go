package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	apperrors "github.com/marveldo/gogin/internal/application/errors"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/validator"
)

type HandlerInterface interface {
	Initialize(r *gin.Engine)
}

type Testhandler struct {
	service services.TesterService
}

func (h *Testhandler) InputMessage(g *gin.Context) {
	var input dto.MessageInput
	err := g.ShouldBindBodyWithJSON(&input)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": &input.Message,
		})
	}
}

func (h *Testhandler) Greet(g *gin.Context) {
	message := h.service.Hello()
	g.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})

}

func (h *Testhandler) Message(g *gin.Context) {
	message := h.service.Message()
	g.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})
}

func (h *Testhandler) GetAllTests(g *gin.Context) {
	tests, err := h.service.GetAllTests()
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   tests,
	})

}

func (h *Testhandler) CreateTest(g *gin.Context) {
	input := dto.TestInput{}
	d_test := domain.TestInput{}

	domain_input := validator.Validate(g, &input, &d_test)
	if domain_input == nil {
		return
	}
	test, err := h.service.CreateTest(domain_input)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   test,
	})
}

func (h *Testhandler) UpdateTest(g *gin.Context) {
	input := dto.TestInputUpdate{}
	d_i := domain.TestInputUpdate{}
	id := g.Param("id")

	domain_input := validator.Validate(g, &input, &d_i)
	if domain_input == nil {
		return
	}
	test, err := h.service.UpdateTest(id, domain_input)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   test,
	})
}

func (h *Testhandler) DeleteTest(g *gin.Context) {
	id := g.Param("id")
	err := h.service.DeleteTest(id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusNoContent, nil)
}

func (h *Testhandler) GetTest(g *gin.Context) {
	id := g.Param("id")
	test, err := h.service.GetTest(id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   test,
	})
}

func (h *Testhandler) Initialize(r *gin.Engine) {
	hg := r.Group("/")
	hg.GET("", h.Greet)
	hg.GET("/message", h.Message)
	hg.GET("/tests", h.GetAllTests)
	hg.POST("/write-message", h.InputMessage)
	hg.POST("/tests", h.CreateTest)
	hg.PUT("/tests/:id", h.UpdateTest)
	hg.DELETE("/tests/:id", h.DeleteTest)
	hg.GET("/tests/:id", h.GetTest)
}

func NewTestHandler(r *gin.Engine, s *services.TesterService) {

	h := &Testhandler{
		service: *s,
	}
	h.Initialize(r)

}
