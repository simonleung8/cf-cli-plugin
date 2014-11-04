package all_apps

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
)

func Show() {
	orgs := parseOrgs()

	for _, org := range orgs {
		plugin.CliCommand("target", "-s", org)
		plugin.CliCommand("apps")
	}
}

func parseOrgs() []string {
	output, err := plugin.CliCommand("spaces")

	if err != nil {
		fmt.Println("Error getting orgs: ", err)
		os.Exit(1)
	}

	var i int
	for i = 0; i < len(output); i++ {
		if strings.Contains(output[i], "name") {
			i++
			break
		}
	}

	orgs := output[i:]
	for i = 0; i < len(orgs); i++ {
		orgs[i] = Decolorize(strings.TrimSpace(orgs[i]))
	}

	return orgs
}

var decolorizerRegex = regexp.MustCompile(`\x1B\[([0-9]{1,2}(;[0-9]{1,2})?)?[m|K]`)

func Decolorize(message string) string {
	return string(decolorizerRegex.ReplaceAll([]byte(message), []byte("")))
}
