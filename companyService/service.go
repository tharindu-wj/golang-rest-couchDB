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
	"strings"
)

//get company bucket reference
var bucket = couchBase.SelectBucket("company")

func main() {

	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	subRouter.HandleFunc("/companies/{id}", GetCompaniesByBranchID).Methods("GET")
	subRouter.HandleFunc("/companies", GetCompaniesByBranchIDs).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}

//get companies by 'branch_id'
func GetCompaniesByBranchID(w http.ResponseWriter, r *http.Request) {

	// Grab the parameter for the company
	branchID := mux.Vars(r)["id"]

	// New query, a really generic one with high selectivity
	query := gocb.NewN1qlQuery("SELECT company.*,META().id FROM company WHERE branch_id = $1")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{branchID})

	// Interfaces for handling streaming return values
	var row models.Company
	var retValues []models.Company

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		// Set the row to an empty struct, to prevent current values
		//  being added to the next row in the results collection
		//  returned by the query
		row = models.Company{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	// Return the JSON
	w.Write(bytes)
}

//get companies by multiple 'branch_ids'
func GetCompaniesByBranchIDs(w http.ResponseWriter, r *http.Request) {
	// Grab the branch_id's for the company
	r.ParseForm()
	responseIds := r.Form.Get("ids")

	//convert branch id's string to []string array
	branchIDs := strings.Split(responseIds, ",")

	// New query, a really generic one with high selectivity
	query := gocb.NewN1qlQuery("SELECT company.*,META().id FROM company WHERE branch_id IN $1")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{branchIDs})

	// Interfaces for handling streaming return values
	var row models.Company
	var retValues []models.Company

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = models.Company{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	// Return the JSON
	w.Write(bytes)
}
