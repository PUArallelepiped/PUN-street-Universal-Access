package delivery

import (
	"net/http"
	"strconv"

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
		v1.GET("/user/get-info/:userID", handler.GetUserById)
		v1.GET("/admin/get-all-users", handler.GetUsers)
		v1.POST("/login", handler.Login)
		v1.GET("/validate", handler.ValidateToken)
		v1.POST("/register", handler.RegisterUser)
		v1.POST("/check-email", handler.CheckEmail)
		v1.GET("/userID", handler.GetUserIdByCookie)
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

func (u *UserHandler) Login(c *gin.Context) {
	var user swagger.LoginInfo

	if err := c.BindJSON(&user); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	token, err := u.UserUsecase.Login(c, user.UserEmail, user.Password)
	if err != nil {
		if err.Error() == "banned" {
			c.JSON(403, "banned")
			return
		}
		logrus.Error(err)
		c.Status(500)
		return
	}
	// fmt.Println(c.Cookie("jwttoken"))
	// c.SetSameSite(http.SameSiteStrictMode)
	// c.SetCookie("jwttoken", token, (int)(24*time.Hour), "/", "localhost", false, true)

	// c.JSON(200, "Login Success")
	c.JSON(200, token)
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

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var user swagger.RegisterInfo

	if err := c.BindJSON(&user); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}

	err := u.UserUsecase.RegisterUser(c, &user)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, "Register Success")
}

func (u *UserHandler) CheckEmail(c *gin.Context) {
	var email swagger.EmailInfo

	if err := c.BindJSON(&email); err != nil {
		logrus.Error(err)
		c.Status(400)
		return
	}
	if email.UserEmail == "" {
		c.JSON(200, false)
		return
	}
	isExist, err := u.UserUsecase.CheckEmail(c, email.UserEmail)
	if err != nil {
		logrus.Error(err)
		c.Status(500)
		return
	}

	c.JSON(200, isExist)
}

func (u *UserHandler) GetUserIdByCookie(c *gin.Context) {
	token, err := c.Cookie("jwttoken")
	if err != nil {
		logrus.Error(err)
		c.Status(500)
	}
	id, err := u.UserUsecase.GetUserIdByCookie(c, token)
	idInfo := swagger.IdInfo{
		UserId: id,
	}
	if err != nil {
		logrus.Error(err)
		c.Status(500)
	}
	c.JSON(200, idInfo)
}
