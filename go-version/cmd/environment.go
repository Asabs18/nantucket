package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/dariubs/percent"
	"github.com/gookit/color"
)

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
func (environment Environment) IsInfected() bool {
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
func (environment *Environment) getNextPopulationStatus() {
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

}

func (environment *Environment) initEnvironment(multiTrials bool) {
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
	environment.getNextPopulationStatus()
	environment.countStatus(multiTrials)

}

// Prints a summary of the people's status on the current day
func (environment Environment) printPopulationSummary() {
	color.HEX(WHITE).Println("=====  DAY ", environment.currDay, " =====")
	color.HEX(BLUE).Println("   Susceptible: %", percent.PercentOf(environment.nSusceptible, environment.populationSize))
	color.HEX(GREEN).Println("   Recovered:   %", percent.PercentOf(environment.nRecovered, environment.populationSize))
	color.HEX(GREEN).Println("   Vaccinated:  %", percent.PercentOf(environment.nVaccinated, environment.populationSize))
	color.HEX(RED).Println("   Dead:        %", percent.PercentOf(environment.nDead, environment.populationSize))
	color.HEX(ORANGE).Println("   Infected:    %", percent.PercentOf(environment.nInfected, environment.populationSize))

}

// Counts the amount of people with each status on the current day
func (environment *Environment) countStatus(printSummary bool) {
	// Reset each counter to
	environment.nSusceptible = 0
	environment.nRecovered = 0
	environment.nVaccinated = 0
	environment.nDead = 0
	environment.nInfected = 0

	// Loop through the status array and update the counter for each appropriate status
	for i := ZERO; i < environment.populationSize; i++ {
		switch environment.currPopulationStatus[i].getStatus() {
		case Susceptible:
			environment.nSusceptible++
		case Recovered:
			environment.nRecovered++
		case Vaccinated:
			environment.nVaccinated++
		case Dead:
			environment.nDead++
		default:
			environment.nInfected++
		}
	}

	// Print the summary for the current day
	if printSummary {
		environment.printPopulationSummary()
	}
}

// Updates all environment variables that change when the day changes
func (environment *Environment) updateDay(printSummary bool) {
	// Copies over the next status array to the current array
	copy(environment.currPopulationStatus, environment.nextPopulationStatus)

	// Generates the status array for the next day
	environment.getNextPopulationStatus()

	// Prints a summary of the day and updates the nRecovered and nDead variables for that day
	environment.countStatus(printSummary)
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
