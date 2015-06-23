package main

import (
	"os"
	"github.com/appfac/cli/command"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "appfac"
	app.Usage = "CLI Tool for WSO2 Appfactory"
	app.Action = func(c *cli.Context) {
		println("first appfac CLI command!")
	}

	//app.Run(os.Args)

	//command `appfac` without argument
	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "h" {
		println("Showing help commands")
		app.Run(os.Args)
	}else if os.Args[1] == "login"{
		c := command.CommandConfigs{"https://apps.cloud.wso2.com/appmgt/site/blocks/user/login/ajax/login.jag", "action=login&userName=dilhasha.wso2.com@dilhashan&password=MASHAALLAH_dilu1" , ""}
		c.Run()
	}

}








