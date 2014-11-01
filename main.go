package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/howeyc/gopass"
)

type IBM_Bluemix struct{}

func (c *IBM_Bluemix) Run(args []string, reply *bool) error {
	if args[0] == "bluemix-login" {
		bluemix_login()
	} else if args[0] == "bluemix-push" {
		bluemix_push()
	}
	return nil
}

func (c *IBM_Bluemix) GetCommands() []plugin.Command {
	return []plugin.Command{
		{
			Name:     "bluemix-login",
			HelpText: "help text for test_1_cmd1",
		},
		{
			Name:     "bluemix_push",
			HelpText: "help text for test_1_cmd2",
		},
	}
}

func bluemix_login() {
	var host, user string
	fmt.Print("login api: (https://api.ng.bluemix.net)")
	fmt.Scanf("%s", &host)
	if host == "" {
		host = "https://api.ng.bluemix.net"
	}

	fmt.Print("login for bluemix: ")
	fmt.Scanf("%s", &user)

	fmt.Print("password for bluemix: ")
	pass := gopass.GetPasswd()

	plugin.CliCommand("login", "-a", host, "-u", user, "-p", string(pass), "-o", "cfplayground", "-s", "test")
}

func bluemix_push() {
	var gitUrl string

	fmt.Print("Source code repo url: (https://github.com/simonleung8/cf_sample_app.git) ")
	fmt.Scanf("%s", &gitUrl)

	if gitUrl == "" {
		gitUrl = "https://github.com/simonleung8/cf_sample_app.git"
	}
}

func main() {
	plugin.Start(new(Sample))
}
