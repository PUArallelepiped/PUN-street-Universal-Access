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

	v1 := e.Group("/api/v1")
	{
		v1.GET("/product/:productID", handler.GetProductByProductID)
		v1.GET("/store/:storeID/products", handler.GetProductsByStoreId)
		v1.POST("/store/:storeID/add-product", handler.AddProduct)
		v1.PUT("/product/:productID/delete-product", handler.DeleteProduct)
	}
}

func (s *ProductHandler) GetProductByProductID(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("productID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	product, err := s.ProductUsecase.GetProductByID(c, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, product)
}

func (s *ProductHandler) GetProductsByStoreId(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	products, err := s.ProductUsecase.GetProductsByStoreID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	c.JSON(200, products)
}

func (s *ProductHandler) AddProduct(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	var product swagger.ProductInfoWithLabelAndDiscount
	if err := c.BindJSON(&product); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = s.ProductUsecase.AddByStoreId(c, storeID, &product)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}

func (s *ProductHandler) DeleteProduct(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("productID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err = s.ProductUsecase.DeleteProduct(c, productID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
