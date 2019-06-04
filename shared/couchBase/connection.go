package couchBase

import (
	"encoding/json"
	"gopkg.in/couchbase/gocb.v1"
	"io/ioutil"
)

type Auth struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//couch cluster connection
func Connection() *gocb.Cluster {

	plan, _ := ioutil.ReadFile("auth.json")
	var authentication Auth
	err := json.Unmarshal(plan, &authentication)

	cluster, err := gocb.Connect(authentication.Host)
	if err != nil {
		panic(err.Error())
	}
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: authentication.Username,
		Password: authentication.Password,
	})
	if err != nil {
		panic(err.Error())
	}

	return cluster
}
