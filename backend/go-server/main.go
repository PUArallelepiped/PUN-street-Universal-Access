package main

import (
	"database/sql"
	_helloDelivery "main/hello/delivery"
	_helloRepo "main/hello/repository"
	_helloUsecase "main/hello/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()

	db := sql.DB{}

	helloRepo := _helloRepo.NewPostgressqlHelloRepo(&db)

	helloUsecase := _helloUsecase.NewHelloUsecase(helloRepo)

	_helloDelivery.NewHelloHandler(r, helloUsecase)

	logrus.Fatal(r.Run())
}
