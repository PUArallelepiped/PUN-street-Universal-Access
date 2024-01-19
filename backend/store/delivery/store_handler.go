package delivery

import (
	"strconv"

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
	v1 := e.Group("api/v1")
	{
		v1.GET("/store/:storeID", handler.GetStoreById)
		v1.POST("/stores", handler.GetStores)
		v1.GET("/store/:storeID/get-statistics/:year", handler.GetStatistics)
		v1.GET("/store/:storeID/get-selling/:year/:month", handler.GetSelling)
		v1.POST("/customer/:userID/cart/:cartID/store/:storeID/rate", handler.PostRate)
	}
}

func (s *StoreHandler) GetStoreById(c *gin.Context) {
	storeID, err := strconv.ParseInt(c.Param("storeID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	store, err := s.StoreUsecase.GetByID(c, storeID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, store)
}

func (s *StoreHandler) GetStores(c *gin.Context) {
	var searchInfo swagger.SearchInfo
	if err := c.BindJSON(&searchInfo); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	stores, err := s.StoreUsecase.GetAllStore(c, &searchInfo)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, stores)
}

func (s *StoreHandler) GetStatistics(c *gin.Context) {
	storeID, storeErr := strconv.ParseInt(c.Param("storeID"), 10, 64)
	year, yearErr := strconv.ParseInt(c.Param("year"), 10, 64)
	for _, err := range []error{storeErr, yearErr} {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}
	priceArray, err := s.StoreUsecase.GetStatisticsById(c, storeID, year)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, priceArray)
}

func (s *StoreHandler) GetSelling(c *gin.Context) {
	storeID, storeErr := strconv.ParseInt(c.Param("storeID"), 10, 64)
	year, yearErr := strconv.ParseInt(c.Param("year"), 10, 64)
	month, monthErr := strconv.ParseInt(c.Param("month"), 10, 64)
	for _, err := range []error{storeErr, yearErr, monthErr} {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	productStatistic, err := s.StoreUsecase.GetAllProductSellingById(c, storeID, year, month)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, productStatistic)
}

func (s *StoreHandler) PostRate(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("userID"), 10, 64)
	cartID, cartErr := strconv.ParseInt(c.Param("cartID"), 10, 64)
	storeID, storeErr := strconv.ParseInt(c.Param("storeID"), 10, 64)
	for _, err := range []error{storeErr, userErr, cartErr} {
		if err != nil {
			logrus.Error(err)
			c.Status(400)
			return
		}
	}

	var rate swagger.RateInfo
	if err := c.BindJSON(&rate); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	if err := s.StoreUsecase.CalculateRate(c, userID, cartID, storeID, rate); err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.Status(200)
}
