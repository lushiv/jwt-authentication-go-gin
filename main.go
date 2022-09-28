package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

var DB *gorm.DB

func connectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(DB)
	if err != nil {
		panic("Failed to connect database")
	}
}

func migrateDB() {
	DB.AutoMigrate(&User{})
}

func loadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading .env files")
	}
}

func init() {
	loadEnvVariables()
	connectDB()
	migrateDB()
}

func main() {
	router := gin.Default()
	fmt.Println("Starting server on Port: 8080")
	http.ListenAndServe(os.Getenv("PORT"), router)
}
