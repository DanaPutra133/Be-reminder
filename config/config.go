package config

import (
	"backend-noted/domain"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "noted.db"
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke Database SQLite:", err)
	}

	db.AutoMigrate(&domain.Note{}, &domain.TrafficStat{})
	log.Println("Database SQLite berhasil terhubung dan dimigrasi.")
	return db
}

func SetupRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	return rdb
}