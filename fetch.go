package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseUrl string = "api.coincap.io"
var version string = "v2"

type rate struct {
	Data struct {
		Id             string `json:"id"`
		Symbol         string `json:"symbol"`
		Currencysymbol string `json:"currencySymbol"`
		Type           string `json:"type"`
		Rateusd        string `json:"rateUsd"`
	} `json:"data"`
	Timestamp uint64 `json:"timestamp"`
}

func fetch(resource string, id string) rate {
	url := fmt.Sprintf("https://%v/%v/%v/%v", baseUrl, version, resource, id)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, getErr := client.Do(req)
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
	rate := rate{}
	if err := json.Unmarshal(body, &rate); err != nil {
		log.Fatal(err)
	}

	return rate
}
