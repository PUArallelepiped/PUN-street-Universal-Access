package delivery

import (
	"main/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	name, err := h.HelloUsecase.SayHello(c)
	if err != nil {
		logrus.Error(err)
	}
	c.JSON(200, gin.H{"name": name.Name})
}
