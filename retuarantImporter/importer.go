package main

import (
	"fmt"
	"./utils"
)

func main() {
	restaurants := utils.ImportJsonToRestaurants()

	insertCompanies := utils.InsertCompanies(restaurants)
	if (insertCompanies) {
		fmt.Println("Inserted bulk companies")
	}

	insertgeoLocations := utils.InsertGeoLocations(restaurants)
	if (insertgeoLocations) {
		fmt.Println("Inserted bulk geo locations")
	}

}
