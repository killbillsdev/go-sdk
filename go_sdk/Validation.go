package go_sdk

import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Payment struct {
	Bin           string `validate:""`
	LastFour      string `validate:""`
	AuthCode      string `validate:""`
	Scheme        string `validate:""`
	TransactionID string `validate:""`
}

type Transaction struct {
	ReferenceID       string     `validate:"required"`
	Amount            float64    `validate:"required,gt=0"`
	CustomerID        string     `validate:"required"`
	TransactionDate   time.Time  `validate:"required"`
	StoreName         string     `validate:"required"`
	BillingDescriptor string     `validate:"required"`
	Siret             string     `validate:"omitempty,number"`
	Payment           Payment    `validate:"required"`
	Currency          string     `validate:"required,oneof=EUR USD"`
	MerchantName      string     `validate:""`
}

type ReceiptPayload struct {
	ReferenceID    string     `validate:"required"`
	Amount         float64    `validate:"required"`
	TotalTaxAmount float64    `validate:""`
	Currency       string     `validate:"required,oneof=EUR USD"`
	Date           string     `validate:"required,datetime"`
	Covers         int        `validate:""`
	Table          string     `validate:""`
	Invoice        int        `validate:""`
	TotalDiscount  float64    `validate:""`
	Mode           string     `validate:""`
	PartnerName    string     `validate:"required"`
	Merchant       Merchant   `validate:"required"`
	Store          Store      `validate:"required"`
	Taxes          []Tax      `validate:""`
	Items          []Item     `validate:"required"`
	Payments       []Payment  `validate:"required"`
}

type Merchant struct {
	MerchantName string `validate:""`
	ReferenceID  string `validate:"required"`
	MerchantID   int    `validate:""`
}

type Store struct {
	StoreName         string   `validate:"required"`
	ReferenceID       string   `validate:"required"`
	BillingDescriptor string   `validate:"required"`
	Siret             string   `validate:"required"`
	CodeApe           string   `validate:""`
	TvaIntra          string   `validate:""`
	Address           Address  `validate:"required"`
}

type Address struct {
	PostalCode    int    `validate:"required"`
	StreetAddress string `validate:""`
	Country       string `validate:""`
	City          string `validate:""`
	FullAddress   string `validate:""`
	Number        int    `validate:""`
}

type Tax struct {
	Description string  `validate:""`
	Amount      float64 `validate:"required"`
	Rate        float64 `validate:"oneof=550 1000 2000"`
}

type Item struct {
	ReferenceID  string `validate:""`
	Name         string `validate:"required"`
	Description  string `validate:""`
	Type         string `validate:""`
	Quantity     int    `validate:"required"`
	Price        float64 `validate:"required"`
	Discount     float64 `validate:""`
	TotalAmount  float64 `validate:""`
	Tax          Tax     `validate:"required"`
	Subitems     []Subitem `validate:""`
}

type Subitem struct {
	ReferenceID  string `validate:""`
	Name         string `validate:"required"`
	Description  string `validate:""`
	Quantity     int    `validate:""`
	Price        float64 `validate:""`
	Discount     float64 `validate:""`
	TotalAmount  float64 `validate:""`
	Tax          Tax     `validate:"required"`
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
	if err := json.Unmarshal(data, &receipt); err != nil {
		return false
	}
	err := validate.Struct(receipt)
	if err != nil {
		return false
	}
	return true
}