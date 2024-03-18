package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"reminder_service/internal/adapters/external_api"
	"reminder_service/internal/adapters/http_gin"
	"reminder_service/internal/adapters/repositories"
	"reminder_service/internal/core"
	"reminder_service/internal/infrastructures"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	db         *sql.DB
	httpClient *http.Client
)

func init() {
	godotenv.Load()

	dsn := os.Getenv("REMINDER_DB")
	db = infrastructures.ConnectRDB("mysql", dsn)
	httpClient = &http.Client{}
}

func main() {
	reminderRepository := repositories.NewReminderRepositorySQL(db)

	notificationService := external_api.NewNotificationService(httpClient)

	reminderService := core.NewReminderService(reminderRepository, notificationService)

	reminderHandler := http_gin.NewReminderHandler(reminderService)

	r := gin.Default()
	// r.Use(http_gin.Cors())

	http_gin.ReminderRoutes(r, reminderHandler)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
