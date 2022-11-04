package cmd

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

type StatusTest int64

const (
	TestNewlyInfected StatusTest = 10
	TestSusceptible   StatusTest = 0
	TestRecovered     StatusTest = -11
	TestVaccinated    StatusTest = -22
	TestDead          StatusTest = -34
)

//go:generate stringer -type=Status

//go:generate stringer -type=StatusTest

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
