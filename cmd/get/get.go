package get

import (
	"fmt"
	"github.com/spf13/cobra"
	//"log"
	//"reflect"
)

var _ = fmt.Printf

var (
	getLong = `
		Display one or many resources
		Prints a table of the basic information about the specified resource`

	getExample = `
		# List all checks in table output format.
		columbus-cli get checks`
)

var Cmd = &cobra.Command{
	Use:     "get",
	Short:   "display on or many resources",
	Long:    getLong,
	Example: getExample,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {

			switch args[0] {
			case "checks":
				listChecks()
			case "connections":
				fmt.Println("Get Connections")
			}

		} else if len(args) == 2 {
			fmt.Println("bla")
		} else {
			fmt.Println("To many arguments")
		}

	},
}

func init() {
	Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	Cmd.PersistentFlags().String("foo", "", "A help for foo")
}
