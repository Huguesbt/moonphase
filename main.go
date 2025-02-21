package main

import (
	"flag"
	"fmt"
	"github.com/HuguesBt/moonphase/pkg/moonphase"
	"time"
)

var (
	action        string
	phaseStr      string
	dateStr       string
	dateFormatStr = "2006-01-02"
)

func initFlags() {
	flag.StringVar(&action, "a", "", "Action to realize; phase/date")
	flag.StringVar(&phaseStr, "p", "", "Phase to found after date; require date")
	flag.StringVar(&dateStr, "d", time.Now().Format(dateFormatStr), "Date for operation")
	flag.Parse()
}

func main() {
	initFlags()

	switch action {
	case "phase":
		if phaseObj := moonphase.ParsePhaseStr(phaseStr); phaseObj.String == "" {
			fmt.Println("Could not parse phase:", phaseStr)
		} else if dateTimeParsed, err := time.Parse(dateFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else if moonPhaseDate, err := moonphase.FindDateOfPhase(dateTimeParsed, phaseObj); err != nil {
			fmt.Println("Could not find date:", err)
		} else {
			fmt.Println(fmt.Sprintf("The next %s is for date %s", phaseObj.String, moonPhaseDate.Format(dateFormatStr)))
		}
		break
	case "date":
		if dateTimeParsed, err := time.Parse(dateFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else {
			moonPhaseObj := moonphase.CalculateMoonPhase(dateTimeParsed)
			fmt.Println(fmt.Sprintf("The phase is %s for date %s", moonPhaseObj.String, dateTimeParsed.Format(dateFormatStr)))
		}
		break
	case "calendar":
		if dateTimeParsed, err := time.Parse(dateFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else {
			calStr := calendar(dateTimeParsed.Year())
			fmt.Println(calStr)
		}
		break
	case "calendar-phases":
		if dateTimeParsed, err := time.Parse(dateFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else {
			calStr := calendarPhases(dateTimeParsed)
			fmt.Println(calStr)
		}
		break
	default:
		fmt.Println("Invalid action")
	}
}

func calendar(year int) (ret string) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	ret += fmt.Sprintf("Liste for year %d\n", year)
	ret += fmt.Sprintf("%s, %s\n", "Date", "Moon phase")

	for d := start; d.Unix() <= end.Unix(); d = d.AddDate(0, 0, 1) {
		moonPosition := moonphase.CalculateMoonPhase(d)

		ret += fmt.Sprintf(
			"%s, %s\n",
			d.Format("2006-01-02"),
			moonPosition.String,
		)
	}

	return
}

func calendarPhases(dateTimeParsed time.Time) (ret string) {
	ret += fmt.Sprintln("Next phases")

	for _, phaseObj := range []moonphase.Phase{moonphase.FullMoon, moonphase.FirstQuarter, moonphase.LastQuarter, moonphase.NewMoon} {
		if moonPhaseDate, err := moonphase.FindDateOfPhase(dateTimeParsed, phaseObj); err != nil {
			ret += fmt.Sprintln("Could not find date:", err)
		} else {
			ret += fmt.Sprintf("The next %s is for date %s\n", phaseObj.String, moonPhaseDate.Format(dateFormatStr))
		}
	}

	return
}
