package api

import (
	"net/http"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/user/register"
	"github.com/gin-gonic/gin"
)

func RecieveData(c *gin.Context) {
	var req swagger.UserData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	register.UpdateData(req)
	c.JSON(http.StatusOK, req)
}
