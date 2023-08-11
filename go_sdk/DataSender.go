package go_sdk

import (
	"bytes"
	"fmt"
	"net/http"
)
type ValidatorFunc func(data []byte) bool
func SendDataWithHMAC(env string, endpoint string, data []byte, hmacSignature string,  validator ValidatorFunc) (*http.Response, bool) {
	if data == nil || hmacSignature == "" {
		return nil, false
	}

	payloadValidationResult := validator(data)
	if !payloadValidationResult {
		return nil, payloadValidationResult
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://in.%skillbills.%s/%s", env, "dev", endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, false
	}

	hashedPayload := CipherHmacPayload(string(data), hmacSignature)
	req.Header.Set("Authorization", fmt.Sprintf("hmac %s", hashedPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, false
	}

	return resp, true
}


