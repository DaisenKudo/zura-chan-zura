package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Use(favicon.New("./assets/icon/favicon.ico"))
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./assets")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title" : "ずらちゃんずら💓",
			"text" : "ずらちゃんずら💓",
		})
	})
	engine.Run(":8080")
}