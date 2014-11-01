package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

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
			Name:     "bluemix-push",
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

	plugin.CliCommand("login", "-a", host, "-u", user, "-p", string(pass), "-o", "cfplayground", "-s", "dev")
}

func bluemix_push() {
	var gitUrl string

	fmt.Print("Source code repo url: (https://github.com/simonleung8/dora.git) ")
	fmt.Scanf("%s", &gitUrl)

	gitUrl = "https://github.com/simonleung8/dora.git"

	tmpDir, err := ioutil.TempDir("./", "app-gitpush")
	if err != nil {
		fmt.Println("Error creating temp dir: ", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	if err = cloneRepo(gitUrl, tmpDir); err != nil {
		fmt.Println("Error cloning git repo:", err)
		os.Exit(1)
	}

	if err = pushApp(tmpDir); err != nil {
		fmt.Println("Error pushing app to bluemix", err)
		os.Exit(1)
	}
}

func main() {
	plugin.Start(new(IBM_Bluemix))
}

func cloneRepo(url, outDir string) error {
	fmt.Println("Cloning git repo ", url, "...")
	cmd := exec.Command("git", "clone", url, outDir)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func pushApp(outDir string) error {
	fmt.Println("Pushing app to bluemix...")
	_, err := plugin.CliCommand("push", "-p", outDir, "cli_plugin_dora")
	return err
}
