package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}

	/*
		a.InitializeApplication(os.Getenv("AWS_DB_USERNAME"),
			os.Getenv("AWS_DB_PASSWORD"),
			os.Getenv("AWS_DB_NAME"),
			os.Getenv("AWS_DB_URL"))
	*/
	a.InitializeApplication(os.Getenv("GCP_API_DB_USER"),
		os.Getenv("GCP_API_DB_PASSWORD"),
		os.Getenv("GCP_API_DB_NAME"),
		os.Getenv("GCP_API_DB_HOST"))

	a.RunApplication(":4040")
}
