package main

import (
	"go-gin-project/config"
	"go-gin-project/model"
	"go-gin-project/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	db := config.DatabaseConnection()
	if db == nil {
		log.Fatal("Database connection failed")
	}

	if err := model.Migration(db); err != nil {
		log.Fatal("Database migration failed:", err)
	}

	router := router.SetupRouter()

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server running on port", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
