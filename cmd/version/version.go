package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	VERSION string
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "prints columbus-cli version",
	Run: func(cmd2 *cobra.Command, args []string) {
		fmt.Println("columbus-cli version:" + " v" + VERSION)
	},
}

func SetVersion(version string) {
	VERSION = version
}
