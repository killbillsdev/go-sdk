package go_sdk
import (
	"encoding/json"
)

func SendBankingTransaction(env string, transactionData []byte, hmacKey string) (string, bool) {
	compactTransactionData, err := json.Marshal(json.RawMessage(transactionData))
    if err != nil {
        return "", false
    }
	return SendDataWithHMAC(env, "transaction", compactTransactionData, hmacKey, ValidateTransactionPayload)
}