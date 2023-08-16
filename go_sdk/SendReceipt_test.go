package go_sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func testSendReceiptValidPayload(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier .env:", err)
		os.Exit(1)
	}

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
		"partner_name": os.Getenv("TEST_POS_PARTNER_NAME"),
		"reference_id": "1221554511",
	}
	jsonPayload, err := json.Marshal(validPayload2)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
		return
	}

	fmt.Println(os.Getenv("TEST_POS_HMAC"))

	result, status := SendReceipt("dev", jsonPayload, os.Getenv("TEST_POS_HMAC"))
	//print result
	fmt.Println(status)
	fmt.Println(result)

	expectedStatus := true
	if status != expectedStatus {
		t.Errorf("Expected status: %v", expectedStatus)
	}
}
