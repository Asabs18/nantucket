/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"github.com/dariubs/percent"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	//"gonum.org/v1/gonum/stat"
)

// CONSTANTS
const SICK_DAYS = 14
const POPULATION_SIZE = 10000
const V_PROB = 0.0
const T_PROB = 1.5
const D_PROB = 1.5
const N_DAYS = 100
const INITIAL_INFECTION_PERCENT = 1
const ZERO = 0
const NUM_EXPOSERS_MULTIPLIER = 20
const RED = "#FF0000"
const GREEN = "#065535"
const BLUE = "#0000FF"
const ORANGE = "#FFA500"
const YELLOW = "#FFFF00"
const WHITE = "#FFFFFF"
const N_TRIALS = 50

//SECTION ===== DRIVER FUNCTION =====

// Final Summary of simulation
func printFinalSummary(environment Environment) {
	color.HEX(WHITE).Println("\n\n\n================== FINAL SUMMARY ==================")
	color.HEX(WHITE).Println("     Population:               ", environment.cmdLineVars.populationSize)
	color.HEX(WHITE).Println("     Vaccination Probability:  ", environment.cmdLineVars.vProb)
	color.HEX(WHITE).Println("     Transmission Probability: ", environment.cmdLineVars.tProb)
	color.HEX(WHITE).Println("     Death Probability:        ", environment.cmdLineVars.dProb)
	color.HEX(WHITE).Println("     Initial Infections:       ", environment.nInfected)
	color.HEX(WHITE).Println("     Simulation Period:        ", environment.cmdLineVars.nDays, "Days")
	color.HEX(WHITE).Println("     Number Recovered:         ", environment.nRecovered)
	color.HEX(WHITE).Println("     Number Dead:              ", environment.nDead)
	color.HEX(WHITE).Println("     Case Fatality Rate:       ", percent.PercentOf(environment.nDead, environment.cmdLineVars.populationSize), "Percent")
}

// Driver function which updates the simulation for a certain number of days and prints a summary
func Simulate(cmdLineVars CmdLineVars, fullSummary bool) Environment {
	// Create Environment and local environment variables
	var environment Environment
	var currPopStatus []Person
	var nextPopStatus []Person

	// Init environment variables
	environment.New()
	environment.cmdLineVars = cmdLineVars
	environment.currDay = ZERO

	// Open output file
	pFile := environment.openFile(cmdLineVars.outputFile)
	defer environment.closeFile(pFile)

	// Update the current and next status arrays for each day and print the changes
	for i := ZERO; i < cmdLineVars.nDays; i++ {

		// Update environment variables
		environment.currDay = i + 1
		environment.updateDay()
		copy(environment.currPopulationStatus, currPopStatus)
		copy(environment.nextPopulationStatus, nextPopStatus)

		// Write current days data to output file
		environment.writeDayToFile(pFile)

		// Check if the virus has already run its course
		if environment.nSusceptible == ZERO || environment.nInfected == ZERO {
			if cmdLineVars.numTrials == 1 {
				color.HEX(WHITE).Println("\n\n      Simulation ended after", environment.currDay, "days \n The virus has run its course in Nantucket.")
			}
			break
		}
	}

	// Display a final summary of the simulation

	if fullSummary {
		printFinalSummary(environment)
	}

	return environment
}

func printMultiTrialSummary(environments []Environment) {
	// Extract certain fields to perform stat operations on
	var deaths []float64
	var recovered []float64
	var susceptible []float64
	var infected []float64
	for i := 0; i < N_TRIALS; i++ {
		deaths = append(deaths, float64(environments[i].nDead))
		recovered = append(recovered, float64(environments[i].nRecovered))
		susceptible = append(susceptible, float64(environments[i].nSusceptible))
		infected = append(infected, float64(environments[i].nInfected))
	}

	color.HEX(WHITE).Println("\n\n\n=============== MULTI-TRIAL SUMMARY ===============\n")
	color.HEX(WHITE).Println("Number of Trials:             ", N_TRIALS, "\n")
	//TODO: Fix the rest of the multitrial summary
	// color.HEX(YELLOW).Println("SUSCEPTIBLE:")
	// color.HEX(WHITE).Println("     Average Susceptible Across Trials: ", stat.Mean(susceptible, nil))
	// color.HEX(WHITE).Println("     STD of Susceptible Across Trials:  ", math.Sqrt(stat.Variance(susceptible, nil)))
	// color.HEX(ORANGE).Println("INFECTED:")
	// color.HEX(WHITE).Println("     Average Infections Across Trials:  ", stat.Mean(infected, nil))
	// color.HEX(WHITE).Println("     STD of Infections Across Trials:   ", math.Sqrt(stat.Variance(infected, nil)))
	// color.HEX(GREEN).Println("RECOVERIES:")
	// color.HEX(WHITE).Println("     Average Recoveries Across Trials:  ", stat.Mean(recovered, nil))
	// color.HEX(WHITE).Println("     STD of Recoveries Across Trials:   ", math.Sqrt(stat.Variance(recovered, nil)))
	// color.HEX(RED).Println("DEATHS:")
	// color.HEX(WHITE).Println("     Average Deaths Across Trials:      ", stat.Mean(deaths, nil))
	// color.HEX(WHITE).Println("     STD of Deaths Across Trials:       ", math.Sqrt(stat.Variance(deaths, nil)))

}

var cmdLineVars CmdLineVars

func init() {

	// var vProb float32
	// var tProb float32
	// var dProb float32
	// var nDays int
	// var output_file string
	// var numTrials int

	rootCmd.Flags().Int("popSize", cmdLineVars.populationSize, "popSize")
	rootCmd.PersistentFlags().Float64P("vProb", "v", cmdLineVars.vProb, "vProb")
	rootCmd.PersistentFlags().Float64P("tProb", "t", cmdLineVars.tProb, "tProb")
	rootCmd.PersistentFlags().Float64P("dProb", "d", cmdLineVars.dProb, "dProb")
	rootCmd.Flags().Int("nDays", cmdLineVars.nDays, "nDays")
	simulateCmd.Flags().Int("multi", cmdLineVars.numTrials, "Run multiple trials or not")
	rootCmd.PersistentFlags().String("output-file", cmdLineVars.outputFile, "output-file")
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(simulateCmd)

	//TODO: Use "var := cmd.Flag("varName")" in what context?
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

		if cmdLineVars.numTrials > 1 {
			var environments []Environment

			for i := 0; i < N_TRIALS; i++ {
				environments = append(environments, Simulate(cmdLineVars, false))
				color.HEX(GREEN).Println("RUNNING SIMULATION #", i+1, "| Deaths: ", environments[i].nDead, "| Recoveries: ", environments[i].nRecovered, "| Infected: ", environments[i].nInfected, "| Susceptible: ", environments[i].nSusceptible, "|")
			}

			printMultiTrialSummary(environments)
		} else {
			Simulate(cmdLineVars, true)
		}
	},
}

// COMMAND LINE INSTRUCTIONS:
//     RUN DEFAULT -       "go run main.go simulate"
//     POPULATION SIZE -   "[RUN DEFAULT] pop=[INPUT: int]"
//     VACCINE PROB -      "[RUN DEFAULT] vprob=[INPUT: float64]"
//     TRANSMISSION PROB - "[RUN DEFAULT] tprob=[INPUT: float64]"
//     DEATH PROB -        "[RUN DEFAULT] dprob=[INPUT: float64]"
//     NUM DAYS -          "[RUN DEFAULT] days=[INPUT: int]"
//     OUTPUT FILE -       "[RUN DEFAULT] file=[INPUT: string]"
//     MULTI-TRIALS -      "[RUN DEFAULT] multi=[INPUT: bool]"
