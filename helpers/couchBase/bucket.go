package couchBase

import "gopkg.in/couchbase/gocb.v1"

var cluster = Connection()

func SelectBucket(bucketName string) *gocb.Bucket {

	bucket, err := cluster.OpenBucket(bucketName, "")

	if err != nil {
		panic(err.Error())
	}

	err = bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	if err != nil {
		panic(err.Error())
	}

	return bucket
}
