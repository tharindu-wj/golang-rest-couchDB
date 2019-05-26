package queryBuilder

import (
	"fmt"
	"github.com/tharindu-wj/golang-couchdb/couchBase"
	"gopkg.in/couchbase/gocb.v1"
)

func GetSingle(bucketName string, column string, value string){
	bucket := couchBase.SelectBucket(bucketName)
	queryString:= fmt.Sprintf("SELECT * FROM company")
	n1Query := gocb.NewN1qlQuery(queryString)
	rows, _ := bucket.ExecuteN1qlQuery(n1Query, nil)

	var result interface{}
	var row interface{}
	for rows.Next(&result){}
	fmt.Println(row)
}
