package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tpeetz/pull-task/pkg/server"
)

const (
	// TaskwarriorConfiguration defines the parameter to disable the config confirmation
	TaskwarriorConfiguration = "rc.confirmation=off"
)

var (
	cfgFile string
	// Verbose is set, when verbose on the commandline is given.
	Verbose bool
)

var (
	rootCmd = &cobra.Command{
		Use:   "pull-task",
		Short: "pull-task loads issues from Github, Gitlab and Redmine",
		Long:  `pull-task loads issues from Github, Gitlab and Redmine and synchronizes with TaskWarrior`,
	}
	serverConfiguration = server.Configuration{}
)

// Execute sets the version of ibtp-gitlab abd executes the command.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/pull-task.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("pull-task")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.config")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if Verbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
		settings := viper.AllSettings()
		for _, v := range settings {
			gitServer, err := server.NewGitServer(v.(map[string]interface{}))
			if err == nil && gitServer != nil {
				serverConfiguration.Add(gitServer)
			}
			//fmt.Printf("\nServer List: %v\n", serverConfiguration)
		}
	}
}
