package log

import (
	"fmt"
	"github.com/n3xtdata/columbus-cli/cmd/config"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var _ = fmt.Printf

var (
	getLong = `
		Displays the Columbus log`

	getExample = `
		# print the latest columbus logs
		columbus log
		`
)



var Cmd = &cobra.Command{
	Use:     "log",
	Short:   "prints the columbus logs",
	Long:    getLong,
	Example: getExample,

	Run: func(cmd *cobra.Command, args []string) {

		var command *exec.Cmd


		if cmd.Flag("follow").Changed {
			command = exec.Command("tail", "-f", config.ColumbusLogFile)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Run()
		} else {
			command = exec.Command("tail", config.ColumbusLogFile)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Start()
		}


	},
}

func init() {
	Cmd.Flags().BoolP("follow", "f", false, "Specify if the logs should be streamed.mak")
}
