package go_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStores(env string, apiKey string) ([]map[string]interface{}, error) {
	if env == "" {
		return nil, fmt.Errorf("No environment specified")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("No API key provided")
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://w.%skillbills.%s/stores", env, "co")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	items, ok := result["items"].([]map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Failed to parse response")
	}

	return items, nil
}