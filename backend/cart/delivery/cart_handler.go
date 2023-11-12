package delivery

import (
	"strconv"

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
	e.GET("/api/v1/customer/:userID/cart/:cartID/store/:storeID/get-total-price", handler.GetTotalPrice)
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

func (s *CartHandler) GetTotalPrice(c *gin.Context) {
	customerID, customerErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	cartID, cartErr := strconv.ParseInt(c.Param("cartID"), 10, 64)
	storeID, storeErr := strconv.ParseInt(c.Param("storeID"), 10, 64)
	errArr := []error{customerErr, cartErr, storeErr}
	for _, err := range errArr {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	price, err := s.CartUsecase.GetTotalPriceByID(c, customerID, cartID, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"price": price,
	})

}
