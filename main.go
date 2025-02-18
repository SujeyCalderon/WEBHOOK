package main

import (
	"os"
	infrastructure "pull-request-checker/src/infrastructure/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.POST("/webhook", infrastructure.GithubWebhookHanlder)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
