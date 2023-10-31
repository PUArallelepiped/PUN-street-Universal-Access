package delivery

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CartHandler struct {
	CartUsecase domain.CartUsecase
}

func NewCartHandler(e *gin.Engine, cartUsecase domain.CartUsecase) {
	handler := &CartHandler{
		CartUsecase: cartUsecase,
	}
	e.POST("/api/v1/cart", handler.PostCart)
}

func (s *CartHandler) PostCart(c *gin.Context) {
	var cart swagger.CartInfo

	if err := c.BindJSON(&cart); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err := s.CartUsecase.PostCart(c, &cart)

	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
