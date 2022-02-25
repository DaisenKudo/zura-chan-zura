package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	yaml "gopkg.in/yaml.v2"
	"net/http"
	"math/rand"
	"time"
	"os"
)

type Routing struct {
	Gin *gin.Engine
	FaceList FaceList
	AbsolutePath string
}

type FaceList struct {
	Faces []string `yaml:"faces"`
}

func NewRouting() *Routing {
	c, _ := NewConfig()
	//facelistをyamlとして読み込んでRoutingへ格納
	f, _ := os.readFile("facelist.yaml") //TODO:pathを合わせる
	var f2 FaceList
	err := yaml.UnmarshalStrict(f, &f2)
	if err != nil {
		panic(err)
	}
	
	r := &Routing{
		Gin: gin.Default(),
		Faces: f2
		AbsolutePath: c.AbsolutePath,
	}
	r.loadTemplates()
	r.setRouting()
	return r
}

func (r *Routing) loadTemplates() {
	r.Gin.Use(favicon.New("./dist/assets/favicon.ico"))
	r.Gin.Static("/assets", r.AbsolutePath + "/dist/assets")
	r.Gin.LoadHTMLGlob(r.AbsolutePath + "/app/interfaces/presenters/*")
}

func (r *Routing) setRouting() {
	const ZURA = "ずらちゃんずら"
	const DEPLOY = "https://zura-chan-zura.herokuapp.com"

	r.Gin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"title" : ZURA + "💓",
			"text" : ZURA,
			"face" : getFace() + "💓",
			"href" : "\nhttps://twitter.com/share" +
				"?url=" + DEPLOY +
				"&text=" + ZURA + "💓",
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

func (r *Routing) getFace() string {
	rand.Seed(time.Now().UnixNano())
	return r.FaceList[rand.Intn(len(r.FaceList))]
}
