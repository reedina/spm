package model

import (
	"database/sql"
	"fmt"
	"log"

	//Initialize postgreSQL provider
	_ "github.com/lib/pq"
)

var db *sql.DB

//ConnectDB - Establish DB connection
func ConnectDB(user, password, dbname, url string) {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s",
		user, password, url, dbname)

	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
	}
}
