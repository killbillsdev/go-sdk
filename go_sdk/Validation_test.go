package go_sdk

import (
	"testing"
	"time"
	"encoding/json"
)

func TestValidateTransactionPayload_Valid(t *testing.T) {

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
			TransactionDate:   time.Now(),
			BillingDescriptor: "Billing Desc",
		},
		CallbackURL:   "https://eolpqff4hbbj6q5.m.pipedream.net",
		PartnerName:   "mooncard",
		ReceiptFormat: "PDF",
	}

	jsonData, _ := json.Marshal(validTransaction)

	if !ValidateTransactionPayload(jsonData) {
		t.Errorf("Expected valid transaction payload, but got invalid")
	}
}

func TestValidateTransactionPayload_Invalid(t *testing.T) {
	invalidTransaction := Transaction{
		// ... Populate fields with invalid data ...
	}

	jsonData, _ := json.Marshal(invalidTransaction)

	if ValidateTransactionPayload(jsonData) {
		t.Errorf("Expected invalid transaction payload, but got valid")
	}
}


