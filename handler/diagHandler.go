package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var StartTime time.Time

/*
Handler for the diag
*/
func HandleGetRequestDiag(w http.ResponseWriter, r *http.Request) {

	// Doing get request for universitiesapi
	url1, err := http.Get("http://universities.hipolabs.com/")
	if err != nil {
		fmt.Errorf("Error in get request:", err.Error())
	}

	// Doing get request for countriesapi
	url2, err := http.Get("https://restcountries.com/")
	if err != nil {
		fmt.Errorf("\nError in get request:", err.Error())
	}

	// Calculate how long the service has been running
	elapsed := time.Since(StartTime)

	data := DIAGNOSTIC{
		Universitiesapi: url1.Status,
		Countriesapi:    url2.Status,
		Version:         "v1",
		Uptime:          elapsed.String(),
	}

	output, _ := json.MarshalIndent(data, "", " ")
	fmt.Fprintf(w, "%v", string(output))

}
