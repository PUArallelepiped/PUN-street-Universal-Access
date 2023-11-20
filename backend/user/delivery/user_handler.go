package delivery

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
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
		v1.POST("/login", handler.Login)
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var user swagger.UserData

	if err := c.BindJSON(&user); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	token, err := u.UserUsecase.Login(c, user.UserEmail, user.Password)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(200, "Login Success")
}
