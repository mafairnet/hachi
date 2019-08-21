package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var configuration = getProgramConfiguration()

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
