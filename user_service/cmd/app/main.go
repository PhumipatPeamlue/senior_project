package main

import (
	"log"
	"os"
	"user_service/internal/adapters/http_gin"
	"user_service/internal/adapters/repositories"
	"user_service/internal/core"
	"user_service/internal/infrastructures"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	driver := "mysql"
	userDBDSN := os.Getenv("USER_DB_DSN")
	userDB := infrastructures.ConnectRDB(driver, userDBDSN)
	defer userDB.Close()

	userRepository := repositories.NewUserRepositorySQL(userDB)
	userService := core.NewUserService(userRepository)
	userHandler := http_gin.NewUserHandler(userService)

	r := gin.Default()
	r.Use(http_gin.Cors())
	http_gin.UserRoutes(r, userHandler)
	err := r.Run()
	if err != nil {
		log.Fatal("http server error:", err)
	}
}
