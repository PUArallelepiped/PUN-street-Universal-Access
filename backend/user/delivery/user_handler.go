package delivery

import (
	"net/http"
	"time"

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
		v1.GET("/validate", handler.ValidateToken)
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var user swagger.LoginInfo

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
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("jwttoken", token, (int)(24*time.Hour), "/", "localhost", false, true)

	c.JSON(200, "Login Success")
}

func (u *UserHandler) ValidateToken(c *gin.Context) {
	token, err := c.Cookie("jwttoken")
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}
	err = u.UserUsecase.ValidateToken(c, token)
	if err != nil {
		logrus.Error(err)
		c.Status(http.StatusUnauthorized)
		return
	}
	c.JSON(200, "Validate Success")
}
