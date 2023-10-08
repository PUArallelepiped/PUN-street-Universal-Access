package delivery

import (
	"main/domain"

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
	e.GET("/api/v1/stores/:storeID", handler.GetStoreById)
}

func (s *StoreHandler) GetStoreById(c *gin.Context) {
	storeID := c.Param("storeID")

	store, err := s.StoreUsecase.GetByID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, gin.H{"Message": "Internal error"})
		return
	}

	c.JSON(200, gin.H{
		"id":      store.ID,
		"name":    store.Name,
		"address": store.Address,
		"email":   store.Email,
		"phone":   store.Phone,
	})

}
