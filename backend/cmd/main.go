package main

import (
	"database/sql"
	"fmt"
	"os"

	_storeDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/store/delivery"
	_storeRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/store/repository"
	_storeUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/store/usecase"

	_userDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/user/delivery"
	_userRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/user/repository"
	_userUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/user/usecase"
	"github.com/gin-contrib/cors"

	_productDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/product/delivery"
	_productRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/product/repository"
	_productUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/product/usecase"

	_cartDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/cart/delivery"
	_cartRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/cart/repository"
	_cartUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/cart/usecase"

	_discountDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/discount/delivery"
	_discountRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/discount/repository"
	_discountUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/discount/usecase"

	_categoryDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/category/delivery"
	_categoryRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/category/repository"
	_categotyUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/category/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func HandlHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}
}

func main() {
	logrus.SetReportCaller(true)

	logrus.Info("HTTP server started!!!")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPassword := viper.GetString("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	// if go not run in docker, host will be localhost
	if dbHost == "" {
		dbHost = "localhost"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)

	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default(), HandlHeaders())
	storeRepo := _storeRepo.NewPostgressqlStoreRepo(db)
	storeUsecase := _storeUsecase.NewStoreUsecase(storeRepo)
	_storeDelivery.NewStoreHandler(r, storeUsecase)

	productRepo := _productRepo.NewPostgressqlProductRepo(db)
	productUsecase := _productUsecase.NewProductUsecase(productRepo)
	_productDelivery.NewProductHandler(r, productUsecase)

	cartRepo := _cartRepo.NewPostgressqlCartRepo(db)
	cartUsecase := _cartUsecase.NewCartUsecase(cartRepo)
	_cartDelivery.NewCartHandler(r, cartUsecase)

	UserRepo := _userRepo.NewPostgressqlUserRepo(db)
	UserUsecase := _userUsecase.NewUserUsecase(UserRepo)
	_userDelivery.NewUserHandler(r, UserUsecase)

	discountRepo := _discountRepo.NewPostgressqlDiscountRepo(db)
	discountUsecase := _discountUsecase.NewDiscountUsecase(discountRepo)
	_discountDelivery.NewDiscountHandler(r, discountUsecase)

	categoryRepo := _categoryRepo.NewPostgressqlCategoryRepo(db)
	categoryUsecase := _categotyUsecase.NewCategoryUsecase(categoryRepo)
	_categoryDelivery.NewCategoryHandler(r, categoryUsecase)

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
