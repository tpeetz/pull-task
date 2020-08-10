package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load issues",
	Long:  "Load issues from configured servers",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("load issues from configured servers")
		}
		serverList := serverConfiguration.List()
		for _, server := range serverList {
			server.LoadIssues()
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
