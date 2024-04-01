package main

import (
	"log"
	"os"
	"user_web_service/internal/adapters/http/http_gin"
	"user_web_service/internal/adapters/repositories"
	"user_web_service/internal/core"
	"user_web_service/internal/infrastructures"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	driver := "mysql"
	DB_DSN := os.Getenv("DB_DSN")
	db := infrastructures.ConnectRDB(driver, DB_DSN)
	defer db.Close()

	userRepository := repositories.NewUserRepositorySQL(db)
	petRepository := repositories.NewPetRepositorySQL(db)
	notificationRepository := repositories.NewNotificationRepositorySQL(db)
	notificationRecordRepository := repositories.NewNotificationRecordRepositorySQL(db)

	notificationRecordService := core.NewNotificationRecordService(notificationRecordRepository)
	notificationService := core.NewNotificationService(notificationRepository, notificationRecordService)
	userService := core.NewUserService(userRepository)
	petService := core.NewPetService(petRepository)

	userHandler := http_gin.NewUserHandler(userService)
	petHandler := http_gin.NewPetHandler(petService)
	notificationHandler := http_gin.NewNotificationHandler(notificationService)

	r := gin.Default()
	r.Use(http_gin.Cors())
	http_gin.UserRoutes(r, userHandler)
	http_gin.PetRoutes(r, petHandler)
	http_gin.NotificationRoutes(r, notificationHandler)
	if err := r.Run(); err != nil {
		log.Fatal("http server error:", err)
	}
}
