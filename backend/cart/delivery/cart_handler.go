package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/gorilla/websocket"
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
		v1.GET("/customer/:userID/carts", handler.GetCurrentCarts)
		v1.GET("/customer/:userID/cart/:cartID/store/:storeID/carts", handler.GetHistoryCart)
		v1.GET("/customer/:userID/get-history", handler.GetAllHistory)
		v1.GET("/customer/:userID/order-status", handler.GetRunOrder)

		v1.POST("/customer/:userID/cart", handler.AddProductToCart)
		v1.POST("/customer/:userID/checkout", handler.Checkout)
		v1.DELETE("/customer/:userID/delete/product/:productID", handler.DeleteProduct)

		v1.PUT("/seller/update-order-status/customer/:userID/cart/:cartID/store/:storeID", handler.UpdateOrderStatus)
		v1.GET("/seller/store/:storeID/orders", handler.GetSellerOrders)
	}
	e.GET("/socket", handler.SocketHandler)
}
func (ch *CartHandler) GetAllHistory(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	historyArray, err := ch.CartUsecase.GetAllHistoryById(c, userID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, historyArray)
}

func (ch *CartHandler) GetRunOrder(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	runOrderArray, err := ch.CartUsecase.GetRunOrderByID(c, userID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, runOrderArray)
}

func (ch *CartHandler) DeleteProduct(c *gin.Context) {
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

	err := ch.CartUsecase.DeleteProduct(c, customerID, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (ch *CartHandler) AddProductToCart(c *gin.Context) {
	customerID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	var cartInfo *swagger.CartInfo
	if err := c.BindJSON(&cartInfo); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = ch.CartUsecase.AddProductToCart(c, customerID, cartInfo)
	if err != nil {
		logrus.Error(err)
		if err.Error() == "The inventory is not enough for the supply" {
			c.Status(418)
			return
		}
		c.Status(500)
		return
	}

	c.Status(200)
}

func (ch *CartHandler) GetHistoryCart(c *gin.Context) {
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

	storeOrder, err := ch.CartUsecase.GetHistoryCart(c, customerID, cartID, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, storeOrder)
}

func (ch *CartHandler) GetCurrentCarts(c *gin.Context) {
	customerID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	cartOrder, err := ch.CartUsecase.GetCurrentCartsByUserID(c, customerID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, cartOrder)
}

func (ch *CartHandler) Checkout(c *gin.Context) {
	customerID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = ch.CartUsecase.Checkout(c, customerID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (ch *CartHandler) UpdateOrderStatus(c *gin.Context) {
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

	order, err := ch.CartUsecase.UpdateOrderStatusByID(c, customerID, cartID, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, order)
}

func (ch *CartHandler) GetSellerOrders(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	orders, err := ch.CartUsecase.GetSellerOrders(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, orders)
}

func (ch *CartHandler) SocketHandler(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			panic(err)
		}
	}()

	for {
		// msgType, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: "Echo...",
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(3 * time.Second)
	}
}
