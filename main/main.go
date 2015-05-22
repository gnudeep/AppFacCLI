package main

import (
	"os"
	"github.com/codegangsta/cli"

)

func main() {
	app := cli.NewApp()
	app.Name = "appfac"
	app.Usage = "show help"
	app.Action = func(c *cli.Context) {
		println("first appfac CLI command!")
	}

	//app.Run(os.Args)

	//command `cf` without argument
	if len(os.Args) == 1 || os.Args[1] == "help" {
		println("Showing help commands")
		app.Run(os.Args)
	}
}
