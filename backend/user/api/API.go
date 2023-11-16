package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user = []User{}

func CreateUser(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user = append(user, User{
		Name:      req.Name,
		Password:  req.Password,
		Address:   req.Address,
		Email:     req.Email,
		Birthiday: req.Birthiday,
	})
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var param struct {
		Name string `uri:"name"`
	}
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, u := range user {
		if u.Name == param.Name {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("user %s not found", param.Name),
	})
}
func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, user)
}
func test() {
	r := gin.Default()
	r.GET("/user", ListUsers)
	r.GET("/users/:name", GetUser)
	r.POST("/user", CreateUser)
	if err := r.Run(":7788"); err != nil {
		panic(err)
	}
}
