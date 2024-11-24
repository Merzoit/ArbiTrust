package api

import (
	"arbbot/constants"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostToAPI(url string, data interface{}) error {
	log.Printf("API: "+constants.CallPOST, url, data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("API: "+constants.ErrEncodingData, err)
		return fmt.Errorf(constants.ErrEncodingData, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("API: "+constants.ErrSendingRequest, err)
		return fmt.Errorf(constants.ErrSendingRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("API: "+constants.LogReturnedStatus, resp.Status, bodyString)
		return fmt.Errorf(constants.LogReturnedStatus, resp.Status, bodyString)
	}

	log.Println(constants.LogPOSTCompleted)
	return nil
}
