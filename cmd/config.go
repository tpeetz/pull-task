package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tpeetz/pull-task/pkg/server"
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
		settings := viper.AllSettings()
		fmt.Printf("Settings: %v\n", settings)
		for k, v := range settings {
			fmt.Printf("%v: %v\n", k, v)
			server.NewGitServer(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
