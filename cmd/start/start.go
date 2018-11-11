package start

import (
	"fmt"
	"github.com/n3xtdata/columbus-cli/cmd/config"
	"github.com/n3xtdata/columbus-cli/utils"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var _ = fmt.Printf

var (
	getLong = `
		Starts the columbus jar`

	getExample = `
		# Starts the columbus jar
		columbus start`
)
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "starts columbus",
	Long:    getLong,
	Example: getExample,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Starting columbus ...")

		if _, err := os.Stat(config.ColumbusJar); os.IsNotExist(err) {

			fmt.Println("the columbus jar does not exist. Downloading from github ...")

			downloadColumbusJar()

		}

		logfile, err := os.Create(config.ColumbusLogFile)

		if err != nil {
			panic(err)
		}

		defer logfile.Close()

		command := exec.Command("java", "-jar", config.ColumbusJar)
		command.Stdout = logfile

		command.Start()

		var pid = strconv.Itoa(command.Process.Pid)

		utils.WritePid(pid)

		fmt.Println("Columbus was started, PID = " + pid)

	},
}

func downloadColumbusJar() {

	// Create the file
	out, err := os.Create(config.ColumbusJar)
	if err != nil {

	}

	defer out.Close()

	// Get the data
	resp, err := http.Get(config.ColumbusJarGithubUrl)
	if err != nil {
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
	}

}
