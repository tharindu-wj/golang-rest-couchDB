package couchBase

import (
	"gopkg.in/couchbase/gocb.v1"
)

//couch cluster connection
func Connection() *gocb.Cluster {
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		panic(err.Error())
	}
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "admin",
		Password: "asdf1234",
	})
	if err != nil {
		panic(err.Error())
	}

	return cluster
}
