package main

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"os"
)

type Parameters struct {
	EarthDate string `url:"earth_date"`
	APIKey string `url:"api_key"`
}

func main() {
	parameters := Parameters{"2015-06-03", "DEMO_KEY"}
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

	fmt.Printf("%s", responseBody)

	//fmt.Println(http.ListenAndServe(":8080", nil))
}

type nasa struct {
	Id uint64 `json:"id"`
}
