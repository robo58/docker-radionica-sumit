package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
			"id":      context.Param("id"),
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
