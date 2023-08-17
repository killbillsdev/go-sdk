package go_sdk

import (
	"encoding/json"
)

func SendReceipt(env string, receiptData []byte, hmacKey string) (string, bool) {

	compactReceiptData, err := json.Marshal(json.RawMessage(receiptData))
	if err != nil {

		return "", false
	}
	return SendDataWithHMAC(env, "receipt", compactReceiptData, hmacKey, ValidateReceiptPayload)
}
