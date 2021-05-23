package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"time"
)

// inputInteger, prompt that will validate integer input.
func inputInteger(name string) int {
	d := promptui.Prompt{
		Label:    name,
		Validate: validateNumber,
	}
	input, _ := d.Run()
	result, _ := strconv.Atoi(input)

	return result
}

// inputString, prompt that will validate input type string.
func inputString(name string) string {
	d := promptui.Prompt{
		Label:    name,
		Validate: validateNumber,
	}
	input, _ := d.Run()

	return input
}

// inputDoseNumber, this will allow you to select dose number
// Currently there are 2 dose available.
func inputDoseNumber() int {
	dz := promptui.Select{
		Label: "Select Dose Number",
		Items: []string{"1", "2"},
	}
	selectedDoseIndex, _, _ := dz.Run()

	if selectedDoseIndex == 0 {
		return 1
	}

	return 2
}

// inputVaccine: Input Vaccine.
func inputVaccine() string {
	v := promptui.Select{
		Label: "Select Vaccine",
		Items: []string{"ANY", "COVAXIN", "COVISHIELD"},
	}
	_, vaccine, err := v.Run()

	if err != nil {
		panic("inputVaccine: Invalid Vaccine Name Selected")
	}

	return vaccine
}

// inputAge, prompt to allow select age group.
func inputAge() int {
	a := promptui.Select{
		Label: "Select Age",
		Items: []string{"18", "45"},
	}
	selectedAgeIndex, _, err := a.Run()

	if err != nil {
		panic("inputAge: Invalid Age Selected")
	}

	if selectedAgeIndex == 0 {
		return 18
	}

	return 45
}

// inputStartDate: Prompt that will take input date.
func inputStartDate() string {
	date := promptui.Prompt{
		Label:    fmt.Sprintf("Enter Start Date DD-MM-YYYY, Today Is: %s", time.Now().Format("02-01-2006")),
		Validate: validateDate,
	}
	startDate, err := date.Run()

	if err != nil {
		panic("inputStartDate: Invalid Date Selected")
	}

	return startDate
}
