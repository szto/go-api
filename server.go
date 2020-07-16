package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		// c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, gin.H{"user": name})
	})
	return router

}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
