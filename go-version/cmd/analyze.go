/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analyze called")

		var environments []Environment

		cmdLineVars := parseCmdLine(args)

		if cmdLineVars.multiTrials {
			for i := 0; i < TRIALS; i++ {
				environments = append(environments, Simulate(cmdLineVars, false))
				color.HEX(GREEN).Println("RUNNING SIMULATION #", i+1, "| Deaths: ", environments[i].nDead, "| Recoveries: ", environments[i].nRecovered, "| Infected: ", environments[i].nInfected, "| Susceptible: ", environments[i].nSusceptible, "|")
			}

			printMultiTrialSummary(environments)
		} else {
			Simulate(cmdLineVars, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// analyzeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// analyzeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
