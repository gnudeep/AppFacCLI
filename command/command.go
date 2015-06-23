package command

import (
	"bytes"
	"fmt"
	"net/http"
	"io/ioutil"
)


type Command interface {
	Metadata() CommandMetadata
	Run()
}

type CommandConfigs struct {
	Url string
	Query string
	Cookie string
}

func (c CommandConfigs) Run() {
	fmt.Println("URL:>", c.Url)
	var jsonStr = []byte(c.Query)
	req, err := http.NewRequest("POST", c.Url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("Cookie", c.Cookie)
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
}

