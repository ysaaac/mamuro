package zincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mamuro-backend/config"
	"net/http"
)

// SendJSONPost function performs the actual POST request and returns the error or response
func SendJSONPost(endpoint string, data interface{}) (*http.Response, error) {
	zincsearchUri := config.GetEnv("ZINCSEARCH_URI", "localhost:4080")

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, zincsearchUri+endpoint, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

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
