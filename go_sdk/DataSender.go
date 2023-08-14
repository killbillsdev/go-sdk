package go_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ValidatorFunc func(data []byte) bool

func SendDataWithHMAC(env string, endpoint string, data1 []byte, hmacSignature string, validator ValidatorFunc) (string, bool) {
	if data1 == nil || hmacSignature == "" {
		return "", false
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data1, &jsonData); err != nil {
		fmt.Println("Error:", err)
		return "", false
	}

	payloadValidationResult := validator(data1)
	if !payloadValidationResult {
		return "", false
	}

	client := &http.Client{}
	var baseURL string
	if env == "prod" {
		baseURL = "https://in.killbills.co"
	} else {
		baseURL = fmt.Sprintf("https://in.%s.killbills.dev", env)
	}
	url := fmt.Sprintf("%s/%s", baseURL, endpoint)

	// Use the original data1 as payload directly
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data1))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", false
	}

	hashedPayload := CipherHmacPayload(string(data1), hmacSignature)
	req.Header.Set("Authorization", fmt.Sprintf("hmac %s", hashedPayload))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("hashedPayload", hashedPayload)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	} else {
		fmt.Println("Response Body:", string(body))
	}
	return string(body), true
}
