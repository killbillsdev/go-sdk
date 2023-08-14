package go_sdk
import (
	"fmt"
	"testing"
	"time"
	"encoding/json"
)

func testSendBankingTransaction(t *testing.T){
	constantTime := time.Date(2023, time.August, 14, 12, 0, 0, 0, time.UTC)
	validTransaction := Transaction{
		Bank_id: "fbec0cb5-91c8-4b8b-a194-c018fbfe258d",
		Transaction: struct {
			Siret             string    `json:"siret" validate:"number"`
			Amount            float64   `json:"amount" validate:"required,gt=0"`
			Payment           Payment   `json:"payment" validate:"required"`
			Currency          string    `json:"currency" validate:"required,oneof=EUR USD"`
			StoreName         string    `json:"store_name" validate:"required"`
			CustomerID        string    `json:"customer_id" validate:"required"`
			MerchantID        string    `json:"merchant_id" validate:"required"`
			ReferenceID       string    `json:"reference_id" validate:"required"`
			MerchantName      string    `json:"merchant_name" validate:"required"`
			TransactionDate   time.Time `json:"transaction_date" validate:"required"`
			BillingDescriptor string    `json:"billing_descriptor" validate:"required"`
		}{
			Siret: "123456789",
			Amount: 100.0,
			Payment: Payment{
				Bin: "123",
			},
			Currency:          "EUR",
			StoreName:         "Store A",
			CustomerID:        "cust123",
			MerchantID:        "merchant123",
			ReferenceID:       "ref123",
			MerchantName:      "Merchant A",
			TransactionDate:   constantTime,
			BillingDescriptor: "Billing Desc",
		},
		CallbackURL:   "https://eolpqff4hbbj6q5.m.pipedream.net",
		PartnerName:   "mooncard",
		ReceiptFormat: "PDF",
	}
	transactionData, err := json.Marshal(validTransaction)
	if err != nil {
		t.Fatalf("Failed to marshal transaction data: %v", err)
	}
	fmt.Println(string(transactionData))
	fmt.Println(SendBankingTransaction("test", transactionData,  "DE$G3@iPFq2"))
}