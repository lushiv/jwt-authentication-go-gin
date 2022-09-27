package main

import (
	"fmt"
	"net/http"
	"os"
	"jwt-auth-gin-app/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.loadEnvVariables()
	initializers.connectDB()
	initializers.migrateDB()
}

func main() {
	router := gin.Default()
	fmt.Println("Starting server on Port: 8080")
	http.ListenAndServe(os.Getenv("PORT"), router)
}
