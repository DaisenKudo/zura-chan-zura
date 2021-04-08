package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
	"os"
)

type Routing struct {
	Gin *gin.Engine
	AbsolutePath string
}

func NewRouting() *Routing {
	c, _ := NewConfig()
	r := &Routing{
		Gin: gin.Default(),
		AbsolutePath: c.AbsolutePath,
	}
	r.loadTemplates()
	r.setRouting()
	return r
}

func (r *Routing) loadTemplates() {
	r.Gin.Use(favicon.New("./static/icon/favicon.ico"))
	r.Gin.Static("/static", r.AbsolutePath + "/static")
	r.Gin.LoadHTMLGlob(r.AbsolutePath + "/app/interfaces/presenters/*")
}

func (r *Routing) setRouting() {
	r.Gin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"title" : "ずらちゃんずら💓",
			"text" : "ずらちゃんずら💓",
		})
	})
}

func (r *Routing) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return r.Gin.Run(":" + port)
}
