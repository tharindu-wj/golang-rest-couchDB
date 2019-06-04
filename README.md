
# Project Setup

## Required Packages
    * gorilla/mux : `go get -u github.com/gorilla/mux`
    * couchbase/gocb : `go get gopkg.in/couchbase/gocb.v1`
    
## Couchbase Setup
    * Need to create two buckets in couchbase `company` and `geo`
    
# Project Deployment

## 



# API Endpoints
###Posts
- http://localhost:3000/api/v1/posts
    - `GET`: get posts
    
- http://localhost:3000/api/v1/post/{id}
    - `GET`: get post
    
- http://localhost:3000/api/v1/post/create
    - `POST`: create post
    
- http://localhost:3000/api/v1/post/delete/{id}
    - `DELETE`: delete post
    
###Companies
- http://localhost:3000/api/v1/companies
    - `GET`: get list of companies
    
###Users
- http://localhost:3000/api/v1/users
    - `GET`: get users
    
- http://localhost:3000/api/v1/user/create
    - `POST`: create user
  