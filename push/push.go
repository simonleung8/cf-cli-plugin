package push

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/cloudfoundry/cli/plugin"
)

func Push() {
	var gitUrl, appName string

	fmt.Print("Source code Github Url: (https://github.com/simonleung8/dora.git) ")
	fmt.Scanf("%s", &gitUrl)

	gitUrl = "https://github.com/simonleung8/dora.git"

	fmt.Print("App Name: ")
	fmt.Scanf("%s", &appName)

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

	if err = pushApp(tmpDir, appName); err != nil {
		fmt.Println("Error pushing app to bluemix", err)
		os.Exit(1)
	}
}

func cloneRepo(url, outDir string) error {
	fmt.Println("Cloning git repo ", url, "...")
	cmd := exec.Command("git", "clone", url, outDir)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func pushApp(outDir string, appName string) error {
	fmt.Println("Pushing app to bluemix...")
	_, err := plugin.CliCommand("push", "-p", outDir, "-m", "128M", appName)
	return err
}
