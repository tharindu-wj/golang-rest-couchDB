package couchBase

import "gopkg.in/couchbase/gocb.v1"

func CreateDocument(bucket *gocb.Bucket, key string, data interface{}) bool {
	_, err := bucket.Upsert(key, data, 0)

	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}
