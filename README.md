
# Project Setup

## Required Packages
    * gorilla/mux : `go get -u github.com/gorilla/mux`
    * couchbase/gocb : `go get gopkg.in/couchbase/gocb.v1`
    
## Couchbase Setup
    * Need to create two buckets in couchbase `company` and `geo`
    
    
# Project Deployment
Project consist with 3 binary files for each service. These files located in '/bin' folder. Before run those add
couchbase credentials in `auth.json`
    
    * `companyService` : Run server on `localhost:3000` listen for company services
    * `geoService` : Run server on `localhost:3001` listen for geo services
    * `restuarantImporter` : Import resturants data to couchDB. This needs valid json file with restuarants.
       Example json can be found in `/resources/retuarants.json`


# KrakenD Integration

    * Run Krakend with `/resources/krakend.json`. Before run this `company` and `geo` services should be executed and 
      open their ports
    
    
# API Endpoints

## Company Service
- http://localhost:3000/api/v1/company//{id}
    - `GET`: get company by id
    
- http://localhost:3000/api/v1/companies?ids={[ids]}
    - `GET`:get companies by multiple id's. ex:-`["12345678","14723698"]`
    
## Geo Service
- http://localhost:3001/api/v1/geo?lon={lon}&lat={lat}&radius={radius}
    - `GET`: get list of geo locations by certain range. ex:-`lon=-86.79113&lat=32.806671&radius=1500`
    
    
## Krakend
- http://localhost:8000/findNearbyRestaurants/{lon}/{lat}/{radius}
    - `GET`: get company details near to given geo range.
    
    
# Development
    * All services source codes in `/src` folder
    * Common package helper functions and models are in `shared` folder
