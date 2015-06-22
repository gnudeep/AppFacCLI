package main

import (
	"os"
	"github.com/codegangsta/cli"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
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

		//Testing a curl command

		url := "https://apps.cloud.wso2.com/appmgt/site/blocks/user/login/ajax/login.jag"
		fmt.Println("URL:>", url)

		var jsonStr = []byte(`{action='login',userName='*****',password='*****'}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		app.Run(os.Args)
	}
}

