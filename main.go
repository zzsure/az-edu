package main

import (
	"github.com/urfave/cli"
	"az-edu/cmd/server"
	"az-edu/cmd/tool"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "az-edu"
	app.Commands = []cli.Command{
		server.Server,
		tool.InitDB,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
