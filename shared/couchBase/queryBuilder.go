package couchBase

import (
	"encoding/json"
	"fmt"
	"github.com/tharindu-wj/golang-rest-couchDB/shared/models"
	"gopkg.in/couchbase/gocb.v1"
)

//TODO get result from couchbase query
func QueryBuilder(bucket *gocb.Bucket, queryString string, row interface{}, retValues []interface{}) []uint8 {
	query := gocb.NewN1qlQuery(queryString)
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{})

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

	return bytes
}
