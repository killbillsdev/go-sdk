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
	Bin           string `json:"bin"`
	Scheme        string `json:"scheme"`
	LastFour      string `json:"lastFour"`
	AuthCode      string `json:"auth_code"`
	TransactionID string `json:"transaction_id"`
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
	ReferenceID    string     `validate:"required"`
	Amount         float64    `validate:"required"`
	TotalTaxAmount float64    `validate:""`
	Currency       string     `validate:"required,oneof=EUR USD"`
	Date           string     `validate:"required"`
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
	StoreName         string  `validate:"required"`
	ReferenceID       string  `validate:"required"`
	BillingDescriptor string  `validate:"required"`
	Siret             string  `validate:"required"`
	CodeApe           string  `validate:""`
	TvaIntra          string  `validate:""`
	Address           Address `validate:"required"`
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
	ReferenceID string    `validate:""`
	Name        string    `validate:"required"`
	Description string    `validate:""`
	Type        string    `validate:""`
	Quantity    int       `validate:"required"`
	Price       float64   `validate:"required"`
	Discount    float64   `validate:""`
	TotalAmount float64   `validate:""`
	Tax         Tax       `validate:"required"`
	Subitems    []Subitem `validate:""`
}

type Subitem struct {
	ReferenceID string  `validate:""`
	Name        string  `validate:"required"`
	Description string  `validate:""`
	Quantity    int     `validate:""`
	Price       float64 `validate:""`
	Discount    float64 `validate:""`
	TotalAmount float64 `validate:""`
	Tax         Tax     `validate:"required"`
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
