package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
	"os"
)

type Routing struct {
	Gin *gin.Engine
	AbsolutePath string
}

func NewRouting() *Routing {
	c := NewConfig()
	r := &Routing{
		Gin: gin.Default(),
		AbsolutePath: c.AbsolutePath,
	}
	r.loadTemplates()
	r.setRouting()
	return r
}

func (r *Routing) loadTemplates() {
	println("-------------------------------")
	println(r.AbsolutePath)
	println("-------------------------------")
	r.Gin.Use(favicon.New("./assets/icon/favicon.ico"))
	r.Gin.Static("/assets", r.AbsolutePath + "/assets")
	r.Gin.LoadHTMLGlob(r.AbsolutePath + "/app/interfaces/presenters/index.html")
}

func (r *Routing) setRouting() {
	r.Gin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"title" : "ずらちゃんずら💓",
			"text" : "ずらちゃんずら💓",
		})
	})
}

func (r *Routing) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Gin.Run(":" + port)
	if err != nil {
		log.Fatal()
	}
}