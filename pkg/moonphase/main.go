package moonphase

import (
	"fmt"
	"math"
	"time"
)

func CalculateMoonPhase(date time.Time) Phase {
	jd := julianDate(date)
	knownNewMoon := 2451550.1 // Known new moon date (January 6, 2000)
	daysSinceNewMoon := jd - knownNewMoon
	synodicMonth := 29.53058867
	phaseFloat := (daysSinceNewMoon / synodicMonth) - math.Floor(daysSinceNewMoon/synodicMonth)

	return ParsePhaseFloat(phaseFloat)
}

func FindDateOfPhase(startDate time.Time, phase Phase) (time.Time, error) {
	const searchInterval = 30 // Search within 30 days

	for i := 0; i <= searchInterval; i++ {
		date := startDate.AddDate(0, 0, i)
		if CalculateMoonPhase(date) == phase {
			return date, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not find the phase within %d days", searchInterval)
}

func julianDate(date time.Time) float64 {
	year := float64(date.Year())
	month := float64(date.Month())
	day := float64(date.Day()) + float64(date.Hour())/24.0 + float64(date.Minute())/1440.0 + float64(date.Second())/86400.0

	if month <= 2 {
		year -= 1
		month += 12
	}

	a := math.Floor(year / 100)
	b := 2 - a + math.Floor(a/4)

	return math.Floor(365.25*(year+4716)) + math.Floor(30.6001*(month+1)) + day + b - 1524.5
}
