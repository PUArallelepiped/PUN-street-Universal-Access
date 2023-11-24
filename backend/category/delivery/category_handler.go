package delivery

import (
	"strconv"

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
		v1.POST("/store/:storeID/add-category/:categoryID", handler.AddCategoryToStore)
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

func (s *CategoryHandler) AddCategoryToStore(c *gin.Context) {
	storeID, storeErr := strconv.ParseInt(c.Param("storeID"), 10, 64)
	categoryID, categoryErr := strconv.ParseInt(c.Param("categoryID"), 10, 64)
	for _, err := range []error{storeErr, categoryErr} {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	err := s.CategoryUsecase.AddCategoryToStore(c, storeID, categoryID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
