package stop

import (
	"fmt"
	"github.com/n3xtdata/columbus-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var _ = fmt.Printf

var (
	getLong = `
		Stops the columbus jar
		The process id of the current will be looked up in $COLUMBUS_HOME/tmp/pid.txt
		If this file does not exist or is empty there is currently no process running which
		was started by columbus-cli
		`

	getExample = `
		# To stop columbus type
		kubectl stop
		`
)
var Cmd = &cobra.Command{
	Use:     "stop",
	Short:   "stops columbus",
	Long:    getLong,
	Example: getExample,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Stopping columbus")

		var s, i = utils.GetPid()

		if i == 0 {
			fmt.Println("Columubs is not running")
		} else {
			var p, err3 = os.FindProcess(i)

			if err3 != nil {
				panic(err3)
			}

			p.Kill()

			fmt.Println("Columbus was stopped, PID = " + s)

			utils.WritePid("")

		}

	},
}
