package couchBase

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/couchbase/gocb.v1"
	"os"
)

//couch cluster connection
func Connection() *gocb.Cluster {
	_ = godotenv.Load(".env")
	username := os.Getenv("COUCHBASE_USERNAME")
	password := os.Getenv("COUCHBASE_PASSWORD")
	fmt.Println(username, password)
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		panic(err.Error())
	}
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	})
	if err != nil {
		panic(err.Error())
	}

	return cluster
}
