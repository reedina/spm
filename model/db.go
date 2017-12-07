package model

import (
	"database/sql"
	"fmt"
	"log"

	//Initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//ConnectDB - Establish DB connection
func ConnectDB(user, password, dbname, url string) {
	/*
		connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s",
			user, password, url, dbname)*/
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, dbname)

	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
	}
}
