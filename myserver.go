package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// default route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "My Cats API")
	})

	// return a `catdata.json` file for `/mycats` route
	http.HandleFunc("/mycats", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./catdata.json")
	})

	log.Println("Starting server...")
	// start HTTP server with `http.DefaultServeMux` handler
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("Unable to start HTTP server: %s", err)
		os.Exit(1)
	}
}
