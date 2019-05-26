package utils

import (
	"github.com/tharindu-wj/golang-couchdb/couchBase"
	"github.com/tharindu-wj/golang-couchdb/retuarantImporter/model"
	"gopkg.in/couchbase/gocb.v1"
)

func InsertCompanies(data model.Restaurants) bool {
	var companies []gocb.BulkOp

	for _, v := range data {
		_key := v.RestaurantID + v.BranchID
		_val := model.Company{
			RestaurantID:   v.RestaurantID,
			RestaurantName: v.RestaurantName,
			CurrencyCode:   v.CurrencyCode,
			BranchID:       v.BranchID,
			BranchName:     v.BranchName,
		}
		companies = append(companies, &gocb.InsertOp{Key: _key, Value: &_val})
	}

	bucket := couchBase.SelectBucket("company")

	err := bucket.Do(companies)

	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}

func InsertGeoLocations(data model.Restaurants) bool {
	var geoLocations []gocb.BulkOp

	for _, v := range data {
		_key := v.RestaurantID + v.BranchID
		_val := model.Geo{
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
		}
		geoLocations = append(geoLocations, &gocb.InsertOp{Key: _key, Value: &_val})
	}

	bucket := couchBase.SelectBucket("geo")

	err := bucket.Do(geoLocations)

	if err != nil {
		panic(err.Error())
	} else {
		return true
	}
}
