package handlers

import (

	
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	apperrors "github.com/marveldo/gogin/internal/application/errors"
	"github.com/marveldo/gogin/internal/application/services"
)



type  HandlerInterface interface {
	Initialize(r * gin.Engine)
}

type Testhandler struct {
	service  services.TesterService
}

func(h *Testhandler) InputMessage(g *gin.Context){
	var input dto.MessageInput
	err := g.ShouldBindBodyWithJSON(&input)
	if err != nil {
		g.JSON(http.StatusBadRequest , gin.H{
			"code" : http.StatusBadRequest,
			"error" : err.Error(),
		})
	}else {
		g.JSON(http.StatusOK, gin.H{
			"code" : http.StatusOK,
			"message" : &input.Message,
		})
	}
}

func (h *Testhandler) Greet(g * gin.Context) {
	message := h.service.Hello()
    g.JSON(http.StatusOK , gin.H{
		"code" : http.StatusOK,
		"message" : message,
	})
	
}

func(h *Testhandler) Message(g * gin.Context){
    message := h.service.Message()
	g.JSON(http.StatusOK , gin.H{
		"code": http.StatusOK ,
		"message" : message,
	})
}

func (h *Testhandler) GetAllTests(g * gin.Context){
	tests , err := h.service.GetAllTests()
	if err != nil {
		g.JSON(http.StatusInternalServerError , gin.H{
			"status" : http.StatusInternalServerError,
			"error" : err,
		})
		return 
	}
	g.JSON(http.StatusOK , gin.H{
		"status" : http.StatusOK,
		"data" : tests,
	})
	

}

func(h *Testhandler) CreateTest(g *gin.Context){
	var input dto.TestInput
	err := g.ShouldBindBodyWithJSON(&input)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"error" : err.Error(),
		})
		return
	}
	domain_input := domain.TestInput(input)
	test , err := h.service.CreateTest(&domain_input)
	if err != nil {
	      if apperrors.CheckDuplicatekeyError(err){
			g.JSON(http.StatusBadRequest, gin.H{
				"status" : http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		  } else {
			g.JSON(http.StatusInternalServerError, gin.H{
				"status" : http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		  }
	}
	g.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"data" : test,
	})
}

func (h *Testhandler) UpdateTest(g *gin.Context) {
	var input dto.TestInputUpdate
    id := g.Param("id")
	if err:= g.ShouldBindBodyWithJSON(&input) ; err!= nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"error" : err.Error(),
		})
		return
	}
	d_i := domain.TestInputUpdate(input)
	test , err := h.service.UpdateTest(id , &d_i)
	if err != nil {
		  if apperrors.CheckDuplicatekeyError(err){
			g.JSON(http.StatusBadRequest, gin.H{
				"status" : http.StatusBadRequest,
				"error":  err.Error(),
			})
			return


		  }else if apperrors.UpdateNotFoundError(err){
            g.JSON(http.StatusNotFound, gin.H{
				"status" : http.StatusNotFound,
				"error":  err.Error(),
			})
			return
		  } else {
			g.JSON(http.StatusInternalServerError, gin.H{
				"status" : http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		  }
	}
   	g.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"data" : test,
	})
}

func (h *Testhandler) Initialize(r *gin.Engine){
	hg := r.Group("/")
	hg.GET("", h.Greet)
	hg.GET("/message", h.Message)
	hg.GET("/tests", h.GetAllTests)
	hg.POST("/write-message", h.InputMessage)
	hg.POST("/tests", h.CreateTest)
	hg.PUT("/tests/:id", h.UpdateTest)
}

func NewTestHandler(r *gin.Engine , s *services.TesterService){
 
 h := &Testhandler{
	service: *s,
 }
 h.Initialize(r)
	
}


