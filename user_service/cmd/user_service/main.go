package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"user_service/internal/adapters/handlers"
	"user_service/internal/adapters/repositories"
	"user_service/internal/conf"
	"user_service/internal/core/services"
	"user_service/internal/middlewares"
	"user_service/internal/mysql"
	"user_service/internal/routes"
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
	userRepo := repositories.NewUserRepositoryMySQL(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.ErrorHandler())

	routes.User(r, userHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatal(err)
	}
}
