package command

import (
	"fmt"
	"bytes"
)

type Login struct {
	//should get repo location for login details
}

func NewLogin() (cmd Login) {
	return
}

func (login Login)Metadata() CommandMetadata{
	return CommandMetadata{
		Name:"Login",
		Description : "Login to app factory",
		ShortName : "l",
		Usage:"login",
	}
}

func (login Login)Configs(reqs CommandRequirements)(configs CommandConfigs){
	var buffer bytes.Buffer
	buffer.WriteString("action=login")
	if(reqs.UserName!=""){
		buffer.WriteString("&userName=")
		buffer.WriteString(reqs.UserName)
	}
	if(reqs.Password!=""){
		buffer.WriteString("&password=")
		buffer.WriteString(reqs.Password)

	}

	s := buffer.String()
	return CommandConfigs{
		Url:"https://apps.cloud.wso2.com/appmgt/site/blocks/user/login/ajax/login.jag",
		Query:s,
		Cookie:reqs.Cookie,
	}
}

func (login Login) Requirements()(reqs CommandRequirements){
	var username,password string
	fmt.Println("username:")
	fmt.Scanf("%s", &username)
	fmt.Println("password:")
	fmt.Scanf("%s", &password)
	reqs.UserName=username
	reqs.Password=password
	reqs.Cookie=""
	return
}
