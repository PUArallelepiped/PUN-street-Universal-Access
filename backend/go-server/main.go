package main

import (
	"database/sql"
	"fmt"

	_storeDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/store/delivery"
	_storeRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/store/repository"
	_storeUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/store/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Fatal error config file: %v\n", err)
	}
}

func main() {
	logrus.Info("HTTP server started!!!")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	// dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbDatabase),
	)

	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

	storeRepo := _storeRepo.NewPostgressqlStoreRepo(db)

	storeUsecase := _storeUsecase.NewStoreUsecase(storeRepo)

	_storeDelivery.NewStoreHandler(r, storeUsecase)

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
