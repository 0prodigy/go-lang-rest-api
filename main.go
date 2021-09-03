package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
    router.GET("/hello", helloWorld)

    router.Run("localhost:8080")
}

func helloWorld(c *gin.Context) {
		fmt.Println("Hello World")
	
}