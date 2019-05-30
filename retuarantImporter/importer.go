package main

import (
	"encoding/json"
	"fmt"
	"github.com/tharindu-wj/golang-rest-couchDB/packages/models"
	"github.com/tharindu-wj/golang-rest-couchDB/packages/couchBase"
	"gopkg.in/couchbase/gocb.v1"
	"io/ioutil"
	"os"
)

func main() {
	restaurants := ImportJsonToRestaurants()

	insertCompanies := InsertCompanies(restaurants)
	if (insertCompanies) {
		fmt.Println("Inserted bulk companies")
	}

	insertgeoLocations := InsertGeoLocations(restaurants)
	if (insertgeoLocations) {
		fmt.Println("Inserted bulk geo locations")
	}

}

//insert companies to companies bucket
func InsertCompanies(data models.Restaurants) bool {
	var companies []gocb.BulkOp

	for _, v := range data {
		_key := v.RestaurantID + v.BranchID
		_val := models.Company{
			RestaurantID:   v.RestaurantID,
			RestaurantName: v.RestaurantName,
			CurrencyCode:   v.CurrencyCode,
			BranchID:       v.BranchID,
			BranchName:     v.BranchName,
		}
		companies = append(companies, &gocb.InsertOp{Key: _key, Value: &_val})
	}

	bucket := couchBase.SelectBucket("company")

	err := bucket.Do(companies)

	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}

//insert geo data to geo bucket
func InsertGeoLocations(data models.Restaurants) bool {
	var geoLocations []gocb.BulkOp

	for _, v := range data {
		_key := v.RestaurantID + v.BranchID
		_val := models.Geo{
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
		}
		geoLocations = append(geoLocations, &gocb.InsertOp{Key: _key, Value: &_val})
	}

	bucket := couchBase.SelectBucket("geo")

	err := bucket.Do(geoLocations)

	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}

//read json file from cmd argument and return restaurant map
func ImportJsonToRestaurants() models.Restaurants {
	jsonFileName := os.Args[1]

	plan, _ := ioutil.ReadFile(jsonFileName)
	var outputStruct models.Restaurants
	err := json.Unmarshal(plan, &outputStruct)

	if err != nil {
		fmt.Print(err)
	}
	return outputStruct
}

