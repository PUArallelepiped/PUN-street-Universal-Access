package delivery

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(e *gin.Engine, categoryUsecase domain.CategoryUsecase) {
	handler := &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}

	v1 := e.Group("/api/v1")
	{
		v1.GET("/categorys", handler.GetCategory)
	}
}

func (s *CategoryHandler) GetCategory(c *gin.Context) {
	categorys, err := s.CategoryUsecase.GetAllCategory(c)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, categorys)
}
