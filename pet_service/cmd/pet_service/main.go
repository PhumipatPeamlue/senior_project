package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pet_service/internal/adapters/handlers"
	"pet_service/internal/adapters/repositories"
	"pet_service/internal/conf"
	"pet_service/internal/core/services"
	"pet_service/internal/middlewares"
	"pet_service/internal/mysql"
	"pet_service/internal/routes"
)

var (
	appConfig *conf.AppConfig
	db        *sql.DB
)

func init() {
	appConfig = conf.LoadConfig()
	db = mysql.ConnectMySQL(appConfig.MySQLDsn)
}

func main() {
	petRepo := repositories.NewPetRepositoryMySQL(db)
	petService := services.NewPetService(petRepo)
	petHandler := handlers.NewPetHandler(petService)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.ErrorHandler())

	routes.Pet(r, petHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatal(err)
	}
}
