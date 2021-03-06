package main

import (
	"github.com/tharindu-wj/golang-rest-couchDB/shared/models"
	"github.com/tharindu-wj/golang-rest-couchDB/shared/couchBase"
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


	subRouter.Path("/geo").Queries("lat", "{lat}").HandlerFunc(GetGeoByLocation).Methods("GET")
	subRouter.Path("/geo").Queries("lon", "{lon}").HandlerFunc(GetGeoByLocation).Methods("GET")
	subRouter.Path("/geo").Queries("radius", "{radius}").HandlerFunc(GetGeoByLocation).Methods("GET")
	subRouter.Path("/geo").HandlerFunc(GetGeoByLocation).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", router))
}

//get geo locations by given location and distance range
func GetGeoByLocation(w http.ResponseWriter, r *http.Request) {
	// Grab the current location

	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
	lon, _ := strconv.ParseFloat(r.FormValue("lon"), 64)
	radius, _ := strconv.ParseInt(r.FormValue("radius"), 10, 64) //in miles

	// New query, a really generic one with high selectivity
	radiusCalc := fmt.Sprintf("(3959 * acos(cos(radians(%f)) * cos(radians(lat)) * cos( radians(lon) - radians(%f)) + sin(radians(32.816671)) *sin(radians(lat))))", lat, lon)
	queryString := fmt.Sprintf("SELECT geo.lat,geo.lon,META().id, %s AS distance FROM geo where %s < %d ORDER BY distance", radiusCalc, radiusCalc, radius)
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
