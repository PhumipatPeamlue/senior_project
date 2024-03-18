package main

import (
	"database/sql"
	"log"
	"notification_service/internal/adapters/http_gin"
	"notification_service/internal/adapters/repositories"
	"notification_service/internal/core"
	"notification_service/internal/infrastructures"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

func init() {
	godotenv.Load()

	dsn := os.Getenv("NOTIFICATION_DB")
	db = infrastructures.ConnectRDB("mysql", dsn)
}

func main() {
	notificationRepository := repositories.NewNotificationSQL(db)

	notificationService := core.NewNotificationService(notificationRepository)

	notificationHandler := http_gin.NewNotificationHandler(notificationService)

	r := gin.Default()

	http_gin.NotificationRoutes(r, notificationHandler)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
