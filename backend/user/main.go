package main

import (
	"github.com/PUArallelepiped/PUN-street-Universal-Access/user/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", api.TransportData)
	r.POST("/user", api.RecieveData)
	if err := r.Run("localhost:7788"); err != nil {
		panic(err)
	}

}
