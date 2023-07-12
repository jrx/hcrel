package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type release struct {
	Version string `json:"version"`
}

func main() {

	product := flag.String("product", "vault", "name of the product: vault, consul, boundary")
	license := flag.String("license", "oss", "class of the license: oss, enterprise, hcp")

	flag.Parse()

	p := *product
	l := *license

	url := "https://api.releases.hashicorp.com/v1/releases/" + p + "/latest?license_class=" + l

	apiClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	release1 := release{}
	jsonErr := json.Unmarshal(body, &release1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(release1.Version)

}
