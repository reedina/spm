package main

import (
	"os"

	_ "github.com/lib/pq"
)

func main() {

	a := App{}

	a.InitializeApplication(os.Getenv("AWS_DB_USERNAME"),
		os.Getenv("AWS_DB_PASSWORD"),
		os.Getenv("AWS_DB_NAME"),
		os.Getenv("AWS_DB_URL"))

	a.RunApplication(":4040")
}
