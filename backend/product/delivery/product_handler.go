package delivery

import (
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
	e.POST("/api/v1/products", handler.GetProductById)
}

func (s *ProductHandler) GetProductById(c *gin.Context) {
	productBody := &swagger.ProductsBody{}
	if err := c.BindJSON(productBody); err != nil {
		c.Status(400)
		return
	}
	products, err := s.ProductUsecase.GetByID(c, productBody.StoreId)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	for p := range *products {
		logrus.Info(p)
	}
	c.JSON(200, products)
}
