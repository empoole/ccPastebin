package main

import (
	"ccPasteBinServer/controller"
	"ccPasteBinServer/database"
	"ccPasteBinServer/migrations"
	"ccPasteBinServer/model"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	migrations.Migrate()
	serveApplication()
}

func loadEnv() {
  err := godotenv.Load("../.env.local")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func loadDatabase() {
  database.Connect()
  database.DBConnection.AutoMigrate(&model.Note{})
}


func serveApplication() {
  router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.POST("/save", controller.Save)


  router.Run(":8000")
  fmt.Println("Server running on port 8000")
}