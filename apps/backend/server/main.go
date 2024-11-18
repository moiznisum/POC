package main

import (
	"os"
	"log"
	
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	InitializeRoutes(router)

	router.Run(":" + port)
}


