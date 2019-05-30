package main

import (
	"github.com/tharindu-wj/golang-rest-couchDB/helpers/models"
	"github.com/tharindu-wj/golang-rest-couchDB/helpers/couchBase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"net/http"
	"strconv"
)

//get company bucket reference
var bucket = couchBase.SelectBucket("geo")

func main() {

	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/geo", GetGeoByLocation).Methods("POST")

	log.Fatal(http.ListenAndServe(":3001", router))
}

//get geo locations by given location and distance range
func GetGeoByLocation(w http.ResponseWriter, r *http.Request) {
	// Grab the current location
	r.ParseForm()

	lat, _ := strconv.ParseFloat(r.Form.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.Form.Get("lon"), 64)
	distance, _ := strconv.ParseInt(r.Form.Get("distance"), 10, 64) //in miles

	fmt.Println(lat, lon, distance)
	fmt.Printf("%T, %T, %T", lat, lon, distance)

	// New query, a really generic one with high selectivity
	distanceCalc := fmt.Sprintf("(3959 * acos(cos(radians(%f)) * cos(radians(lat)) * cos( radians(lon) - radians(%f)) + sin(radians(32.816671)) *sin(radians(lat))))", lat, lon)
	queryString := fmt.Sprintf("SELECT geo.lat,geo.lon,META().id, %s AS distance FROM geo where %s < %d ORDER BY distance", distanceCalc, distanceCalc, distance)
	query := gocb.NewN1qlQuery(queryString)

	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{})

	// Interfaces for handling streaming return values
	var row models.Geo
	var retValues []models.Geo

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		// Set the row to an empty struct, to prevent current values
		//  being added to the next row in the results collection
		//  returned by the query
		row = models.Geo{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	// Return the JSON
	w.Write(bytes)
}
