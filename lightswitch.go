package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var Version string

// lightswitch represents the base command when called without any subcommands.
var lightswitch = &cobra.Command{
	Use:     "lightswitch",
	Short:   "Manage Hue devices",
	Long:    `Manage Hue devices from your terminal`,
	Run:     initGui().run,
	Version: Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the lightswitch.
func execute() {
	cobra.CheckErr(lightswitch.Execute())
}

func runLightswitch() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	lightswitch.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lightswitch)")

	// run the app
	execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".lightswitch")
	}

	viper.SetEnvPrefix("LS")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
