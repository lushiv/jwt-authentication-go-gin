package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	fmt.Println("Starting server on Port: 8080")

	http.ListenAndServe(":8080", router)

}
