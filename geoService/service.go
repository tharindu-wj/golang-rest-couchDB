package main

import (
	"../couchBase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"net/http"
)

type Geo struct {
	ID        string  `json:"id"`
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

//get company bucket reference
var bucket = couchBase.SelectBucket("geo")

func main() {

	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/geo", GetGeoByLocation).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", router))
	/*indexName := "geo-location"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("32.806671")).
		Limit(10)

	res, _ := bucket.ExecuteSearchQuery(query)
	for _, hit := range res.Hits() {
		fmt.Println( hit)
	}*/

}

func GetGeoByLocation(w http.ResponseWriter, r *http.Request) {
	// Grab the parameter for the company
	//branchID := mux.Vars(r)["id"]
	lat := 32.806671
	lon := -86.79113

	// New query, a really generic one with high selectivity
	query := gocb.NewN1qlQuery("SELECT geo.*,META().id FROM geo WHERE lat = $1 AND lon = $2")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{lat, lon})

	// Interfaces for handling streaming return values
	var row Geo
	var retValues []Geo

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		// Set the row to an empty struct, to prevent current values
		//  being added to the next row in the results collection
		//  returned by the query
		row = Geo{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	// Return the JSON
	w.Write(bytes)
}
