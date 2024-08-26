# golang restful service
An example of a restful service written in golang using Domain driven design and onion architecture

## tech stack

### carrier-service api
- fiber for handling http requests
- gorm for ORM
- PostgreSql for data persistance
- Docker

### carrier-client cli
- flag package in Go for command-line parsing

## Getting started

### carrier-service api
In order to run the carrier-service api you must have docker installed locally on your machine to install and run PostgreSql image. 
As an option. You can also build and run the carrier-service api in docker, but it is not required. You can choose to run the
service locally and use Postman or curl to interact with the service.  

Make sure you execute go and docker commands in ```.\golangrestfulservice\service\src\carrier-service```

1. Pull down latest image for Postgresql ```docker pull postgres```
   - If you plan on running the carrier-service locally in an ide or build an executable to run locally you can skip the 2nd step and just run the following command
   - ```docker run --rm -d --name carrierdb -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e postgres -p 5432:5432  postgres```
2. [Optional] Create a bridge network in docker ```docker network create -d bridge dev-network``` if you want to add the carrier-service to docker.
   - Run your PostgreSql in docker with the newly created network
   - ```docker run --rm -d --name carrierdb --network dev-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e postgres -p 5432:5432  postgres```
3. Run your carrier-service api from an ide such as goland.
   - Navigate to main.go in ```.\golangrestfulservice\service\src\carrier-service``` and build and run using the ide
4. [Optional] Run your carrier-service api in a docker container. Make sure you followed step 2. [Optional].
   - Navigate to ```.\golangrestfulservice\service\src\carrier-service```
   - Build your carrier-service image ```docker build -t carrierapi .```
   - Run your new carrier-service image and make sure you add it to the docker bridge network you created in step 2 ```docker run --rm -d --name carrierapi --network dev-network -p 8080:8080 carrierapi```

### carrier-client cli
If you plan on using the cli then you will have to run it from a terminal executing ```main.go```. You will want to ensure that you have the 
carrier-service api running in a docker container or running locally in an ide
1. Navigate to ```.\golangrestfulservice\client\src\carrier-client```
2. Execute command-line functions from your terminal. Execute this one to get started. ```go run main.go```
   - Authenticate: ```go run main.go -f authenticate -u jack -p burton```
   - Get: ```go run main.go -f get -i 1```
   - Create: ```go run main.go -f create -cn "Test Carrier" -add "123 Test Street 12345" -act true```
   - Update: ```go run main.go -f update -add "Update address" -i 1``` or ```go run main.go -f update -act false -i 1```
   - Delete: ```go run main.go -f delete -i 1```

### postman
- authenticate: ```POST: http://localhost:8080/api/authenticate/ {
  "User": "jack",
  "Password":"burton"
  }```
- Read a carrier: ```GET: http://localhost:8080/api/carriers/1 "Authorization" bearer, "Content-Type" application/json```
- Create a carrier: ```POST: http://localhost:8080/api/carriers "Authorization" bearer, "Content-Type" application/json {
  "name":"Bandito Trucking",
  "address":"9458 Outback way, Syndney, AU",
  "active":false
  }```
- Update a carrier address: ```PUT: http://localhost:8080/api/carriers/2/address "Authorization" bearer, "Content-Type" application/json {
  "id":2,
  "address":"1313 Chinatown San Francisco, CA, 90210"
  }```
- Update a carrier Active Status: ```PUT: http://localhost:8080/api/carriers/2/active "Authorization" bearer, "Content-Type" application/json {
  "id":2,
  "active":false
  }```
- Delete a carrier: ```DELETE: http://localhost:8080/api/carriers/1```





## Docker
- ```docker run --rm -d --name carrierdb --network dev-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -e postgres -p 5432:5432  postgres```
- ```docker run --rm -d --name carrierapi --network dev-network -p 8080:8080 carrierapi```
- ```docker build -t carrierapi .```
- ```docker network create -d bridge dev-network```
- ```docker pull postgres```

