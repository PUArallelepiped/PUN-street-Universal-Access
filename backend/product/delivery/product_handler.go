package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
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
