package main

import (
	"github.com/simonleung8/cf-cli-plugin/all_apps"
	"github.com/simonleung8/cf-cli-plugin/login"
	"github.com/simonleung8/cf-cli-plugin/push"

	"github.com/cloudfoundry/cli/plugin"
)

type IBM_Bluemix struct{}

func (c IBM_Bluemix) Run(args []string) {
	if args[0] == "bluemix-login" {
		login.Login()
	} else if args[0] == "bluemix-push" {
		push.Push()
	} else if args[0] == "bluemix-all-apps" {
		all_apps.Show()
	}
}

func (c IBM_Bluemix) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "IBM_Bluemix",
		Commands: []plugin.Command{
			{
				Name:     "bluemix-login",
				HelpText: "Options to login to Bosh-lite or Bluemix environment",
			},
			{
				Name:     "bluemix-push",
				HelpText: "Clone source code from a Github repo and push to Cloud Foundry",
			},
			{
				Name:     "bluemix-all-apps",
				HelpText: "List all apps from all spaces",
			},
		},
	}
}

func main() {
	plugin.Start(new(IBM_Bluemix))
}
