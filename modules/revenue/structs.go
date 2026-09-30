package revenue

import (
	"encoding/json"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Filter narrows revenue.get; FastBill takes the filters of invoice.get.
// Dates are YYYY-MM-DD.
type Filter struct {
	InvoiceID       fastbill.ID `json:"INVOICE_ID,omitempty"`
	InvoiceNumber   string      `json:"INVOICE_NUMBER,omitempty"`
	InvoiceTitle    string      `json:"INVOICE_TITLE,omitempty"`
	CustomerID      fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	StartDueDate    string      `json:"START_DUE_DATE,omitempty"`
	EndDueDate      string      `json:"END_DUE_DATE,omitempty"`
	StartPaidDate   string      `json:"START_PAID_DATE,omitempty"`
	EndPaidDate     string      `json:"END_PAID_DATE,omitempty"`
	StartDate       string      `json:"START_DATE,omitempty"`
	EndDate         string      `json:"END_DATE,omitempty"`
	StartLastUpdate string      `json:"START_LASTUPDATE,omitempty"`
	EndLastUpdate   string      `json:"END_LASTUPDATE,omitempty"`
	// Type is outgoing, draft or credit.
	Type string `json:"TYPE,omitempty"`
	// Subtype is standard, progress_invoice or final_invoice.
	Subtype string        `json:"SUBTYPE,omitempty"`
	IsGross fastbill.Flag `json:"IS_GROSS,omitempty"`
	// Deprecated: FastBill deprecated MONTH and YEAR; use StartDate and
	// EndDate.
	Month int `json:"MONTH,omitempty"`
	// Deprecated: see Month.
	Year int `json:"YEAR,omitempty"`
}

// Revenue is a revenue as revenue.get returns it.
type Revenue struct {
	InvoiceID            fastbill.ID `json:"INVOICE_ID"`
	Type                 string      `json:"TYPE"`
	CustomerID           fastbill.ID `json:"CUSTOMER_ID"`
	CustomerNumber       string      `json:"CUSTOMER_NUMBER"`
	CustomerCostCenterID fastbill.ID `json:"CUSTOMER_COSTCENTER_ID"`
	ContactID            fastbill.ID `json:"CONTACT_ID"`
	ProjectID            fastbill.ID `json:"PROJECT_ID"`
	CurrencyCode         string      `json:"CURRENCY_CODE"`
	BaseCurrencyCode     string      `json:"BASE_CURRENCY_CODE"`
	// ExchangeRate converts to BaseCurrencyCode: base amount = amount ÷
	// ExchangeRate. It is 1 if both currencies are the same.
	ExchangeRate        fastbill.Number        `json:"EXCHANGE_RATE"`
	DeliveryDate        string                 `json:"DELIVERY_DATE"`
	InvoiceTitle        string                 `json:"INVOICE_TITLE"`
	CashDiscountPercent fastbill.Number        `json:"CASH_DISCOUNT_PERCENT"`
	CashDiscountDays    fastbill.Number        `json:"CASH_DISCOUNT_DAYS"`
	SubTotal            fastbill.Number        `json:"SUB_TOTAL"`
	VatTotal            fastbill.Number        `json:"VAT_TOTAL"`
	VatCase             string                 `json:"VAT_CASE"`
	VatItems            fastbill.List[VatItem] `json:"VAT_ITEMS"`
	Items               fastbill.List[Item]    `json:"ITEMS"`
	Total               fastbill.Number        `json:"TOTAL"`
	Organization        string                 `json:"ORGANIZATION"`
	Note                string                 `json:"NOTE"`
	Salutation          string                 `json:"SALUTATION"`
	FirstName           string                 `json:"FIRST_NAME"`
	LastName            string                 `json:"LAST_NAME"`
	Address             string                 `json:"ADDRESS"`
	Address2            string                 `json:"ADDRESS_2"`
	ZipCode             string                 `json:"ZIPCODE"`
	City                string                 `json:"CITY"`
	ServicePeriodStart  string                 `json:"SERVICE_PERIOD_START"`
	ServicePeriodEnd    string                 `json:"SERVICE_PERIOD_END"`
	// PaymentType is 1 transfer, 2 direct debit, 3 cash, 4 PayPal,
	// 5 advance payment, 6 credit card.
	PaymentType       fastbill.Number        `json:"PAYMENT_TYPE"`
	BankName          string                 `json:"BANK_NAME"`
	BankAccountNumber string                 `json:"BANK_ACCOUNT_NUMBER"`
	BankCode          string                 `json:"BANK_CODE"`
	BankAccountOwner  string                 `json:"BANK_ACCOUNT_OWNER"`
	BankIBAN          string                 `json:"BANK_IBAN"`
	BankBIC           string                 `json:"BANK_BIC"`
	CountryCode       string                 `json:"COUNTRY_CODE"`
	VatID             string                 `json:"VAT_ID"`
	TemplateID        fastbill.ID            `json:"TEMPLATE_ID"`
	InvoiceNumber     string                 `json:"INVOICE_NUMBER"`
	IntroText         string                 `json:"INTROTEXT"`
	PaidDate          string                 `json:"PAID_DATE"`
	IsCanceled        fastbill.Flag          `json:"IS_CANCELED"`
	InvoiceDate       string                 `json:"INVOICE_DATE"`
	DueDate           string                 `json:"DUE_DATE"`
	PaymentInfo       string                 `json:"PAYMENT_INFO"`
	Payments          fastbill.List[Payment] `json:"PAYMENTS"`
	LastUpdate        string                 `json:"LASTUPDATE"`
	DocumentURL       string                 `json:"DOCUMENT_URL"`
	Comments          fastbill.List[Comment] `json:"COMMENTS"`
}

// Comment is a comment on a revenue.
type Comment struct {
	Date          string        `json:"DATE"`
	Comment       string        `json:"COMMENT"`
	CommentPublic fastbill.Flag `json:"COMMENT_PUBLIC"`
}

// Item is a line of a revenue.
type Item struct {
	InvoiceItemID fastbill.ID     `json:"INVOICE_ITEM_ID,omitempty"`
	ArticleNumber string          `json:"ARTICLE_NUMBER,omitempty"`
	Description   string          `json:"DESCRIPTION,omitempty"`
	Quantity      fastbill.Number `json:"QUANTITY,omitempty"`
	UnitPrice     fastbill.Number `json:"UNIT_PRICE,omitempty"`
	VatPercent    fastbill.Number `json:"VAT_PERCENT,omitempty"`
	VatValue      fastbill.Number `json:"VAT_VALUE,omitempty"`
	CompleteNet   fastbill.Number `json:"COMPLETE_NET,omitempty"`
	CompleteGross fastbill.Number `json:"COMPLETE_GROSS,omitempty"`
	// Category is kept as sent: FastBill does not document its shape.
	Category     json.RawMessage `json:"CATEGORY,omitempty"`
	SortOrder    fastbill.Number `json:"SORT_ORDER,omitempty"`
	CurrencyCode string          `json:"CURRENCY_CODE,omitempty"`
}

// Payment is a payment booked on a revenue.
type Payment struct {
	PaymentID    fastbill.ID     `json:"PAYMENT_ID"`
	Date         string          `json:"DATE"`
	Amount       fastbill.Number `json:"AMOUNT"`
	CurrencyCode string          `json:"CURRENCY_CODE"`
	Note         string          `json:"NOTE"`
	Type         string          `json:"TYPE"`
}

// VatItem sums up one VAT rate of a revenue.
type VatItem struct {
	VatPercent  fastbill.Number `json:"VAT_PERCENT"`
	CompleteNet fastbill.Number `json:"COMPLETE_NET"`
	VatValue    fastbill.Number `json:"VAT_VALUE"`
}

// Request is the data of revenue.create. InvoiceDate, CustomerID and
// SubTotal are required.
type Request struct {
	InvoiceDate   string          `json:"INVOICE_DATE,omitempty"`
	DueDate       string          `json:"DUE_DATE,omitempty"`
	CustomerID    fastbill.ID     `json:"CUSTOMER_ID,omitempty"`
	InvoiceNumber string          `json:"INVOICE_NUMBER,omitempty"`
	Comment       string          `json:"COMMENT,omitempty"`
	SubTotal      fastbill.Number `json:"SUB_TOTAL,omitempty"`
	VatTotal      fastbill.Number `json:"VAT_TOTAL,omitempty"`
	CurrencyCode  string          `json:"CURRENCY_CODE,omitempty"`
	Items         []Item          `json:"ITEMS,omitempty"`
}

// CreateResponse is the answer of revenue.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

// SetPaidRequest is the data of revenue.setpaid.
type SetPaidRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
	// PaidDate is YYYY-MM-DD.
	PaidDate string `json:"PAID_DATE,omitempty"`
}

// SetPaidResponse is the answer of revenue.setpaid.
type SetPaidResponse struct {
	Status        string `json:"STATUS"`
	InvoiceNumber string `json:"INVOICE_NUMBER"`
}

type idRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

type getResponse struct {
	Revenues fastbill.List[Revenue] `json:"REVENUES"`
}
