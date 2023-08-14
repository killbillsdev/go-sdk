package go_sdk

import (
	"time"
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}


type Payment struct {
	BIN            string  `json:"bin"`
	LastFour       string  `json:"last_four"`
	AuthCode       string  `json:"auth_code"`
	Scheme         string  `json:"scheme"`
	Amount         float64 `json:"amount" validate:"required"`
	TransactionDate string  `json:"transaction_date" validate:"regexp=^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$"`
	TransactionID   string  `json:"transaction_id"`
	PaymentType    string  `json:"payment_type"`
}

type Transaction struct {
	Bank_id   string `json:"bank_id" validate:"required"`
	Transaction struct {
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
	} `json:"transaction" validate:"required"`
	CallbackURL    string `json:"callback_url" validate:"required"`
	PartnerName    string `json:"partner_name" validate:"required"`
	ReceiptFormat  string `json:"receipt_format" validate:"required"`
}

type ReceiptPayload struct {
	ReferenceID      string     `json:"reference_id" validate:"required"`
	Amount           float64    `json:"amount" validate:"required"`
	TotalTaxAmount   float64    `json:"total_tax_amount"`
	Currency         string     `json:"currency" validate:"required,oneof=EUR USD"`
	Date             string     `json:"date" validate:"required,regexp=^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}$"`
	Covers           int        `json:"covers"`
	Table            string     `json:"table"`
	Invoice          int        `json:"invoice"`
	TotalDiscount    float64    `json:"total_discount"`
	Mode             string     `json:"mode"`
	PartnerName      string     `json:"partner_name" validate:"required"`
	Merchant         Merchant   `json:"merchant" validate:"required"`
	Store            Store      `json:"store" validate:"required"`
	Taxes            []Tax      `json:"taxes"`
	Items            []Item     `json:"items" validate:"required"`
	Payments         []Payment  `json:"payments" validate:"required"`
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

type SubItem struct {
	ReferenceID   string  `json:"reference_id"`
	Name          string  `json:"name" validate:"required"`
	Description   string  `json:"description"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Discount      float64 `json:"discount"`
	TotalAmount   float64 `json:"total_amount"`
	Tax           Tax     `json:"tax"`
}

type Item struct {
	ReferenceID   string    `json:"reference_id"`
	Name          string    `json:"name" validate:"required"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	Quantity      int       `json:"quantity" validate:"required"`
	Price         float64   `json:"price" validate:"required"`
	Discount      float64   `json:"discount"`
	TotalAmount   float64   `json:"total_amount"`
	Tax           Tax       `json:"tax" validate:"required"`
	SubItems      []SubItem `json:"subitems"`
}

type Merchant struct {
	MerchantName  string `json:"merchant_name"`
	ReferenceID   string `json:"reference_id" validate:"required"`
	MerchantID    int    `json:"merchant_id"`
}

type Store struct {
	StoreName        string  `json:"store_name" validate:"required"`
	ReferenceID      string  `json:"reference_id" validate:"required"`
	BillingDescriptor string  `json:"billing_descriptor" validate:"required"`
	Siret            string  `json:"siret" validate:"required,min=14,max=14"`
	CodeAPE          string  `json:"code_ape"`
	TVAIntra         string  `json:"tva_intra"`
	Address          Address `json:"address" validate:"required"`
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