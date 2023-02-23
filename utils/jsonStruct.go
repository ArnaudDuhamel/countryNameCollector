package utils

// CountryNames is the data structure used to gather and store the country names
type CountryNames struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}
