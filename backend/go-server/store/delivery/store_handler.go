package delivery

import (
	"main/domain"
	"main/swagger"

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
		Id:      store.ID,
		Name:    store.Name,
		Address: store.Address,
		Email:   store.Email,
		Phone:   store.Phone,
	})

}
