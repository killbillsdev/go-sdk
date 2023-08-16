package go_sdk

import (
	"encoding/json"
	"testing"
	"time"

)

func TestValidateTransactionPayload_Valid(t *testing.T) {
	constantTime := time.Date(2023, time.August, 14, 12, 0, 0, 0, time.UTC)

	validTransaction := struct {
		Bank_id     string `json:"bank_id"`
		Transaction struct {
			Siret             string    `json:"siret" validate:"number"`
			Amount            int       `json:"amount" validate:"required,gt=0"`
			Payment           Payment   `json:"payment" validate:"required"`
			Currency          string    `json:"currency" validate:"required,oneof=EUR USD"`
			StoreName         string    `json:"store_name" validate:"required"`
			CustomerID        string    `json:"customer_id" validate:"required"`
			MerchantID        string    `json:"merchant_id" validate:"required"`
			ReferenceID       string    `json:"reference_id" validate:"required"`
			MerchantName      string    `json:"merchant_name" validate:"required"`
			TransactionDate   time.Time `json:"transaction_date" validate:"required"`
			BillingDescriptor string    `json:"billing_descriptor" validate:"required"`
		} `json:"transaction"`
		CallbackURL   string `json:"callback_url"`
		PartnerName   string `json:"partner_name"`
		ReceiptFormat string `json:"receipt_format"`
	}{
		Bank_id: "fbec0cb5-91c8-4b8b-a194-c018fbfe258d",
		Transaction: struct {
			Siret             string    `json:"siret" validate:"number"`
			Amount            int       `json:"amount" validate:"required,gt=0"`
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
			Siret:  "123456789",
			Amount: 100,
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

	jsonData, _ := json.Marshal(validTransaction)

	if !ValidateTransactionPayload(jsonData) {
		t.Errorf("Expected valid transaction payload, but got invalid")
	}
}

func TestValidateTransactionPayload_Invalid(t *testing.T) {
	LoadEnv()
	invalidTransaction := Transaction{
		// ... Populate fields with invalid data ...
	}

	jsonData, _ := json.Marshal(invalidTransaction)

	if ValidateTransactionPayload(jsonData) {
		t.Errorf("Expected invalid transaction payload, but got valid")
	}
}

func TestValidateReceiptPayload_Valid(t *testing.T) {
	

	var validPayload2 = map[string]interface{}{
		"date": "2023-07-30T09:04:08",
		"mode": "0",
		"items": []map[string]interface{}{
			{
				"tax": map[string]interface{}{
					"rate":        1000,
					"amount":      85,
					"description": "TVA",
				},
				"name":     "Salade ATCHOUM",
				"price":    850,
				"quantity": 1,
				"subitems": []map[string]interface{}{
					{
						"tax": map[string]interface{}{
							"rate":        1000,
							"amount":      30,
							"description": "TVA",
						},
						"name":         "Atchoum V1",
						"price":        1555,
						"quantity":     1,
						"description":  "",
						"reference_id": "5df1e0fa-3bdc-461a-9170-a74bb2f0792b",
						"total_amount": 300,
					},
					{
						"tax": map[string]interface{}{
							"rate":        1000,
							"amount":      40,
							"description": "TVA",
						},
						"name":         "Saucisses v1",
						"quantity":     1,
						"description":  "",
						"reference_id": "d15e20c6-925c-491a-8381-153c9352aadd",
						"total_amount": 400,
					},
					{
						"tax": map[string]interface{}{
							"rate":        1000,
							"amount":      25,
							"description": "TVA",
						},
						"name":         "Thé v1",
						"quantity":     1,
						"description":  "",
						"reference_id": "72b2479f-9210-44fc-8187-a4f40bc31ee6",
						"total_amount": 250,
					},
				},
				"description":  "",
				"reference_id": "1c49ad5c-2610-4bd7-bbb5-e235639a0a42",
				"total_amount": 850,
			},
			// ... (other items)
		},
		"store": map[string]interface{}{
			"store_name":         "RESTAU TEST",
			"siret":              "66666666666666",
			"billing_descriptor": "RESTAU TEST",
			"address": map[string]interface{}{
				"city":           "Paris",
				"number":         0,
				"country":        "FRANCE",
				"postal_code":    75014,
				"street_address": "17 rue du Smart Receipt",
			},
			"code_ape":     "4410",
			"tva_intra":    "FR 000 000 00",
			"reference_id": "1",
		},
		"table": "31",
		"taxes": []map[string]interface{}{
			{
				"rate":        1000,
				"aamount":     85,
				"description": "TVA",
			},
			{
				"rate":        2000,
				"amount":      190,
				"description": "TVA",
			},
		},
		"amount":   871741,
		"covers":   2,
		"invoice":  1,
		"currency": "EUR",
		"merchant": map[string]interface{}{
			"merchant_name": "Restaurant test",
			"reference_id":  "1234",
		},
		"payments": []map[string]interface{}{
			{
				"bin":              "0",
				"amount":           871741,
				"scheme":           "",
				"auth_code":        "",
				"last_four":        "0",
				"payment_type":     "CB",
				"transaction_id":   "null",
				"transaction_date": "2023-07-30T09:04:08",
			},
		},
		"partner_name": "any",
		"reference_id": "1221554511",
	}

	

	jsonData, _ := json.Marshal(validPayload2)

	if !ValidateReceiptPayload(jsonData) {
		t.Errorf("Expected valid transaction payload, but got invalid")
	}
}
