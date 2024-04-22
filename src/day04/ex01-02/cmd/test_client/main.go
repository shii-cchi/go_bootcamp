package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type requestParams struct {
	CandyType  string `json:"candyType,omitempty"`
	CandyCount int    `json:"candyCount,omitempty"`
	Money      int    `json:"money,omitempty"`
}

type responceResult struct {
	Change  int64  `json:"change"`
	Thanks  string `json:"thanks"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func main() {
	params := handleFlags()

	client, err := makeClient()

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := makeRequest(*client, params)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	getResult(res)
}

func handleFlags() requestParams {
	var params requestParams

	flag.StringVar(&params.CandyType, "k", "", "Define candy type")
	flag.IntVar(&params.CandyCount, "c", 0, "Define candy count")
	flag.IntVar(&params.Money, "m", 0, "Define money received")

	flag.Parse()

	return params
}

func makeClient() (*http.Client, error) {
	rootCA, err := os.ReadFile("ca/minica.pem")

	if err != nil {
		return nil, fmt.Errorf("Error opening root certificate ca: %s\n", err)
	}

	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM(rootCA)

	cert, err := tls.LoadX509KeyPair("ca/client.candy.tld/cert.pem", "ca/client.candy.tld/key.pem")

	if err != nil {
		return nil, fmt.Errorf("Error loading client certificate and key: %s\n", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      cp,
	}

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}, nil
}

func makeRequest(client http.Client, params requestParams) (*http.Response, error) {
	requestBody, err := json.Marshal(params)

	if err != nil {
		return nil, fmt.Errorf("Error marshaling request: %s\n", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://127.0.0.1:3333/buy_candy", bytes.NewReader(requestBody))

	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s\n", err)
	}

	return res, nil
}

func getResult(res *http.Response) {
	var result responceResult

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Printf("Error parsing the response body: %s\n", err)
		return
	}

	if res.StatusCode == 201 {
		fmt.Printf("%s Your change is %d\n", result.Thanks, result.Change)
		return
	}

	if res.StatusCode == 400 || res.StatusCode == 402 {
		if result.Error == "" {
			fmt.Printf("Error: %s\n", result.Message)
			return
		}

		fmt.Printf("Error: %s\n", result.Error)
		return
	}

	fmt.Printf("Error: %s\n", result.Message)
	return
}
