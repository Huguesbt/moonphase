package moonphase

import (
	"fmt"
	"github.com/hablullah/go-sampa"
	"time"
)

var (
	location = sampa.Location{Latitude: 48.85661400, Longitude: 2.35222190}
)

func CalculateMoonPhase(dt time.Time) Phase {
	moonPosition, _ := sampa.GetMoonPosition(dt, location, nil)
	return Phase{
		String: moonPosition.Phase.String(),
	}
}

func FindDateOfPhase(startDate time.Time, phase Phase) (dt time.Time, err error) {
	p := sampa.GetMoonPhases(startDate, nil)
	switch phase.String {
	case NewMoon.String:
		dt = p.NewMoon
	case FirstQuarter.String:
		dt = p.FirstQuarter
	case FullMoon.String:
		dt = p.FullMoon
	case LastQuarter.String:
		dt = p.LastQuarter
	default:
		return time.Time{}, fmt.Errorf("unknown phase")
	}
	if dt.Before(startDate) {
		fmt.Println("reload")
		return FindDateOfPhase(startDate.AddDate(0, 0, 1), phase)
	} else {
		return
	}
}
