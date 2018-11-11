package cmd

import (
	"fmt"
	"github.com/n3xtdata/columbus-cli/cmd/get"
	"github.com/n3xtdata/columbus-cli/cmd/log"
	"github.com/n3xtdata/columbus-cli/cmd/start"
	"github.com/n3xtdata/columbus-cli/cmd/status"
	"github.com/n3xtdata/columbus-cli/cmd/stop"
	"github.com/n3xtdata/columbus-cli/cmd/version"
	"os"

	"github.com/spf13/cobra"
)

var (
	ConfigFile string
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "columbus",
	Short: "columbus-cli manages the columbus monitoring tools",
}

// Execute adds all child commands to the root command
func Execute(vers string) {

	version.SetVersion(vers)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(get.Cmd)
	rootCmd.AddCommand(log.Cmd)
	rootCmd.AddCommand(start.Cmd)
	rootCmd.AddCommand(stop.Cmd)
	rootCmd.AddCommand(status.Cmd)
	rootCmd.AddCommand(version.Cmd)
}
