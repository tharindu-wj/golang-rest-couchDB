package main

import (
	"../couchBase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"net/http"
	"strings"
)

type Company struct {
	RestaurantID   string `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
	CurrencyCode   string `json:"currency_code"`
	BranchID       string `json:"branch_id"`
	BranchName     string `json:"branch_name"`
}

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
	branchID := mux.Vars(r)["id"]

	query := gocb.NewN1qlQuery("SELECT company.* FROM company WHERE branch_id = $1")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{branchID})

	// Interfaces for handling streaming return values
	var row Company
	var retValues []Company

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = Company{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	w.Write(bytes)
}

//get companies by multiple 'branch_ids'
func GetCompaniesByBranchIDs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	responseIds := r.Form.Get("ids")
	branchIDs := strings.Split(responseIds, ",")

	query := gocb.NewN1qlQuery("SELECT company.* FROM company WHERE branch_id IN $1")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{branchIDs})

	// Interfaces for handling streaming return values
	var row Company
	var retValues []Company

	// Stream the values returned from the query into a typed array
	//  of structs
	for rows.Next(&row) {
		retValues = append(retValues, row)
		row = Company{}
	}

	// Marshal array of structs to JSON
	bytes, err := json.Marshal(retValues)
	if err != nil {
		fmt.Println("ERROR PROCESSING STREAMING OUTPUT:", err)
	}

	w.Write(bytes)
}
