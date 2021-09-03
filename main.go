package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
    router.GET("/hello/:id", helloWorld)
	router.POST("/upload",handleUpload)
    router.Run("localhost:8080")
}

func handleUpload(c *gin.Context){
	file,_ := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func helloWorld(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
		fmt.Println("Hello World")
	
}