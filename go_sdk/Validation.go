package go_sdk

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Payment struct {
	Bin           string `json:"bin"`
	Scheme        string `json:"scheme"`
	LastFour      string `json:"lastFour"`
	AuthCode      string `json:"auth_code"`
	TransactionID string `json:"transaction_id"`
}

type Transaction struct {
	Bank_id     string `json:"bank_id" validate:"required"`
	Transaction struct {
		Siret             string    `json:"siret" validate:"number"`
		Amount            int       `json:"amount" validate:"required"`
		Payment           Payment   `json:"payment" validate:"required"`
		Currency          string    `json:"currency" validate:"required,oneof=EUR USD"`
		StoreName         string    `json:"store_name" validate:"required"`
		CustomerID        string    `json:"customer_id" validate:"required"`
		MerchantID        string    `json:"merchant_id" validate:"required"`
		ReferenceID       string    `json:"reference_id" validate:"required"`
		MerchantName      string    `json:"merchant_name" validate:"required"`
		TransactionDate   time.Time `json:"transaction_date" validate:"required"`
		BillingDescriptor string    `json:"billing_descriptor" validate:"required"`
	} `json:"transaction" validate:"required"`
	CallbackURL   string `json:"callback_url" validate:"required"`
	PartnerName   string `json:"partner_name" validate:"required"`
	ReceiptFormat string `json:"receipt_format" validate:"required"`
}
type ReceiptPayload struct {
	ReferenceID    string    `json:"reference_id" validate:"required"`
	Amount         float64   `json:"amount" validate:"required"`
	TotalTaxAmount float64   `json:"total_tax_amount"`
	Currency       string    `json:"currency" validate:"required,oneof=EUR USD"`
	Date           string    `json:"date" validate:"required"`
	Covers         int       `json:"covers"`
	Table          string    `json:"table"`
	Invoice        int       `json:"invoice"`
	TotalDiscount  float64   `json:"total_discount"`
	Mode           string    `json:"mode"`
	PartnerName    string    `json:"partner_name" validate:"required"`
	Merchant       Merchant  `json:"merchant" validate:"required"`
	Store          Store     `json:"store" validate:"required"`
	Taxes          []Tax     `json:"taxes"`
	Items          []Item    `json:"items" validate:"required"`
	Payments       []Payment `json:"payments" validate:"required"`
}

type Merchant struct {
	MerchantName string `json:"merchant_name"`
	ReferenceID  string `json:"reference_id" validate:"required"`
	MerchantID   int    `json:"merchant_id"`
}

type Store struct {
	StoreName         string  `json:"store_name" validate:"required"`
	ReferenceID       string  `json:"reference_id" validate:"required"`
	BillingDescriptor string  `json:"billing_descriptor" validate:"required"`
	Siret             string  `json:"siret" validate:"required,min=14,max=14"`
	CodeAPE           string  `json:"code_ape"`
	TVAIntra          string  `json:"tva_intra"`
	Address           Address `json:"address" validate:"required"`
}

type Address struct {
	PostalCode    int    `json:"postal_code" validate:"required"`
	StreetAddress string `json:"street_address"`
	Country       string `json:"country"`
	City          string `json:"city"`
	FullAddress   string `json:"full_address"`
	Number        int    `json:"number"`
}

type Tax struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount" validate:"required"`
	Rate        int     `json:"rate" validate:"oneof=550 1000 2000"`
}

type Item struct {
	ReferenceID string    `json:"reference_id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Quantity    int       `json:"quantity" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Discount    float64   `json:"discount"`
	TotalAmount float64   `json:"total_amount"`
	Tax         Tax       `json:"tax" validate:"required"`
	SubItems    []SubItem `json:"subitems"`
}

type SubItem struct {
	ReferenceID string  `json:"reference_id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	TotalAmount float64 `json:"total_amount"`
	Tax         Tax     `json:"tax"`
}

func ValidateTransactionPayload(data []byte) bool {
	var transaction Transaction
	if err := json.Unmarshal(data, &transaction); err != nil {
		return false
	}
	err := validate.Struct(transaction)
	if err != nil {
		return false
	}
	return true
}

func ValidateReceiptPayload(data []byte) bool {
	var receipt ReceiptPayload
	fmt.Println("qsfpjjkfsjklsfjklsfd")
	if err := json.Unmarshal(data, &receipt); err != nil {
		fmt.Println(err)
		return false
	}

	err := validate.Struct(receipt)
	fmt.Println(err)
	if err != nil {
		//print error
		fmt.Println(err)
		return false
	}
	return true
}
