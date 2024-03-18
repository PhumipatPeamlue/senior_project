package main

import (
	"context"
	"cronjob/internal/adapters/repositories"
	"cronjob/internal/core"
	"cronjob/internal/infrastructures"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

var (
	reminderDB     *sql.DB
	notificationDB *sql.DB
	petDB          *sql.DB
	httpClient     *http.Client
)

func init() {
	godotenv.Load()

	reminderDbDsn := os.Getenv("REMINDER_DB")
	reminderDB = infrastructures.ConnectRDB("mysql", reminderDbDsn)

	notificationDbDsn := os.Getenv("NOTIFICATION_DB")
	notificationDB = infrastructures.ConnectRDB("mysql", notificationDbDsn)

	petDbDsn := os.Getenv("PET_DB")
	petDB = infrastructures.ConnectRDB("mysql", petDbDsn)

	httpClient = &http.Client{}
}

func main() {
	defer reminderDB.Close()
	defer notificationDB.Close()
	defer petDB.Close()

	reminderRepository := repositories.NewReminderRepositorySQL(reminderDB)
	notificationRepository := repositories.NewNotificationRepositorySQL(notificationDB)
	petRepository := repositories.NewPetRepositorySQL(petDB)

	lineNotificationService := core.NewLineNotificationService(httpClient)

	service := core.NewService(reminderRepository, notificationRepository, petRepository, lineNotificationService)

	c := cron.New()

	c.AddFunc("0 * * * *", func() {
		log.Println("Run send notification job")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := service.SendNotification(ctx)
		if err != nil {
			log.Println(err)
		}
	})

	c.AddFunc("@midnight", func() {
		log.Println("Run renew notification job")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := service.Renew(ctx)
		if err != nil {
			log.Println(err)
		}
	})

	c.Start()

	select {}
}
