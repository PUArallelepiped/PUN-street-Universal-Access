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
	v1 := e.Group("/api/v1")
	{
		v1.POST("/customer/:userID/cart", handler.PostCart)
		v1.GET("/customer/:userID/cart/:cartID/store/:storeID/get-total-price", handler.GetTotalPrice)
		v1.DELETE("/customer/:userID/cart/:cartID/delete/product/:productID", handler.DeleteProduct)
		v1.POST("/customer/:userID/cart/:cartID/store/:storeID/checkout", handler.CheckoutCart)
		v1.GET("/customer/:userID/cart/:cartID/carts", handler.GetCartByCustomerCartID)
		v1.PUT("/customer/:userID/cart/:cartID/update/product/:productID", handler.UpdateCartProductQuantity)
		v1.GET("/customer/:userID/orders", handler.GetOrders)
	}

}

func (s *CartHandler) PostCart(c *gin.Context) {
	customerID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	var cart swagger.CartInfo

	if err := c.BindJSON(&cart); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = s.CartUsecase.PostCart(c, &cart, customerID)

	if err.Error() == "The inventory is not enough for the supply" {
		logrus.Error(err)
		c.Status(418)
		return
	} else if err != nil {
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

func (s *CartHandler) DeleteProduct(c *gin.Context) {
	customerID, customerErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	cartID, cartErr := strconv.ParseInt(c.Param("cartID"), 10, 64)
	productID, productErr := strconv.ParseInt(c.Param("productID"), 10, 64)
	errArr := []error{customerErr, cartErr, productErr}
	for _, err := range errArr {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	err := s.CartUsecase.DeleteProduct(c, customerID, cartID, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *CartHandler) CheckoutCart(c *gin.Context) {
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

	var checkout swagger.CheckoutInfo

	if err := c.BindJSON(&checkout); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	err := s.CartUsecase.Checkout(c, customerID, cartID, storeID, &checkout)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *CartHandler) GetOrders(c *gin.Context) {
	customerID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	orders, err := s.CartUsecase.GetOrderArrayByCustomerID(c, customerID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, orders)
}

func (s *CartHandler) UpdateCartProductQuantity(c *gin.Context) {
	customerID, customerErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	cartID, cartErr := strconv.ParseInt(c.Param("cartID"), 10, 64)
	productID, productErr := strconv.ParseInt(c.Param("productID"), 10, 64)
	errArr := []error{customerErr, cartErr, productErr}
	for _, err := range errArr {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	var cart swagger.CartInfo
	if err := c.BindJSON(&cart); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	productQuantity := cart.ProductQuantity
	err := s.CartUsecase.UpdateProduct(c, customerID, cartID, productID, productQuantity)

	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *CartHandler) GetCartByCustomerCartID(c *gin.Context) {
	customerID, customerErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	cartID, cartErr := strconv.ParseInt(c.Param("cartID"), 10, 64)
	for _, err := range []error{customerErr, cartErr} {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	carts, err := s.CartUsecase.GetCartByCustomerCartID(c, customerID, cartID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, carts)
}
