package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandler(e *gin.Engine, productUsecase domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: productUsecase,
	}
	e.GET("/api/v1/store/:storeID/products", handler.GetProductById)
	e.POST("/api/v1/store/:storeID/add-product", handler.PostProductByStoreId)
	e.PUT("/api/v1/store/:storeID/update-product/:productID", handler.PutByStoreIdProductId)
}

func (s *ProductHandler) GetProductById(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	products, err := s.ProductUsecase.GetByID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	c.JSON(200, products)
}

func (s *ProductHandler) PostProductByStoreId(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	var product swagger.ProductInfo
	if err := c.BindJSON(&product); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = s.ProductUsecase.PostByStoreId(c, storeID, &product)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *ProductHandler) PutByStoreIdProductId(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	productID, err := strconv.ParseInt(c.Param("productID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	var product swagger.ProductInfo
	if err := c.BindJSON(&product); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = s.ProductUsecase.PutById(c, storeID, productID, &product)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
