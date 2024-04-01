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
	db         *sql.DB
	httpClient *http.Client
)

func init() {
	godotenv.Load()

	DB_DSN := os.Getenv("DB_DSN")
	db = infrastructures.ConnectRDB("mysql", DB_DSN)
	httpClient = &http.Client{}
}

func main() {
	defer db.Close()

	reminderRepository := repositories.NewReminderRepositorySQL(db)
	notificationRepository := repositories.NewNotificationRepositorySQL(db)
	petRepository := repositories.NewPetRepositorySQL(db)

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
