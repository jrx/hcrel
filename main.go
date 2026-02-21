package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

	url := "https://api.releases.hashicorp.com/v1/releases/" + *product + "/latest?license_class=" + *license

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rel release
	if err := json.Unmarshal(body, &rel); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rel.Version)
}
