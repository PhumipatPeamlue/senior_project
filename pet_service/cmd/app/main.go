package main

import (
	"log"
	"os"
	"pet_service/internal/adapters/http_gin"
	"pet_service/internal/adapters/repositories"
	"pet_service/internal/core"
	"pet_service/internal/infrastructures"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	driver := "mysql"
	petDBDSN := os.Getenv("PET_DB_DSN")
	petDB := infrastructures.ConnectRDB(driver, petDBDSN)
	defer petDB.Close()

	userRepository := repositories.NewPetRepositorySQL(petDB)

	userService := core.NewPetService(userRepository)

	userHandler := http_gin.NewPetHandler(userService)

	r := gin.Default()
	r.Use(http_gin.Cors())

	http_gin.PetRoutes(r, userHandler)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
