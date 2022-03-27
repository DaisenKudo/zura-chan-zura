package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	yaml "gopkg.in/yaml.v2"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Routing struct {
	Gin          *gin.Engine
	FaceList     FaceList
	AbsolutePath string
}

type FaceList struct {
	Faces []string `yaml:"faces"`
}

func NewRouting() *Routing {
	c, _ := NewConfig()
	//facelistをyamlとして読み込んでRoutingへ格納
	f, _ := os.ReadFile("./dist/assets/facelist.yaml") //TODO:pathを合わせる
	var f2 FaceList
	err := yaml.UnmarshalStrict(f, &f2)
	if err != nil {
		panic(err)
	}

	r := &Routing{
		Gin:          gin.Default(),
		FaceList:     f2,
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
	const ZURA = "ずらちゃんずら"
	const DEPLOY = "http://zura-chan-zura.com"

	r.Gin.GET("/", func(c *gin.Context) {
		face := r.getFace()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": ZURA + "💓",
			"text":  ZURA,
			"face":  face,
			"href": "https://twitter.com/share" +
				"?url=" + "\n\n" + DEPLOY +
				"&text=" + ZURA + face,
		})
	})
}

func (r *Routing) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	return r.Gin.Run(":" + port)
}

func (r *Routing) getFace() string {
	rand.Seed(time.Now().UnixNano())
	return r.FaceList.Faces[rand.Intn(len(r.FaceList.Faces))]
}
