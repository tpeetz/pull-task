package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show configuration",
	Long:  "Show configuration details",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("show configuration details")
		}
		fmt.Println("Configuration:")
		serverList := serverConfiguration.List()
		for _, server := range serverList {
			fmt.Printf("Server configuration: %v\n", server)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
