// Code generated by "stringer -type=StatusTest"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TestNewlyInfected-10]
	_ = x[TestSusceptible-0]
	_ = x[TestRecovered - -11]
	_ = x[TestVaccinated - -22]
	_ = x[TestDead - -34]
}

const (
	_StatusTest_name_0 = "TestDead"
	_StatusTest_name_1 = "TestVaccinated"
	_StatusTest_name_2 = "TestRecovered"
	_StatusTest_name_3 = "TestSusceptible"
	_StatusTest_name_4 = "TestNewlyInfected"
)

func (i StatusTest) String() string {
	switch {
	case i == -34:
		return _StatusTest_name_0
	case i == -22:
		return _StatusTest_name_1
	case i == -11:
		return _StatusTest_name_2
	case i == 0:
		return _StatusTest_name_3
	case i == 10:
		return _StatusTest_name_4
	default:
		return "StatusTest(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
