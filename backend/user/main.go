package main

import (
	"database/sql"
	"fmt"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/user/delivery"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/user/repository"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	// connStr := "user=orange dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%s dbname=%s sslmode=disable", "orange", "user_data"))
	if err != nil {
		logrus.Fatal(err)
	}
	UserRepo := repository.NewPostgressqlUserRepo(db)
	UserUsecase := usecase.NewUserUsecase(UserRepo)
	delivery.NewUserHandler(r, UserUsecase)
	if err := r.Run("localhost:7788"); err != nil {
		panic(err)
	}

}
