package main

import (
	"log"
	"os"

	"backend-noted/config"
	"backend-noted/handler"
	"backend-noted/middleware"
	"backend-noted/repository"
	"backend-noted/service"
	"backend-noted/worker"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Tidak menemukan file .env")
	}

	db := config.SetupDatabase()
	_ = config.SetupRedis()
	noteRepo := repository.NewSqliteNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	trafficRepo := repository.NewSqliteTrafficRepository(db)
	trafficService := service.NewTrafficService(trafficRepo)
	cronWorker := worker.SetupCron(noteRepo)
	defer cronWorker.Stop()
	r := gin.Default()
	r.Use(middleware.TrafficLogger(trafficRepo))
	r.Use(middleware.AuthMiddleware())
	handler.NewNoteHandler(r, noteService)
	handler.NewTrafficHandler(r, trafficService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3678"
	}
	log.Printf("Server berjalan di port :%s", port)
	r.Run(":" + port)
}