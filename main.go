package main

import (
	"flag"
	"fmt"
	"github.com/HuguesBt/moonphase/pkg/moonphase"
	"time"
)

var (
	action            string
	phaseStr          string
	dateStr           string
	dateTimeFormatStr = "2006-01-02 03:04:05"
	dateFormatStr     = "2006-01-02"
)

func initFlags() {
	flag.StringVar(&action, "a", "", "Action to realize; phase/date")
	flag.StringVar(&phaseStr, "p", "", "Phase to found after date; require date")
	flag.StringVar(&dateStr, "d", time.Now().Format(dateTimeFormatStr), "Date for calcul")
	flag.Parse()
}

func main() {
	initFlags()

	switch action {
	case "phase":
		if dateTimeParsed, err := time.Parse(dateTimeFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else {
			moonPhaseObj := moonphase.CalculateMoonPhase(dateTimeParsed)
			fmt.Println(fmt.Sprintf("The phase is %s for date %s", moonPhaseObj.String, dateTimeParsed.Format(dateFormatStr)))
		}
		break
	case "date":
		if phaseObj := moonphase.ParsePhaseStr(phaseStr); phaseObj.String == "" {
			fmt.Println("Could not parse phase:", phaseStr)
		} else if dateTimeParsed, err := time.Parse(dateTimeFormatStr, dateStr); err != nil {
			fmt.Println("Could not parse time:", err)
		} else if moonPhaseDate, err := moonphase.FindDateOfPhase(dateTimeParsed, phaseObj); err != nil {
			fmt.Println("Could not find date:", err)
		} else {
			fmt.Println(fmt.Sprintf("The next %s is for date %s", phaseObj.String, moonPhaseDate.Format(dateFormatStr)))
		}
		break
	default:
		fmt.Println("Invalid action")
	}
}
