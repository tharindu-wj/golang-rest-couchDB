package utils

import (
	"encoding/json"
	"fmt"
	"github.com/tharindu-wj/golang-couchdb/retuarantImporter/model"
	"io/ioutil"
	"os"
)

func ImportJsonToRestaurants() model.Restaurants {
	jsonFileName := os.Args[1]

	plan, _ := ioutil.ReadFile(jsonFileName)
	var outputStruct model.Restaurants
	err := json.Unmarshal(plan, &outputStruct)

	if err != nil {
		fmt.Print(err)
	}
	return outputStruct
}
