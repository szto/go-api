package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
)

var html = template.Must(template.New("http").Parse(`
<html>
	<head>
		<title>http test</title>
	</head>
	<body>
		<h1 style="color:red;"> Welcome.</h1>
	</body>
</html>
`))

func setupRouter() *gin.Engine {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("DON'T USE THIS")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// return user name
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		// c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, gin.H{"user": name})
	})

	// return html
	router.SetHTMLTemplate(html)
	router.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "http", gin.H{
			"status": "success",
		})
	})

	return router

}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
