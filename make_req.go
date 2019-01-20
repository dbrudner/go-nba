package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getNBAJSON(endpoint string, target interface{}) {
	var client = &http.Client{Timeout: 10 * time.Second}
	root := "http://data.nba.net/10s/"
	url := root + endpoint

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)
}
