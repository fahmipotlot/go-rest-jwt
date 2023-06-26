package main

import (
	"diary_go_api/controller"
	"diary_go_api/database"
	"diary_go_api/middleware"
	"diary_go_api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/api")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api/entry")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/", controller.AddEntry)
	protectedRoutes.GET("/", controller.GetAllEntries)
	protectedRoutes.GET("/:id", controller.GetDetailEntry)
	protectedRoutes.PUT("/", controller.UpdateEntry)
	protectedRoutes.DELETE("/:id", controller.DeletEntry)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
