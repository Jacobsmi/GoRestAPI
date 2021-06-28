# Car Tracker API

## Setting up the Database

- Step One:
  - Create a PostreSQL database that the API can connect to and persist data in
- Step Two:
  - Create a Go file in the database directory called `databaseVars.go` with the following variables
    -	host, port, user, password, dbname  
    - Example `databaseVars.go` file
        ```go
        package database

        const (
	        host     = "localhost"
	        port     = 5432
	        user     = "postgres"
	        password = "Password"
	        dbname   = "go_rest_api"
        )  
        ```

## Running the Server

- cd into the goRestAPI directory and `go run main.go` to run the server