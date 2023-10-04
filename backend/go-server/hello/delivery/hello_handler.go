package delivery

import (
	"main/domain"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	HelloUsecase domain.HelloUsecase
}

func NewHelloHandler(e *gin.Engine, helloUsecase domain.HelloUsecase) {
	handler := &HelloHandler{
		HelloUsecase: helloUsecase,
	}
	e.GET("/api/v1/hello", handler.Hello)
}

func (h *HelloHandler) Hello(c *gin.Context) {
	c.JSON(200, "hello")
}
