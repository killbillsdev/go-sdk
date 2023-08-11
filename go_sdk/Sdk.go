package go_sdk
import (
	"net/http"
)

func SendBankingTransaction(env string, transactionData []byte, hmacKey string) (*http.Response, bool) {
	return SendDataWithHMAC(env, "transaction", transactionData, hmacKey, ValidateTransactionPayload)
}
