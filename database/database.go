package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectToDB() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)
}

func CreateUser(name string, username string, password string) {
	connectToDB()
	fmt.Println("Creating a user")
	db, err := connectToDB()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`

	_, err = db.Exec(sqlStatement, name, username, password)

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
