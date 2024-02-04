package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reminder_service/internal/adapters/handlers"
	"reminder_service/internal/adapters/repositories"
	"reminder_service/internal/conf"
	"reminder_service/internal/core/services"
	"reminder_service/internal/middlewares"
	"reminder_service/internal/mysql"
	"reminder_service/internal/routes"
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
	notificationRepo := repositories.NewNotificationRepositoryMySQL(db)
	reminderRepo := repositories.NewReminderRepositoryMySQL(db)
	periodReminderRepo := repositories.NewPeriodReminderRepositoryMySQL(reminderRepo, db)
	hourReminderRepo := repositories.NewHourReminderRepositoryMySQL(reminderRepo, db)

	notificationService := services.NewNotificationService(notificationRepo)
	reminderService := services.NewReminderService(reminderRepo, periodReminderRepo, hourReminderRepo, notificationService)

	reminderHandler := handlers.NewReminderHandler(reminderService)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.ErrorHandler())

	routes.Reminder(r, reminderHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatal(err)
	}
}
