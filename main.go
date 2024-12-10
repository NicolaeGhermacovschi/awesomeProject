package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	connectToDatabase()

	r := gin.Default()

	setupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Nu s-a putut porni serverul:", err)
	}
}
