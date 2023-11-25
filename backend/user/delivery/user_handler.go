package delivery

import (
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
	e.GET("/api/v1/user/get-info/:userID", handler.GetUserById)
	e.GET("/api/v1/admin/get-all-users", handler.GetUsers)
}

func (s *UserHandler) GetUserById(c *gin.Context) {
	UserID := c.Param("userID")
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
