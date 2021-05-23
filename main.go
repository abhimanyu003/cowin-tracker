package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gen2brain/beeep"
	"strings"
	"time"
)

const findByPinURL = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=%d&date=%s"
const defaultScanInterval = 60 // 1 minute

// App Base structure of app.
type App struct {
	Cowin        CoWin
	PinCode      int
	Age          int
	ScanDays     int
	StartDate    string
	ScanInterval int
	DoseNumber   int
	VaccineName  string
}

func main() {
	var app App
	app.PinCode = inputInteger("Enter Pin Code")
	app.Age = inputAge()
	app.DoseNumber = inputDoseNumber()
	app.VaccineName = inputVaccine()
	app.StartDate = inputStartDate()
	app.ScanDays = inputInteger("Enter number ScanDays to scan")
	app.ScanInterval = inputInteger("Enter Scan Interval in seconds, default is 60 sec.")

	// Default Scan Interval
	if app.ScanInterval == 0 {
		app.ScanInterval = defaultScanInterval
	}

	ttl, err := time.Parse("02-01-2006", app.StartDate)

	if err != nil {
		panic("Invalid Date Provided, Please Use DD-MM-YYY")
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Color("green")
	s.Prefix = "Keep this windows open, locating vaccine!! You will get notification once available."
	s.Start()

	for {
		for i := 0; i <= app.ScanDays; i++ {
			date := ttl.AddDate(0, 0, i).Format("02-01-2006")

			url := fmt.Sprintf(findByPinURL, app.PinCode, date)

			// We will ignore error to continue the loop
			GetJson(url, &app.Cowin)

			for _, session := range app.Cowin.Sessions {
				if app.isVaccineAvailable(&session, date) {

					alertMsg := fmt.Sprintf("Vaccine are now available at %s for date %s", session.Name, date)

					beeep.Alert("Co-Win Alert", alertMsg, "")
				}
			}
		}
		time.Sleep(time.Duration(app.ScanInterval) * time.Second)
	}
	s.Stop()
}

func (app *App) isVaccineAvailable(session *Sessions, date string) bool {
	var isAvailable bool

	var availableCapacityDose int

	// Choose if doze one or two selected
	if app.DoseNumber == 1 {
		availableCapacityDose = session.AvailableCapacityDose1
	} else {
		availableCapacityDose = session.AvailableCapacityDose2
	}

	if availableCapacityDose <= 0 {
		return false
	}

	if session.MinAgeLimit == app.Age && session.Date == date {
		isAvailable = true
		if app.VaccineName != "ANY" {
			isAvailable = strings.EqualFold(session.Vaccine, app.VaccineName)
		}
	}

	return isAvailable
}
