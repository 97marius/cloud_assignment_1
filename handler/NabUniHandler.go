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
func HandleGetRequestNeighboUni(w http.ResponseWriter, r *http.Request) {

	URLparts := strings.Split(r.URL.Path, "/")
	name := URLparts[4]

	url := ""
	if name == "" {
		// URL for all countries if no country has been named in the address bar
		url = "https://restcountries.com/v3.1/all/"
	} else {
		// URL for the named country
		url = "https://restcountries.com/v3.1/name/" + name
	}

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

	nabuni := []NABUNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&nabuni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	borderCountries := strings.Join(nabuni[0].Borders, ",")

	url1 := "https://restcountries.com/v3.1/alpha?codes=" + borderCountries

	NewRequest, err = http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	NewRequest.Header.Add("content-type", "application/json")

	client = &http.Client{}
	defer client.CloseIdleConnections()

	res, err = client.Do(NewRequest)
	//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// Instantiate decoder
	decoder = json.NewDecoder(res.Body)
	// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
	// Note: This does not lead to a check whether an actually provided field is empty!

	// Prepare empty struct to populate
	nabuni = []NABUNI{}

	// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
	err = decoder.Decode(&nabuni)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(nabuni) > 0 {
		for i := range nabuni {

			countryName := nabuni[i].Name.Common

			url2 := "http://universities.hipolabs.com/search?country=" + countryName

			if URLparts[5] != "" {
				url2 += "&name=" + URLparts[5]
			}

			NewRequest, err = http.NewRequest(http.MethodGet, url2, nil)
			if err != nil {
				fmt.Errorf("Error in creating request:", err.Error())
			}

			// Setting content type -> effect depends on the service provider
			NewRequest.Header.Add("content-type", "application/json")

			client = &http.Client{}
			defer client.CloseIdleConnections()

			res, err = client.Do(NewRequest)
			//res, err := client.Get(url) // Alternative: Direct issuing of requests, but fewer configuration options
			if err != nil {
				fmt.Errorf("Error in response:", err.Error())
			}

			// Instantiate decoder
			decoder = json.NewDecoder(res.Body)
			// Ensure parser fails on unknown fields (baseline way of detecting different structs than expected ones)
			// Note: This does not lead to a check whether an actually provided field is empty!

			uni := []UNI{}

			// Decode uni instance --> Alternative: "err := json.NewDecoder(r.Body).Decode(&uni)"
			err = decoder.Decode(&uni)
			if err != nil {
				// Note: more often than not is this error due to client-side input, rather than server-side issues
				http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Flat printing
			//fmt.Println(uni)
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
		}
	} else {
		fmt.Fprintf(w, "Can't find borders for "+name)
		fmt.Println("Can't find borders for " + name)
	}

	//fmt.Println(nabuni)
	//fmt.Fprintf(w, "%v", nabuni)

	// Return status code (good practice)
	http.Error(w, "OK", http.StatusOK)
}
