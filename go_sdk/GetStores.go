package go_sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStores(env string, apiKey string, offset int, limit int) ([]interface{}, error) {
	if env == "" || apiKey == "" {
		return nil, fmt.Errorf("No environment specified or API key provided")
	}

	var baseURL string
	if env == "prod" {
		baseURL = "https://w.killbills.co"
	} else {
		baseURL = fmt.Sprintf("https://w.%s.killbills.dev", env)
	}

	url := fmt.Sprintf("%s/stores?offset=%d&limit=%d", baseURL, offset, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", apiKey)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	items, ok := data["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Response does not contain 'items'")
	}

	return items, nil
}
