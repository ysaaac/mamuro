package zincsearch

import (
	"api/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request function performs request and handle permissions
func Request(method string, endpoint string, data interface{}) (*http.Response, error) {
	zincsearchUri := config.GetEnv("ZINCSEARCH_URI", "http://localhost:4080")

	var jsonData []byte
	if data != nil {
		var err error
		jsonData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, zincsearchUri+endpoint, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(
		config.GetEnv("ZINC_FIRST_USER", "admin"),
		config.GetEnv("ZINC_FIRST_PASSWORD", "admin123"),
	)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// HandleResponse function checks the status code and unmarshals the response if successful
func HandleResponse(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
