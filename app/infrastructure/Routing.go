package infrastructure

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	yaml "gopkg.in/yaml.v2"
)

type Routing struct {
	Gin          *gin.Engine
	School       School
	AbsolutePath string
}

type School struct {
	Uranohoshi []Character `yaml:"uranohoshi"`
	Hashnosora []Character `yaml:"hasunosora"`
}

type Character struct {
	Faces []string `yaml:"faces"`
	Lines []string `yaml:"lines"`
}

func NewRouting() *Routing {
	c, _ := NewConfig()
	//CharacterListをyamlとして読み込んでRoutingへ格納
	f, _ := os.ReadFile("./dist/assets/CharacterList.yaml") //TODO:pathを合わせる
	var f2 School
	err := yaml.UnmarshalStrict(f, &f2)
	if err != nil {
		panic(err)
	}

	r := &Routing{
		Gin:          gin.Default(),
		School:       f2,
		AbsolutePath: c.AbsolutePath,
	}
	r.loadTemplates()
	r.setRouting()
	return r
}

func (r *Routing) loadTemplates() {
	r.Gin.Use(favicon.New("./dist/assets/favicon.ico"))
	r.Gin.Static("/assets", r.AbsolutePath+"/dist/assets")
	r.Gin.LoadHTMLGlob(r.AbsolutePath + "/app/interfaces/presenters/*")
}

func (r *Routing) setRouting() {
	const DEPLOY = "https://zura-chan-zura.com"

	r.Gin.GET("/", func(c *gin.Context) {
		char := r.School.Uranohoshi[rand.Intn(len(r.School.Uranohoshi))]
		line := r.getLine(char)
		face := r.getFace(char)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": line + "💓",
			"text":  line,
			"face":  face,
			"href": "https://twitter.com/intent/tweet" +
				"?url=" + "\n\n" + DEPLOY +
				"&text=" + line + face,
		})
	})

	r.Gin.GET("/hasu", func(c *gin.Context) {
		char := r.School.Hashnosora[rand.Intn(len(r.School.Hashnosora))]
		line := r.getLine(char)
		face := r.getFace(char)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": line + "💓",
			"text":  line,
			"face":  face,
			"href": "https://twitter.com/intent/tweet" +
				"?url=" + "\n\n" + DEPLOY +
				"&text=" + line + face,
		})
	})

	r.Gin.HEAD("/", func(c *gin.Context) {
		char := r.School.Uranohoshi[rand.Intn(len(r.School.Uranohoshi))]
		face := r.getFace(char)
		line := r.getLine(char)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": line,
			"text":  line,
			"face":  face,
			"href": "https://twitter.com/share" +
				"?url=" + "\n\n" + DEPLOY +
				"&text=" + line + face,
		})
	})

	r.Gin.POST("/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})

	r.Gin.PUT("/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})

	r.Gin.DELETE("/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})

	r.Gin.Handle(http.MethodConnect, "/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})

	r.Gin.OPTIONS("/*any", func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"text": "ずらちゃん is member of Aqours.",
		})
	})

	r.Gin.PATCH("/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})

	r.Gin.Handle(http.MethodTrace, "/*any", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusForbidden)
	})
}

func (r *Routing) Run() error {
	port := "8703"
	return r.Gin.Run(":" + port)
}

func (r *Routing) getFace(char Character) string {
	return char.Faces[rand.Intn(len(char.Faces))]
}

func (r *Routing) getLine(char Character) string {
	return char.Lines[rand.Intn(len(char.Lines))]
}
