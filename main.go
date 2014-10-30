package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type SamplePlugin struct {
}

func (c *SamplePlugin) Run(args []string, reply *bool) error {
	if args[0] == "test_1_cmd1" {
		theFirstCmd()
	} else if args[0] == "test_1_cmd2" {
		theSecondCmd()
	}
	return nil
}

func (c *SamplePlugin) GetCommands() []plugin.Command {
	return []plugin.Command{
		{
			Name:     "test_1_cmd1",
			HelpText: "help text for test_1_cmd1",
		},
		{
			Name:     "test_1_cmd2",
			HelpText: "help text for test_1_cmd2",
		},
	}
}

func theFirstCmd() {
	fmt.Println("You called cmd1 in test_1")
}

func theSecondCmd() {
	fmt.Println("You called cmd2 in test_1")
}

func main() {
	plugin.Start(new(SamplePlugin))
}
