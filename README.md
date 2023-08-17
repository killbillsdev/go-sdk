
# KillBills - SDK (Go)

The main purpose of this Software Development Kit (SDK) is to facilitate the integration of KillBills' services into GO applications. The SDK will act as an intermediary layer, allowing developers to communicate with the KillBills API and access its features in a straightforward manner. By providing a consistent and easy-to-use interface, the SDK aims to simplify & speed up the process of integrating KillBills' services into Go applications.



# Dependencies
![GO](https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg)

## Features

- getStores : The method getStores returns a list of killbills integrated stores as an array of objects with relevant properties.
- sendTransaction : The method sendTransaction validates and send transaction to the killbills gate transaction.
- sendReceipt : The method sendReceipt validates and send a receipt to the killbills gate receipt.




## Usage/Examples
#### Installation:
```shell
go get github.com/killbillsdev/go_sdk/go_sdk   
```
#### Import:
```go
import (
  sdk "github.com/killbillsdev/go_sdk/go_sdk"
)
```
#### Method getStores :
```go
sdk.GetStores("dev", "your-api-key");
```
##### Output:
```yaml
[
  {
    "id": "id",
    "billing_descriptor": "store billing descriptor",
    "store_name": "store name",
    "merchant_name": "merchant name",
    "full_address": "00 street, 00000, city",
    "siret": 00000000000000
  },
   {
    "id": "id",
    "billing_descriptor": "store billing descriptor",
    "store_name": "store name",
    "merchant_name": "merchant name",
    "full_address": "00 street, 00000, city",
    "siret": 00000000000000
  },
  //...other results
]
```
# 
#### Method sendTransaction :
##### note that the transactionData object only contains minimal required values see (insert link to all possibilities)
```go
constantTime := time.Date(2023, time.August, 14, 12, 0, 0, 0, time.UTC)
	validTransaction := Transaction{
		Bank_id: "your-org-id",
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
		CallbackURL:   "callback",
		PartnerName:   "partnername",
		ReceiptFormat: "PDF",
	}
	transactionData, err := json.Marshal(validTransaction)
	if err != nil {
		t.Fatalf("Failed to marshal transaction data: %v", err)
	}
	sdk.SendBankingTransaction("test", transactionData,  "your-api-key")
```
##### Output: 
```yaml
{
"status": "success",
"message": "published to gate transaction",
"messageId": "xxxxxxxxxxxxxxxxx",
"previewLink": "https://banks.killbills.co/payloads/xxxxxxxxxxxxxxxxx"
}
```
# 
#### Method sendReceipt :
##### note that the receiptData object only contains minimal required values see (insert link to all possibilities)
```go
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
		"partner_name": "TEST_POS_PARTNER_NAME",
		"reference_id": "1221554511",
	}
	jsonPayload, err := json.Marshal(validPayload2)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
		return
	}

	result, status := sdk.SendReceipt("dev", jsonPayload, "TEST_POS_HMAC")
```
##### Output: 
```yaml
{
"status": "success",
"message": "published to gate receipt",
"messageId": "xxxxxxxxxxxxxxxxx",
"previewLink": "https://merchants.killbills.co/payloads/xxxxxxxxxxxxxxxxx"
}
```

## Transaction & Receipt Architecture

#### TRANSACTION :
```typescript
transactionData = {
    bank_id: '', // string (36 caractères)
    callback_url: '', // string
    partner_name: '', // string
    kb_features: [], // tableau de chaînes de caractères ou nombres vides
    receipt_format: '', // string ('JSON', 'PDF', 'SVG', 'PNG')
    transaction: {
        reference_id: '', // string
        amount: '', // number (positif)
        customer_id: '', // string
        transaction_date: '', // date (au format chaîne de caractères)
        store_name: '', // string ou vide
        billing_descriptor: '', // string
        siret: '', // string ou vide
        payment: '', // objet vide
        currency: '', // string ou vide
        pos_name: '', // string ou vide
        merchant_name: '' // string ou vide
    }
};
```
#### RECEIPT :
```typescript
receiptData = {
    reference_id: '', // string (alphanumérique)
    amount: 0, // number
    total_tax_amount: '', // number ou vide
    currency: '', // string ('EUR' ou 'USD')
    date: '', // string (au format 'YYYY-MM-DDTHH:mm:ss')
    covers: 0, // number ou vide
    table: '', // string ou vide
    invoice: 0, // number ou vide
    total_discount: 0, // number ou vide
    mode: 0, // number ou vide
    partner_name: '', // string
    merchant: {
        merchant_name: '', // string ou vide
        reference_id: '', // string
        merchant_id: 0 // number ou vide
    },
    store: {
        store_name: '', // string
        reference_id: '', // string
        billing_descriptor: '', // string
        siret: '', // string (14 caractères)
        code_ape: '', // string ou vide
        tva_intra: '', // string ou vide
        address: {
            postal_code: 0, // number
            street_address: '', // string ou vide
            country: '', // string ou vide
            city: '', // string ou vide
            full_address: '', // string ou vide
            number: 0 // number ou vide
        }
    },
    taxes: [{ 
        description: '', // string ou vide
        amount: 0, // number
        rate: 550 // number (550, 1000 ou 2000) ou vide
    }],
    items: [{
        reference_id: '', // string ou vide
        name: '', // string
        description: '', // string ou vide
        type: '', // string ou vide
        quantity: 0, // number
        price: 0, // number
        discount: 0, // number ou vide
        total_amount: 0, // number ou vide
        tax: {
            description: '', // string ou vide
            amount: 0, // number
            rate: 550 // number (550, 1000 ou 2000) ou vide
        },
        subitems: [{
            reference_id: '', // string ou vide
            name: '', // string
            description: '', // string ou vide
            quantity: 0, // number ou vide
            price: 0, // number ou vide
            discount: 0, // number ou vide
            total_amount: 0, // number ou vide
            tax: {
                description: '', // string ou vide
                amount: 0, // number
                rate: 550 // number (550, 1000 ou 2000) ou vide
            }
        }]
    }],
    payments: [{
        bin: '', // string ou vide
        last_four: '', // string ou vide
        auth_code: '', // string ou vide
        scheme: '', // string ou vide
        amount: 0, // number
        transaction_date: '', // string (au format 'YYYY-MM-DDTHH:mm:ss')
        transaction_id: '', // string ou vide
        payment_type: '' // string ou vide
    }]
};
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/killbillsdev/go-sdk
```

Go to the project directory

```bash
  cd go-sdk
```

Install dependencies
```bash
make
```



## Running Tests

To run tests, run the following command

```bash
make test / mvn run test
```


## Feedback / Feature request

If you have any feedback, please reach out to us at contact@killbills.co
## License

[MIT](https://choosealicense.com/licenses/mit/) [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


