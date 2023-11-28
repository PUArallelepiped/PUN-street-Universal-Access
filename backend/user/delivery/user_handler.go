package delivery

import (
	"strconv"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(e *gin.Engine, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}
	v1 := e.Group("/api/v1")
	{
		v1.GET("/user/get-info/:userID", handler.GetUserById)
		v1.GET("/admin/get-all-users", handler.GetUsers)
	}
}

func (s *UserHandler) GetUserById(c *gin.Context) {
	UserID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	user, err := s.UserUsecase.GetByID(c, UserID)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, user)
}

func (s *UserHandler) GetUsers(c *gin.Context) {
	stores, err := s.UserUsecase.GetAllUser(c)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, stores)
}
