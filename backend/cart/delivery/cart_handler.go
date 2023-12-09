package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartUsecase domain.CartUsecase
}

func NewCartHandler(e *gin.Engine, cartUsecase domain.CartUsecase) {
	handler := &CartHandler{
		CartUsecase: cartUsecase,
	}
	v1 := e.Group("/api/v1")
	{
		v1.GET("/customer/:userID/cart/:cartID/carts")
		v1.GET("/customer/:userID/cart/:cartID/store/:storeID/carts")
		v1.GET("/customer/:userID/cart/:cartID/store/:storeID/get-total-price")
		v1.GET("/customer/:userID/get-history", handler.GetAllHistory)
		v1.GET("/customer/:userID/order-status", handler.GetRunOrder)

		v1.POST("/customer/:userID/cart")
		v1.POST("/customer/:userID/store/:storeID/checkout")
		v1.PUT("/customer/:userID/cart/:cartID/update/product/:productID")
		v1.DELETE("/customer/:userID/delete/product/:productID", handler.DeleteProduct)
	}
}
func (ch *CartHandler) GetAllHistory(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
	}

	historyArray, err := ch.CartUsecase.GetAllHistoryById(c, userID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
	}

	c.JSON(200, historyArray)
}

func (ch *CartHandler) GetRunOrder(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
	}

	runOrderArray, err := ch.CartUsecase.GetRunOrderByID(c, userID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
	}

	c.JSON(200, runOrderArray)
}

func (ch *CartHandler) GetCartsByUserCartID() {

}

func (ch *CartHandler) PostProductToCart(c *gin.Context) {

}

func (s *CartHandler) DeleteProduct(c *gin.Context) {
	customerID, customerErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	productID, productErr := strconv.ParseInt(c.Param("productID"), 10, 64)
	errArr := []error{customerErr, productErr}
	for _, err := range errArr {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	err := s.CartUsecase.DeleteProduct(c, customerID, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
