package server

import (
	"az-edu/conf"
	"az-edu/controller/v1"
	"az-edu/library/db"
	"az-edu/library/log"
	"az-edu/middleware"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var Server = cli.Command{
	Name:  "server",
	Usage: "az-edu http server",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
			Usage: "toml配置文件",
		},
		cli.StringFlag{
			Name:  "args",
			Value: "",
			Usage: "multi config cmd line args",
		},
	},
	Action: run,
}

func run(c *cli.Context) {
	conf.Init(c.String("conf"), c.String("args"))
	log.Init()
	db.Init()

	_ = GinEngine().Run(conf.Config.Server.Listen)
}

func GinEngine() *gin.Engine {
	var r *gin.Engine
	if conf.Config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(middleware.Recovery)
	} else {
		r = gin.Default()
	}
	r.Use(middleware.Access)
	r.Use(middleware.Auth)
	r.GET("/health")
	V1(r)

	return r
}

func V1(r *gin.Engine) {
	g := r.Group("/api/v1")
	{
		g.POST("/question/add", v1.QuestionAdd)
		g.POST("/label/add", v1.LabelAdd)
	}
}
