package nba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// FetchNBADataJSON function to request NBA stats API with an endpoint
// reassigns response to target of type interface{}
func FetchNBADataJSON(endpoint string, target interface{}) {
	client := &http.Client{Timeout: 10 * time.Second}
	root := "http://data.nba.net/10s"
	url := root + endpoint

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(target)
}
