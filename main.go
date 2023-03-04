package main

import (
	"Assignemnt1/handler"
	"log"
	"net/http"
	"os"
	"time"
)

var StartTime time.Time

func main() {

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	handler.StartTime = time.Now()

	http.HandleFunc("/", handler.DefaultHandler)
	http.HandleFunc("/unisearcher/v1/uniinfo/", handler.HandleGetRequestUni)
	http.HandleFunc("/unisearcher/v1/neighbourunis/", handler.HandleGetRequestNeighborUni)
	http.HandleFunc("/unisearcher/v1/diag/", handler.HandleGetRequestDiag)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
