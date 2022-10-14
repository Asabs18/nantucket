/*
Copyright © 2022 Asabs18 & davegi
This file is part of the `nantucket` project.
*/
package cmd

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/dariubs/percent"
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
const ZERO = 0
const NUM_EXPOSERS_MULTIPLIER = 20
const RED = "#FF0000"
const GREEN = "#065535"
const BLUE = "#0000FF"
const ORANGE = "#FFA500"
const YELLOW = "#FFFF00"
const WHITE = "#FFFFFF"
const TRIALS = 50

//SECTION ===== PERSON STRUCT =====

// Create status type and set enum to specific status values
type Status int64

const (
	NewlyInfected Status = 1
	Susceptible   Status = 0
	Recovered     Status = -1
	Vaccinated    Status = -2
	Dead          Status = -3
)

// Define Person ADT which just holds one status
type Person struct {
	status Status
}

// TODO - Automate getStatus
// Takes a status and converts it to the appropriate string
func (person Person) stringifyStatus() string {
	switch person.status {
	case Susceptible:
		return "Susceptible"
	case Recovered:
		return "Recovered"
	case Vaccinated:
		return "Vaccinated"
	case Dead:
		return "Dead"
	default:
		return "Infected"
	}
}

func (person Person) getStatus() Status {
	return person.status
}

//SECTION ===== ENVIRONMENT STRUCT =====

// Creates Environment ADT which hold all info about the current simulation
type Environment struct {
	currDay              int
	populationSize       int
	nDays                int
	nSusceptible         int
	nRecovered           int
	nVaccinated          int
	nDead                int
	nInfected            int
	vProb                float64
	tProb                float64
	dProb                float64
	currPopulationStatus []Person
	nextPopulationStatus []Person
}

// Decides a random number of exposers and simulates each one to see if a person gets infected or not
func (environment Environment) isInfected() bool {
	isInfected := false

	// Decides number of exposers based on a constant given in the project description
	numExposures := int(NUM_EXPOSERS_MULTIPLIER * rand.Float64())

	// Simulates each exposer to see if the person should be infected
	for j := ZERO; j < numExposures; j++ {

		// Gets a random person from the list to see if they are infected and have the ability to transmit the virus
		randomPerson := environment.currPopulationStatus[rand.Intn(environment.populationSize)]

		// If the random person is infected check if the current person gets infected based on the transmission probability
		if randomPerson.stringifyStatus() == "Infected" {
			if float64(rand.Intn(100)) < environment.tProb {
				isInfected = true
			}
		}
	}

	return isInfected
}

// Generates the next days status array
func (environment Environment) getNextPopulationStatus() []Person {
	// empties the previous nextPopulationStatus array
	environment.nextPopulationStatus = nil

	// Loops through each person in the environment to assign them each a new status
	for i := ZERO; i < environment.populationSize; i++ {

		resolved := false
		// If a person is susceptible, simulate a random number of interactions, if that person gets infect update their new status accordingly
		if environment.currPopulationStatus[i].stringifyStatus() == "Susceptible" {

			if environment.isInfected() {
				environment.nextPopulationStatus = append(environment.nextPopulationStatus, Person{status: NewlyInfected})
				resolved = true
			}

		} else if environment.currPopulationStatus[i].stringifyStatus() == "Infected" {
			// Update the current persons status value by one because they survived another day
			environment.currPopulationStatus[i].status = environment.currPopulationStatus[i].getStatus() + 1

			// If a person stays alive for the designated amount of sick days their status changes to recovered, otherwise they have a certain chance of dying or staying alive but still infected
			if environment.currPopulationStatus[i].getStatus() >= SICK_DAYS {
				environment.nextPopulationStatus = append(environment.nextPopulationStatus, Person{status: Recovered})
				resolved = true
			} else if float64(rand.Intn(100)) < environment.dProb {
				environment.nextPopulationStatus = append(environment.nextPopulationStatus, Person{status: Dead})
				resolved = true
			}
		}

		// If the person's status does not get updated by this point they keep the same status as the previous day
		if !resolved {
			environment.nextPopulationStatus = append(environment.nextPopulationStatus, environment.currPopulationStatus[i])
		}

	}

	return environment.nextPopulationStatus
}

func (environment Environment) initEnvironment(multiTrials bool) ([]Person, []Person, int, int, int, int, int, int) {
	// Loops through each person in the population to assign them a status
	for i := ZERO; i < environment.populationSize; i++ {
		// Checks if a random number out of 100 is less than the probability of starting with a certain status and if so the current person is create with that status
		if rand.Intn(100) < INITIAL_INFECTION_PERCENT {
			environment.currPopulationStatus = append(environment.currPopulationStatus, Person{status: NewlyInfected})
		} else if float64(rand.Intn(100)) < environment.vProb {
			environment.currPopulationStatus = append(environment.currPopulationStatus, Person{status: Vaccinated})
		} else {
			environment.currPopulationStatus = append(environment.currPopulationStatus, Person{status: Susceptible})
		}
	}

	// Gets the status array for the next day and creates a summary of the current day status array
	environment.nextPopulationStatus = environment.getNextPopulationStatus()
	environment.nSusceptible, environment.nRecovered, environment.nVaccinated, environment.nDead, environment.nInfected = environment.countStatus(multiTrials)

	return environment.currPopulationStatus, environment.nextPopulationStatus, environment.populationSize, environment.nSusceptible, environment.nRecovered, environment.nVaccinated, environment.nDead, environment.nInfected
}

// Prints a summary of the people's status on the current day
func (environment Environment) printPopulationSummary(nSusceptible int, nRecovered int, nVaccinated int, nDead int, nInfected int) {
	color.HEX(WHITE).Println("=====  DAY ", environment.currDay, " =====")
	color.HEX(BLUE).Println("   Susceptible: %", percent.PercentOf(nSusceptible, environment.populationSize))
	color.HEX(GREEN).Println("   Recovered:   %", percent.PercentOf(nRecovered, environment.populationSize))
	color.HEX(GREEN).Println("   Vaccinated:  %", percent.PercentOf(nVaccinated, environment.populationSize))
	color.HEX(RED).Println("   Dead:        %", percent.PercentOf(nDead, environment.populationSize))
	color.HEX(ORANGE).Println("   Infected:    %", percent.PercentOf(nInfected, environment.populationSize))

}

// Counts the amount of people with each status on the current day
func (environment Environment) countStatus(printSummary bool) (int, int, int, int, int) {
	// Init each status counter as zero
	nSusceptible, nRecovered, nVaccinated, nDead, nInfected := ZERO, ZERO, ZERO, ZERO, ZERO

	// Loop through the status array and update the counter for each appropriate status
	for i := ZERO; i < environment.populationSize; i++ {
		switch environment.currPopulationStatus[i].getStatus() {
		case Susceptible:
			nSusceptible = nSusceptible + 1
		case Recovered:
			nRecovered = nRecovered + 1
		case Vaccinated:
			nVaccinated = nVaccinated + 1
		case Dead:
			nDead = nDead + 1
		default:
			nInfected = nInfected + 1
		}
	}

	// Print the summary for the current day
	if printSummary {
		environment.printPopulationSummary(nSusceptible, nRecovered, nVaccinated, nDead, nInfected)
	}

	return nSusceptible, nRecovered, nVaccinated, nDead, nInfected
}

// Updates all environment variables that change when the day changes
func (environment Environment) updateDay(printSummary bool) ([]Person, []Person, int, int, int, int) {
	// Copies over the next status array to the current and generates a new next status array
	copy(environment.currPopulationStatus, environment.nextPopulationStatus)
	copy(environment.nextPopulationStatus, environment.getNextPopulationStatus())

	//Prints a summary of the day and updates the nRecovered and nDead variables for that day
	environment.nSusceptible, environment.nRecovered, _, environment.nDead, environment.nInfected = environment.countStatus(printSummary)

	return environment.currPopulationStatus, environment.nextPopulationStatus, environment.nRecovered, environment.nDead, environment.nSusceptible, environment.nInfected
}

// Opens the specified file name
func (environment Environment) openFile(outputFile string) *os.File {
	// Open/create the output file
	pFile, err := os.Create(outputFile)

	// Check for an error when opening/creating output file
	if err != nil {
		fmt.Println(err)
		return os.NewFile(ZERO, "Error")
	}

	return pFile
}

// Writes the current days status information to the output file
func (environment Environment) writeDayToFile(pFile *os.File) {
	// Create output string from current days values and write it to the output file\
	strs := []string{strconv.Itoa(environment.nDays), strconv.Itoa(environment.nSusceptible), strconv.Itoa(environment.nRecovered), strconv.Itoa(environment.nVaccinated), strconv.Itoa(environment.nInfected), strconv.Itoa(environment.nDead), "\n"}
	output := strings.Join(strs, ",")
	_, err := pFile.WriteString(output)

	// Check for an error writing the string to the file
	if err != nil {
		fmt.Println(err)
		pFile.Close()
		return
	}
}

// Closes a given file
func (environment Environment) closeFile(pFile *os.File) {
	err := pFile.Close()

	// Check for an error closing the file
	if err != nil {
		fmt.Println(err)
		return
	}
}

//SECTION ===== DRIVER FUNCTION =====

// Takes a list of args and overwrites the default environment variables as necessary
func parseCmdLine(args []string) (int, float64, float64, float64, int, string, bool) {
	// Defines default environment variables
	popSize, vProb, tProb, dProb, nDays := POPULATION_SIZE, V_PROB, T_PROB, D_PROB, N_DAYS
	outputFile := "output.txt"
	multiTrials := true

	// If special values are passed in, loop through args and overwrite the specified values
	if len(args) > ZERO {
		for i := ZERO; i < len(args); i++ {
			if strings.ToUpper(args[i][0:4]) == "POP=" {
				popSize, _ = strconv.Atoi(args[i][4:len(args[i])])
			} else if strings.ToUpper(args[i][0:6]) == "VPROB=" {
				vProb, _ = strconv.ParseFloat(args[i][6:len(args[i])], 64)
			} else if strings.ToUpper(args[i][0:6]) == "TPROB=" {
				tProb, _ = strconv.ParseFloat(args[i][6:len(args[i])], 64)
			} else if strings.ToUpper(args[i][0:6]) == "DPROB=" {
				dProb, _ = strconv.ParseFloat(args[i][6:len(args[i])], 64)
			} else if strings.ToUpper(args[i][0:5]) == "DAYS=" {
				nDays, _ = strconv.Atoi(args[i][5:len(args[i])])
			} else if strings.ToUpper(args[i][0:5]) == "FILE=" {
				outputFile = args[i][5:len(args[i])]
			} else if strings.ToUpper(args[i][0:6]) == "MULTI=" {
				multiTrials, _ = strconv.ParseBool(args[i][6:len(args[i])])
			}
		}
	}

	return popSize, vProb, tProb, dProb, nDays, outputFile, multiTrials
}

// Final Summary of simulation
func printFinalSummary(environment Environment) {
	color.HEX(WHITE).Println("\n\n\n================== Final SUMMARY ==================")
	color.HEX(WHITE).Println("     Population:               ", environment.populationSize)
	color.HEX(WHITE).Println("     Vaccination Probability:  ", environment.vProb)
	color.HEX(WHITE).Println("     Transmission Probability: ", environment.tProb)
	color.HEX(WHITE).Println("     Death Probability:        ", environment.dProb)
	color.HEX(WHITE).Println("     Initial Infections:       ", environment.nInfected)
	color.HEX(WHITE).Println("     Simulation Period:        ", environment.nDays, "Days")
	color.HEX(WHITE).Println("     Number Recovered:         ", environment.nRecovered)
	color.HEX(WHITE).Println("     Number Dead:              ", environment.nDead)
	color.HEX(WHITE).Println("     Case Fatality Rate:       ", percent.PercentOf(environment.nDead, environment.populationSize), "Percent")
}

// Driver function which updates the simulation for a certain number of days and prints a summary
func Simulate(populationSize int, vProb float64, tProb float64, dProb float64, nDays int, outputFile string, multiTrials bool) Environment {
	// Create Environment and local environment variables
	var environment Environment
	var currPopStatus []Person
	var nextPopStatus []Person

	// Init environment variables
	environment.populationSize = populationSize
	environment.vProb = vProb
	environment.tProb = tProb
	environment.dProb = dProb
	environment.nDays = nDays
	environment.currDay = ZERO
	environment.currPopulationStatus, environment.nextPopulationStatus, environment.populationSize, environment.nSusceptible, environment.nRecovered, environment.nVaccinated, environment.nDead, environment.nInfected = environment.initEnvironment(!multiTrials)

	// Open output file
	pFile := environment.openFile(outputFile)

	// Update the current and next status arrays for each day and print the changes
	for i := ZERO; i < nDays; i++ {

		// Update environment variables
		environment.currDay = i + 1
		currPopStatus, nextPopStatus, environment.nRecovered, environment.nDead, environment.nSusceptible, environment.nInfected = environment.updateDay(!multiTrials)
		copy(environment.currPopulationStatus, currPopStatus)
		copy(environment.nextPopulationStatus, nextPopStatus)

		// Write current days data to output file
		environment.writeDayToFile(pFile)

		// Check if the virus has already run its course
		if environment.nSusceptible == ZERO || environment.nInfected == ZERO {
			if !multiTrials {
				color.HEX(WHITE).Println("\n\n      Simulation ended after", environment.currDay, "days \n The virus has run its course in Nantucket.")
			}
			break
		}
	}

	// Close output file
	environment.closeFile(pFile)

	// Display a final summary of the simulation
	if !multiTrials {
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
	for i := 0; i < TRIALS; i++ {
		deaths = append(deaths, float64(environments[i].nDead))
		recovered = append(recovered, float64(environments[i].nRecovered))
		susceptible = append(susceptible, float64(environments[i].nSusceptible))
		infected = append(infected, float64(environments[i].nInfected))
	}

	color.HEX(WHITE).Println("\n\n\n=============== Multi-Trial Summary ===============\n")
	color.HEX(WHITE).Println("Number of Trials:             ", TRIALS, "\n")
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

func init() {
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
		popSize, vProb, tProb, dProb, nDays, outputFile, multiTrials := parseCmdLine(args)

		var environments []Environment

		if multiTrials {
			for i := 0; i < TRIALS; i++ {
				environments = append(environments, Simulate(popSize, vProb, tProb, dProb, nDays, outputFile, multiTrials))
				color.HEX(GREEN).Println("RUNNING SIMULATION #", i+1, "| Deaths: ", environments[i].nDead, "| Recoveries: ", environments[i].nRecovered, "| Infected: ", environments[i].nInfected, "| Susceptible: ", environments[i].nSusceptible, "|")
			}

			printMultiTrialSummary(environments)
		} else {
			Simulate(popSize, vProb, tProb, dProb, nDays, outputFile, multiTrials)
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
