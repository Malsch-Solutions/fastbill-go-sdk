package estimate

import (
	"encoding/json"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// States of an estimate (STATE).
const (
	StateCreated     = "a"
	StateSent        = "b"
	StateNegotiation = "c"
	StateAccepted    = "d"
	StateRejected    = "e"
	StateInvoiced    = "f"
)

// Filter narrows estimate.get. Dates are YYYY-MM-DD.
type Filter struct {
	CustomerID     fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	EstimateID     fastbill.ID `json:"ESTIMATE_ID,omitempty"`
	EstimateNumber string      `json:"ESTIMATE_NUMBER,omitempty"`
	StartDate      string      `json:"START_DATE,omitempty"`
	EndDate        string      `json:"END_DATE,omitempty"`
}

// Estimate is an estimate as estimate.get returns it.
type Estimate struct {
	EstimateID fastbill.ID `json:"ESTIMATE_ID"`
	// State is one of the State constants.
	State                string      `json:"STATE"`
	CustomerID           fastbill.ID `json:"CUSTOMER_ID"`
	CustomerNumber       string      `json:"CUSTOMER_NUMBER"`
	CustomerCostCenterID fastbill.ID `json:"CUSTOMER_COSTCENTER_ID"`
	ProjectID            fastbill.ID `json:"PROJECT_ID"`
	Organization         string      `json:"ORGANIZATION"`
	Salutation           string      `json:"SALUTATION"`
	FirstName            string      `json:"FIRST_NAME"`
	LastName             string      `json:"LAST_NAME"`
	Address              string      `json:"ADDRESS"`
	Address2             string      `json:"ADDRESS_2"`
	ZipCode              string      `json:"ZIPCODE"`
	City                 string      `json:"CITY"`
	InvoiceTitle         string      `json:"INVOICE_TITLE"`
	// PaymentType is 1 transfer, 2 direct debit, 3 cash, 4 PayPal,
	// 5 advance payment, 6 credit card.
	PaymentType       fastbill.Number `json:"PAYMENT_TYPE"`
	BankName          string          `json:"BANK_NAME"`
	BankAccountNumber string          `json:"BANK_ACCOUNT_NUMBER"`
	BankCode          string          `json:"BANK_CODE"`
	BankAccountOwner  string          `json:"BANK_ACCOUNT_OWNER"`
	BankIBAN          string          `json:"BANK_IBAN"`
	BankBIC           string          `json:"BANK_BIC"`
	CountryCode       string          `json:"COUNTRY_CODE"`
	VatID             string          `json:"VAT_ID"`
	CurrencyCode      string          `json:"CURRENCY_CODE"`
	BaseCurrencyCode  string          `json:"BASE_CURRENCY_CODE"`
	// ExchangeRate converts to BaseCurrencyCode: base amount = amount ÷
	// ExchangeRate. It is 1 if both currencies are the same.
	ExchangeRate   fastbill.Number        `json:"EXCHANGE_RATE"`
	TemplateID     fastbill.ID            `json:"TEMPLATE_ID"`
	EstimateNumber string                 `json:"ESTIMATE_NUMBER"`
	IntroText      string                 `json:"INTROTEXT"`
	EstimateDate   string                 `json:"ESTIMATE_DATE"`
	DueDate        string                 `json:"DUE_DATE"`
	SubTotal       fastbill.Number        `json:"SUB_TOTAL"`
	VatTotal       fastbill.Number        `json:"VAT_TOTAL"`
	VatItems       fastbill.List[VatItem] `json:"VAT_ITEMS"`
	Items          fastbill.List[Item]    `json:"ITEMS"`
	Total          fastbill.Number        `json:"TOTAL"`
	DocumentURL    string                 `json:"DOCUMENT_URL"`
}

// Item is a line of an estimate.
type Item struct {
	EstimateItemID fastbill.ID     `json:"ESTIMATE_ITEM_ID,omitempty"`
	ArticleNumber  string          `json:"ARTICLE_NUMBER,omitempty"`
	Description    string          `json:"DESCRIPTION,omitempty"`
	Quantity       fastbill.Number `json:"QUANTITY,omitempty"`
	UnitPrice      fastbill.Number `json:"UNIT_PRICE,omitempty"`
	VatPercent     fastbill.Number `json:"VAT_PERCENT,omitempty"`
	VatValue       fastbill.Number `json:"VAT_VALUE,omitempty"`
	CompleteNet    fastbill.Number `json:"COMPLETE_NET,omitempty"`
	CompleteGross  fastbill.Number `json:"COMPLETE_GROSS,omitempty"`
	// Category is kept as sent: FastBill does not document its shape.
	Category  json.RawMessage `json:"CATEGORY,omitempty"`
	SortOrder fastbill.Number `json:"SORT_ORDER,omitempty"`
}

// VatItem sums up one VAT rate of an estimate.
type VatItem struct {
	VatPercent  fastbill.Number `json:"VAT_PERCENT"`
	CompleteNet fastbill.Number `json:"COMPLETE_NET"`
	VatValue    fastbill.Number `json:"VAT_VALUE"`
}

// Request is the data of estimate.create. CustomerID and Items are
// required.
type Request struct {
	CustomerID           fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	CustomerCostCenterID fastbill.ID `json:"CUSTOMER_COSTCENTER_ID,omitempty"`
	TemplateID           fastbill.ID `json:"TEMPLATE_ID,omitempty"`
	TemplateHash         string      `json:"TEMPLATE_HASH,omitempty"`
	Items                []Item      `json:"ITEMS,omitempty"`
}

// CreateResponse is the answer of estimate.create.
type CreateResponse struct {
	Status     string      `json:"STATUS"`
	EstimateID fastbill.ID `json:"ESTIMATE_ID"`
}

// CreateInvoiceResponse is the answer of estimate.createinvoice.
type CreateInvoiceResponse struct {
	Status    string      `json:"STATUS"`
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

// SendByEmailRequest is the data of estimate.sendbyemail.
type SendByEmailRequest struct {
	EstimateID fastbill.ID `json:"ESTIMATE_ID"`
	Recipient  Recipient   `json:"RECIPIENT"`
	Subject    string      `json:"SUBJECT,omitempty"`
	Message    string      `json:"MESSAGE,omitempty"`
	// ReceiptConfirmation (1) asks for a read receipt.
	ReceiptConfirmation fastbill.Flag `json:"RECEIPT_CONFIRMATION,omitempty"`
}

// Recipient are the addresses of an email.
type Recipient struct {
	To  string `json:"TO,omitempty"`
	Cc  string `json:"CC,omitempty"`
	Bcc string `json:"BCC,omitempty"`
}

type idRequest struct {
	EstimateID fastbill.ID `json:"ESTIMATE_ID"`
}

type getResponse struct {
	Estimates fastbill.List[Estimate] `json:"ESTIMATES"`
}
