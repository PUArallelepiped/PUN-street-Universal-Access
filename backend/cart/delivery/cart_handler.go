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
	e.GET("/api/v1/cart/:customerID", handler.GetCartByCustomerId)
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
		CustomerID:      cart.CustomerID,
		ProductID:       cart.ProductID,
		StoreID:         cart.StoreID,
		ProductQuantity: cart.ProductQuantity,
	})
}

func (s *CartHandler) GetCartByCustomerId(c *gin.Context) {
	customerID := c.Param("customerID")

	carts, err := s.CartUsecase.GetByCustomerId(c, customerID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	jsonData := &[]swagger.CartInfo{}
	for _, cart := range *carts {
		*jsonData = append(*jsonData, swagger.CartInfo{
			CustomerID:      cart.CustomerID,
			ProductID:       cart.ProductID,
			StoreID:         cart.StoreID,
			ProductQuantity: cart.ProductQuantity,
		})
	}

	c.JSON(200, jsonData)
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

	_cart, err := s.CartUsecase.PostCart(c, &cart)

	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :( ",
		})
		return
	}

	c.JSON(200, &swagger.CartInfo{
		CustomerID:      _cart.CustomerID,
		ProductID:       _cart.ProductID,
		StoreID:         _cart.StoreID,
		ProductQuantity: _cart.ProductQuantity,
	})
}
