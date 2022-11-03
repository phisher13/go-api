package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/phisher13/go-api/config"
	"github.com/phisher13/go-api/internal/composite/auth"
	"github.com/phisher13/go-api/internal/composite/product"
	"github.com/phisher13/go-api/pkg/client/postgres"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	// init gin engine

	router := gin.Default()

	// init config
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error config initializing, %s", err.Error())
	}

	// parse environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// connect to postgres
	db, err := postgres.PostgresClient(postgres.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("PSQL_PASSWORD"),
		Database: viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error connection to postgres, %s", err.Error())
	}

	authComposite, err := auth.NewAuthorizationComposite(db)
	if err != nil {
		log.Fatalf("error initializing auth composite, %s", err.Error())
	}

	productComposite, err := product.NewProductComposite(db)
	if err != nil {
		log.Fatalf("error initializing product composite, %s", err.Error())
	}

	// init auth router
	authComposite.Handler.InitRoutes(router)

	// init product router
	productComposite.Handler.InitRoutes(router)

	// starting web server
	log.Fatal(router.Run(viper.GetString("server.port")))
}
