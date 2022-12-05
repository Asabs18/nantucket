/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"math"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/stat"
)

// CONSTANTS
const SICK_DAYS = 14
const POPULATION_SIZE = 10000
const V_PROB = 0.0
const T_PROB = 1.5
const D_PROB = 1.5
const N_DAYS = 100
const INITIAL_INFECTION_PERCENT = 1
const NUM_EXPOSERS_MULTIPLIER = 20
const N_TRIALS = 100
const OUTPUT_FILE = "output.txt"
const RED = "#FF0000"
const GREEN = "#065535"
const BLUE = "#0000FF"
const ORANGE = "#FFA500"
const YELLOW = "#FFFF00"
const WHITE = "#FFFFFF"

// Driver function which updates the simulation for a certain number of days and prints a summary
func Simulate(cmdLineVars CmdLineVars) Environment {
	// Create Environment and local environment variables
	var environment Environment
	var currPopStatus []Person
	var nextPopStatus []Person

	// Init environment variables
	environment.New(cmdLineVars)
	environment.currDay = 0

	// Open output file
	pFile := environment.openFile(cmdLineVars.outputFile)
	defer environment.closeFile(pFile)

	// Update the current and next status arrays for each day and print the changes
	for i := 0; i < cmdLineVars.nDays; i++ {

		// Update environment variables
		environment.currDay = i + 1
		environment.updateDay()
		copy(environment.currPopulationStatus, currPopStatus)
		copy(environment.nextPopulationStatus, nextPopStatus)

		// Write current days data to output file
		environment.writeDayToFile(pFile)

		// Check if the virus has already run its course
		if environment.nSusceptible == 0 || environment.nInfected == 0 {
			if cmdLineVars.numTrials == 1 {
				color.HEX(WHITE).Println("\n\n      Simulation ended after", environment.currDay, "days \n The virus has run its course in Nantucket.")
			}
			break
		}
	}

	return environment
}

// Change to take a single struct which is created outside the func depending on version
func printSummary(environments []Environment) {
	// Extract certain fields to perform stat operations on
	var deaths []float64
	var recovered []float64
	var susceptible []float64
	var infected []float64
	for i := 0; i < cmdLineVars.numTrials; i++ {
		deaths = append(deaths, float64(environments[i].nDead))
		recovered = append(recovered, float64(environments[i].nRecovered))
		susceptible = append(susceptible, float64(environments[i].nSusceptible))
		infected = append(infected, float64(environments[i].nInfected))
	}

	color.HEX(WHITE).Println("\n\n\n=============== MULTI-TRIAL SUMMARY ===============\n")
	color.HEX(WHITE).Println("Number of Trials:             ", N_TRIALS, "\n")
	color.HEX(YELLOW).Println("SUSCEPTIBLE:")
	color.HEX(WHITE).Println("     Average Susceptible Across Trials: ", stat.Mean(susceptible, nil))
	color.HEX(WHITE).Println("     STD of Susceptible Across Trials:  ", math.Sqrt(stat.Variance(susceptible, nil)))
	color.HEX(ORANGE).Println("INFECTED:")
	color.HEX(WHITE).Println("     Average Infections Across Trials:  ", stat.Mean(infected, nil))
	color.HEX(WHITE).Println("     STD of Infections Across Trials:   ", math.Sqrt(stat.Variance(infected, nil)))
	color.HEX(GREEN).Println("RECOVERIES:")
	color.HEX(WHITE).Println("     Average Recoveries Across Trials:  ", stat.Mean(recovered, nil))
	color.HEX(WHITE).Println("     STD of Recoveries Across Trials:   ", math.Sqrt(stat.Variance(recovered, nil)))
	color.HEX(RED).Println("DEATHS:")
	color.HEX(WHITE).Println("     Average Deaths Across Trials:      ", stat.Mean(deaths, nil))
	color.HEX(WHITE).Println("     STD of Deaths Across Trials:       ", math.Sqrt(stat.Variance(deaths, nil)))

}

var cmdLineVars CmdLineVars

func init() {

	rootCmd.PersistentFlags().Float64VarP(&cmdLineVars.vProb, "vProb", "v", V_PROB, "Vaccination probability for each person when generated")
	rootCmd.PersistentFlags().Float64VarP(&cmdLineVars.tProb, "tProb", "t", T_PROB, "Transmission probability each day when sick")
	rootCmd.PersistentFlags().Float64VarP(&cmdLineVars.dProb, "dProb", "x", D_PROB, "Death probability each day when sick")

	rootCmd.PersistentFlags().StringVarP(&cmdLineVars.outputFile, "outputFile", "o", OUTPUT_FILE, "File to output result to")

	simulateCmd.Flags().IntVarP(&cmdLineVars.populationSize, "popSize", "p", POPULATION_SIZE, "Size of population")
	simulateCmd.Flags().IntVarP(&cmdLineVars.nDays, "nDays", "d", N_DAYS, "Number of Days")
	simulateCmd.Flags().IntVarP(&cmdLineVars.numTrials, "numTrials", "n", N_TRIALS, "Number of Trials")

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(simulateCmd)
}

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.HEX(GREEN).Println("simulate called")

		var environments []Environment

		for i := 0; i < cmdLineVars.numTrials; i++ {
			environments = append(environments, Simulate(cmdLineVars))
			color.HEX(GREEN).Println("RUNNING SIMULATION #", i+1, "| Deaths: ", environments[i].nDead, "| Recoveries: ", environments[i].nRecovered, "| Infected: ", environments[i].nInfected, "| Susceptible: ", environments[i].nSusceptible, "|")
		}

		printSummary(environments)

	},
}
