package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

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
		v1.DELETE("/customer/:userID/cart/:cartID/delete/product/:productID")
	}
}
func (ch *CartHandler) GetAllHistory(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.Status(400)
	}

	historyArray, err := ch.CartUsecase.GetAllHistoryById(c, userID)
	if err != nil {
		c.Status(500)
	}

	c.JSON(200, historyArray)
}

func (ch *CartHandler) GetRunOrder(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.Status(400)
	}

	runOrderArray, err := ch.CartUsecase.GetRunOrderByID(c, userID)
	if err != nil {
		c.Status(500)
	}

	c.JSON(200, runOrderArray)
}
