package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type requestParams struct {
	candyType     string
	candyCount    int
	moneyReceived int
}

func main() {
	params := handleFlags()

	certificate :=

	config := &tls.Config{
		Certificates: certificate,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	requestBody := []byte(fmt.Sprintf(`{"money": %d, "candyType": %s, "candyCount": %d}`, params.moneyReceived, params.candyType, params.candyCount))

	req, err := http.NewRequest(http.MethodPost, "https://candy.tld/buy_candy", bytes.NewReader(requestBody))

	if err != nil {
		fmt.Printf("Error creating request", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error sending request", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 400 || res.StatusCode == 402 {
		var errorResponse map[string]interface{}

		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			fmt.Printf("Error parsing the error response body", err)
			return
		}

		fmt.Printf("Error:", errorResponse["error"])
		return
	}

	if res.StatusCode == 201 {
		var result map[string]interface{}

		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			fmt.Printf("Error parsing the response body", err)
			return
		}

		fmt.Printf("%s Your change is %d", result["thanks"], result["change"])
		return
	}
}

func handleFlags() requestParams {
	var params requestParams
	flag.StringVar(&params.candyType, "k", "", "Define candy type")
	flag.IntVar(&params.candyCount, "c", 0, "Define candy count")
	flag.IntVar(&params.candyCount, "m", 0, "Define money received")

	flag.Parse()

	return params
}
