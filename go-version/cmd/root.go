/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// BACK: run text through MD processor?
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nantucket",
	Short: "An attempt to implement the _nantucket_ project as, defined in *nantucket.pdf*, in idiomatic `GO`",
	Long: `
_nantucket_ is a simulation of how different percentages of
vaccinated people affect the spread of a virus.

Measures like mask-wearing and social distancing are considered when 
determining the probability that the virus will be transmitted from one
person to another. That is: How does transmission probability influence the 
spread of the disease?`,
	// HACK: versioning?
	Version: "0.0.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// HACK: where to put config vars?
	// var popSize int
	// var vProb float32
	// var tProb float32
	// var dProb float32
	// var nDays int
	// var output_file string
	// var multiTrials bool

	// rootCmd.Flags().Int("popSize", popSize, "popSize")
	// rootCmd.PersistentFlags().Float32P("vProb", "v", vProb, "vProb")
	// rootCmd.PersistentFlags().Float32P("tProb", "t", tProb, "tProb")
	// rootCmd.PersistentFlags().Float32P("dProb", "d", dProb, "dProb")
	// rootCmd.Flags().Int("nDays", nDays, "nDays")
	// simulateCmd.Flags().BoolP("multi", "m", multiTrials, "Run multiple trials or not")
	// rootCmd.PersistentFlags().String("output-file", output_file, "output-file")
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your project.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nantucket.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// BACK: does nantucket need a config file?
//   - yes - to save scenarios - so long as it can write out CLAs ala `opt`
//
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nantucket" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nantucket")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
