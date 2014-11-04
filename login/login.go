package login

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/howeyc/gopass"
)

func Login() {
	var choice int
	fmt.Printf("\nSelect a target:\n1. Local Bosh-lite\n2. Bluemix\n\nTarget> ")
	fmt.Scanf("%d", &choice)
	if choice == 1 {
		boshlite_login()
	} else if choice == 2 {
		bluemixLogin()
	}
}

func bluemixLogin() {
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

	plugin.CliCommand("login", "-a", host, "-u", user, "-p", string(pass), "-o", "cf_playground")
}

func boshlite_login() {
	fmt.Print("loging into bosh-lite...")
	plugin.CliCommand("login", "-a", "https://api.10.244.0.34.xip.io", "-u", "admin", "-p", "admin", "--skip-ssl-validation")
}
