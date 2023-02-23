package main

import (
	"countryNames/utils"
	"os"
)

// This program gathers the country names and store them to a "countryNames.go" file
// It gathers the country names from this API:
// https://restcountries.com/v3.1/all
// Website of the API: https://restcountries.com/
// GitHub of the API: https://gitlab.com/amatos/rest-countries

func main() {

	// gathers the country names into a string list
	nameList, err := utils.ObtainCountryNames()
	if err != nil {
		os.Exit(1)
	}

	// writes the country names in a string slice
	// in a file called "countryNames.go"
	err = utils.WriteToFile(nameList)
	if err != nil {
		os.Exit(1)
	}

}
