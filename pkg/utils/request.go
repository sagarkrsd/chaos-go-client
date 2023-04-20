package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// SendRequest sends a HTTP request and returns the response body
func SendRequest(url, method string, payload map[string]interface{}) ([]byte, error) {
	client := &http.Client{}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	// API key should be exported as an ENV: X_API_KEY
	req.Header.Add("x-api-key", os.Getenv("X_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error while registering a new Chaos infra: %+v", err)
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error while reading response for registering a new Chaos infra: %+v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("\nNon-OK HTTP status: %+v, responseBody: %+v\n", res.StatusCode, string(resBody))
		// You may read / inspect response body
		return resBody, fmt.Errorf("Non-OK HTTP status: %+v", res.StatusCode)
	}

	return resBody, nil
}
