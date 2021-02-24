package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
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
	r.Gin.Use(favicon.New("./assets/icon/favicon.ico"))
	r.Gin.LoadHTMLGlob(r.AbsolutePath + "/app/interfaces/presenters/index.html")
	r.Gin.Static("/assets", r.AbsolutePath + "./assets")
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
	err := r.Gin.Run(":8080")
	if err != nil {
		log.Fatal()
	}
}