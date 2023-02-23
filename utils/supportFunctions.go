package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
)

// ObtainCountryNames gathers the country names from the
// "https://restcountries.com/" API
func ObtainCountryNames() ([]string, error) {

	var nameStructList []CountryNames
	var emptyList []string
	var nameList []string

	//Request name information to the API
	url := "https://restcountries.com/v3.1/all"

	newRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error in creating request: " + err.Error())
		return emptyList, errors.New("Error in creating request")
	}

	newRequest.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	requestIssue, err := client.Do(newRequest)
	if err != nil {
		fmt.Println("Did not manage to issue request: " + err.Error())
		return emptyList, errors.New("Did not manage to issue request")
	}

	decoder := json.NewDecoder(requestIssue.Body)

	if err = decoder.Decode(&nameStructList); err != nil {
		fmt.Println("Error in decoding data: " + err.Error())
		return emptyList, errors.New(err.Error())
	}

	//creates a clean slice of strings with the country names
	for _, countryName := range nameStructList {
		nameList = append(nameList, countryName.Name.Common)
	}

	//sorts the list in alphabetical order
	sort.Strings(nameList)

	return nameList, nil

}

// WriteToFile writes the gathered data from the ObtainCountryNames()
// function to a file called "countryNames.go" in a slice of strings
func WriteToFile(countryNamesList []string) error {

	//creates the file
	file, err := os.Create("countryNames.go")
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
		return err
	}
	defer file.Close()

	//Create the data structure and populate it with the
	//country names
	data := []byte("package utils\n\nvar COUNTRYNAMES = []string{\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file: " + err.Error())
		return err
	}

	for index, countryName := range countryNamesList {

		if index != len(countryNamesList)-1 {
			data = []byte("\t" + "\"" + countryName + "\"" + ",\n")
		} else {
			data = []byte("\t" + "\"" + countryName + "\"" + " }")
		}

		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file: " + err.Error())
			return err
		}

	}

	fmt.Println("Data written successfully to file.")
	return nil
}
