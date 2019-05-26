package main

import (
	"fmt"
	"github.com/tharindu-wj/golang-couchdb/couchBase"
	"gopkg.in/couchbase/gocb.v1"
)

func main() {
	bucket := couchBase.SelectBucket("company")
	myQuery := gocb.NewN1qlQuery("SELECT * FROM `company` LIMIT 4")
	rows, _ := bucket.ExecuteN1qlQuery(myQuery, nil)

	row := make(map[string]string)
	for rows.Next(&row) {
		fmt.Printf("Row: %+v\n", row)
	}

}
