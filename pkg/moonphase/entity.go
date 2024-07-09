package moonphase

import (
	"strings"
)

type Phase struct {
	String string
	Int    int
}

var (
	NewMoon      = Phase{String: "new moon", Int: 0}
	FirstQuarter = Phase{String: "first quarter", Int: 1}
	FullMoon     = Phase{String: "full moon", Int: 2}
	LastQuarter  = Phase{String: "last quarter", Int: 3}

	phases = []Phase{NewMoon, FirstQuarter, FullMoon}
)

func ParsePhaseStr(phaseStr string) Phase {
	phaseStr = strings.TrimSpace(phaseStr)
	phaseStr = strings.ToLower(phaseStr)
	for _, phase := range phases {
		if strings.ToLower(phaseStr) == strings.ToLower(phase.String) {
			return phase
		}
	}
	return Phase{}
}

func ParsePhaseInt(phaseInt int) Phase {
	for _, phase := range phases {
		if phaseInt == phase.Int {
			return phase
		}
	}
	return Phase{}
}

func ParsePhaseFloat(phaseFloat float64) Phase {
	if phaseFloat < 0.125 {
		return NewMoon
	} else if phaseFloat < 0.375 {
		return FirstQuarter
	} else if phaseFloat < 0.625 {
		return FullMoon
	} else if phaseFloat < 0.875 {
		return LastQuarter
	} else {
		return NewMoon
	}
}
