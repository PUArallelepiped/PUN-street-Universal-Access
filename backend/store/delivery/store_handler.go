package delivery

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StoreHandler struct {
	StoreUsecase domain.StoreUsecase
}

func NewStoreHandler(e *gin.Engine, storeUsecase domain.StoreUsecase) {
	handler := &StoreHandler{
		StoreUsecase: storeUsecase,
	}
	e.GET("/api/v1/store/:storeID", handler.GetStoreById)
	e.GET("/api/v1/stores", handler.GetAllStore)
}

func (s *StoreHandler) GetStoreById(c *gin.Context) {
	storeID := c.Param("storeID")

	store, err := s.StoreUsecase.GetByID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}

	c.JSON(200, &swagger.StoreInfo{
		StoreId: store.ID,
		Name:    store.Name,
		Address: store.Address,
		Email:   store.Email,
		Phone:   store.Phone,
	})

}

func (s *StoreHandler) GetAllStore(c *gin.Context) {
	stores, err := s.StoreUsecase.GetAll(c)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal Error :(",
		})
		return
	}
	c.JSON(200, stores)
}
