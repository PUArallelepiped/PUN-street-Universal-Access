package main

import (
	"database/sql"
	"fmt"

	_storeDelivery "github.com/PUArallelepiped/PUN-street-Universal-Access/store/delivery"
	_storeRepo "github.com/PUArallelepiped/PUN-street-Universal-Access/store/repository"
	_storeUsecase "github.com/PUArallelepiped/PUN-street-Universal-Access/store/usecase"
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

func main() {
	logrus.Info("HTTP server started!!!")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPassword := viper.GetString("POSTGRES_PASSWORD")

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
	r.Use(cors.Default())
	storeRepo := _storeRepo.NewPostgressqlStoreRepo(db)
	storeUsecase := _storeUsecase.NewStoreUsecase(storeRepo)
	_storeDelivery.NewStoreHandler(r, storeUsecase)

	productRepo := _productRepo.NewPostgressqlProductRepo(db)
	productUsecase := _productUsecase.NewProductUsecase(productRepo)
	_productDelivery.NewProductHandler(r, productUsecase)

	cartRepo := _cartRepo.NewPostgressqlCartRepo(db)
	cartUsecase := _cartUsecase.NewCartUsecase(cartRepo)
	_cartDelivery.NewCartHandler(r, cartUsecase)

	discountRepo := _discountRepo.NewPostgressqlDiscountRepo(db)
	discountUsecase := _discountUsecase.NewDiscountUsecase(discountRepo)
	_discountDelivery.NewDiscountHandler(r, discountUsecase)

	categoryRepo := _categoryRepo.NewPostgressqlCategoryRepo(db)
	categoryUsecase := _categotyUsecase.NewCategoryUsecase(categoryRepo)
	_categoryDelivery.NewCategoryHandler(r, categoryUsecase)

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
