package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
Dedicated handler for POST requests
*/
func HandleGetRequestUni(w http.ResponseWriter, r *http.Request) {

	URLparts := strings.Split(r.URL.Path, "/")
	name := URLparts[4]

	url := "http://universities.hipolabs.com/search?name=" + name

	NewRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	NewRequest.Header.Add("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	res, err := client.Do(NewRequest)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// Instantiate decoder
	decoder := json.NewDecoder(res.Body)
	// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
	// Note: This does not lead to a check whether an actually provided field is empty!

	// Prepare empty struct to populate
	uni := []UNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&uni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Flat printing
	fmt.Println(uni)
	/*
		fmt.Println("Received following request:")
		fmt.Println(uni)
		fmt.Fprintf(w, "%v", uni)
	*/

	// Pretty printing
	output, err := json.MarshalIndent(uni, "", "  ")
	if err != nil {
		log.Println("Error during pretty printing of output: " + err.Error())
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%v", string(output))
	fmt.Fprintf(w, "\n\n")

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)
}
