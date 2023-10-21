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
	e.GET("/api/v1/cart/:customerID/:productID/:storeID", handler.GetCartById)
	e.POST("/api/v1/cart", handler.PostCart)
}

func (s *CartHandler) GetCartById(c *gin.Context) {
	customerID := c.Param("customerID")
	productID := c.Param("productID")
	storeID := c.Param("storeID")

	cart, err := s.CartUsecase.GetByID(c, customerID, productID, storeID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	c.JSON(200, &swagger.CartInfo{
		CustomerId:      cart.CustomerId,
		ProductId:       cart.ProductId,
		StoreId:         cart.StoreId,
		ProductQuantity: cart.ProductQuantity,
	})
}

func (s *CartHandler) PostCart(c *gin.Context) {
	var cart domain.Cart

	if err := c.BindJSON(&cart); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :( ",
		})
		return
	}

	_, err := s.CartUsecase.PostCart(c, &cart)

	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :( ",
		})
		return
	}

	c.Status(200)
}
