package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() { // give json to /
	fmt.Println("Listening on :8080...")
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "ok",
		})
	})
	router.Run()
}
