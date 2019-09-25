package main

import (
	"encoding/json"
	"fmt"
	"github.com/ejcapetillo/nasa-rover-go/models"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	parameters := models.Parameters{EarthDate: "2015-06-03", APIKey: "DEMO_KEY"}
	values, _ := query.Values(parameters)
	url := "https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?" + values.Encode()
	resp, err := http.Get(url)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", responseBody)

	//fmt.Println(http.ListenAndServe(":8080", nil))

	photoList := models.PhotoWrapper{}
	//var photoList []interface{}
	//var photos []models.Photo
	//var photoList []models.PhotoWrapper
	jsonError := json.Unmarshal(responseBody, &photoList)
	if jsonError != nil {
		_, _ = fmt.Fprintf(os.Stderr, "JSON ERROR: %v\n", jsonError)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(photoList)
}
