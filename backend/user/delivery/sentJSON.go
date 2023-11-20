package delivery

import (
	"net/http"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/gin-gonic/gin"
)

func TransportData(c *gin.Context) {
	userData := swagger.UserData{
		UserId:    1,
		UserName:  "John Doe",
		UserEmail: "john@example.com",
		Authority: 2,
		Password:  "hashed_password",
		Address:   "123 Main St",
		Phone:     "555-1234",
		Status:    1,
		CartId:    101,
		Birthday:  "1990-01-01",
	}
	c.JSON(http.StatusOK, userData)
}
