package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"user_management_service/internal/handler"
	"user_management_service/internal/line"
	"user_management_service/internal/middlewares"
)

const (
	lineClientID     = "2002050743"
	lineClientSecret = "069f5a39b6094506217c75a2b31821a5"
	lineRedirectURI  = "http://localhost:8081/line/callback"
	lineState        = "12345"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			// Handle the case where the file does not exist
			log.Println("No .env file found.")
		} else {
			// Handle other errors
			log.Fatal("Error loading .env file:", err)
		}
	}

	lineClient := line.New(lineClientID, lineClientSecret, lineRedirectURI, lineState)

	h := handler.New(lineClient)

	r := gin.Default()
	r.Use(middlewares.Cors())

	lineGroup := r.Group("/line")
	{
		lineGroup.GET("/login", h.LineLogin())
		lineGroup.GET("/callback", h.LineCallback())
		lineGroup.GET("/profile/:access_token", h.GetProfile())
	}

	if err := r.Run(":8081"); err != nil {
		log.Fatalln(err.Error())
	}
}
