package command

import (
	"fmt"
	"bytes"
	"bufio"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/appfac/cli/terminal"
)
//Unix specific
const (
	sttyArg0   = "/bin/stty"
	exec_cwdir = ""
)

// Tells the terminal to turn echo off.
var sttyArgvEOff []string = []string{"stty", "-echo"}

// Tells the terminal to turn echo on.
var sttyArgvEOn []string = []string{"stty", "echo"}

var ws syscall.WaitStatus = 0


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
	var username string
	fmt.Println("username: ")
	fmt.Scanf("%s", &username)
	//fmt.Println("password:")
	//fmt.Scanf("%s", &password)
	reqs.UserName=username
	reqs.Password=login.AskForPassword("Password")
	//reqs.Password=password
	reqs.Cookie=""
	return
}

/**Ask Password functionality for unix*/

func (login Login) AskForPassword(prompt string) (passwd string) {
	sig := make(chan os.Signal, 10)

	// Display the prompt.
	fmt.Println("")
	fmt.Printf(prompt+": ")

	// File descriptors for stdin, stdout, and stderr.
	fd := []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}

	// Setup notifications of termination signals to channel sig, create a process to
	// watch for these signals so we can turn back on echo if need be.
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT,
		syscall.SIGTERM)
	defer signal.Stop(sig)

	go catchSignal(fd, sig)

	pid, err := echoOff(fd)
	defer echoOn(fd)
	if err != nil {
		return
	}

	passwd = readPassword(pid)

	// Carriage return after the user input.
	fmt.Println("")

	return
}

func readPassword(pid int) string {
	rd := bufio.NewReader(os.Stdin)
	syscall.Wait4(pid, &ws, 0, nil)

	line, err := rd.ReadString('\n')
	if err == nil {
		return strings.TrimSpace(line)
	}
	return ""
}


// catchSignal tries to catch SIGKILL, SIGQUIT and SIGINT so that we can turn terminal
// echo back on before the program ends.  Otherwise the user is left with echo off on
// their terminal.
func catchSignal(fd []uintptr, sig chan os.Signal) {
	select {
	case <-sig:
		echoOn(fd)
		os.Exit(2)
	}
}


func echoOff(fd []uintptr) (int, error) {
	pid, err := syscall.ForkExec(sttyArg0, sttyArgvEOff, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})

	if err != nil {
		//Removed error and replaced with nil
		return 0, nil
	}
	return pid, nil
}

// echoOn turns back on the terminal echo.
func echoOn(fd []uintptr) {
	// Turn on the terminal echo.
	pid, e := syscall.ForkExec(sttyArg0, sttyArgvEOn, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})

	if e == nil {
		syscall.Wait4(pid, &ws, 0, nil)
	}
}

