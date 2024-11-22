package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostToAPI(url string, data interface{}) error {
	log.Printf("POST request to %s with data: %v", url, data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("API returned status: %v, response: %s", resp.Status, bodyString)
	}

	return nil
}
