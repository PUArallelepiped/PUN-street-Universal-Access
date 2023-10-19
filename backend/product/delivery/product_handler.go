package delivery

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewStoreHandler(e *gin.Engine, productUsecase domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: productUsecase,
	}
	e.GET("/api/v1/product/:productID", handler.GetProductById)
}

func (s *ProductHandler) GetProductById(c *gin.Context) {
	productID := c.Param("productID")

	product, err := s.ProductUsecase.GetByID(c, productID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	c.JSON(200, &swagger.ProductInfo{
		ProductId:   product.ProductId,
		Name:        product.Name,
		CatogoryId:  product.CatogoryId,
		Description: product.Description,
		Price:       product.Price,
		Storage:     product.Storage,
	})

}
