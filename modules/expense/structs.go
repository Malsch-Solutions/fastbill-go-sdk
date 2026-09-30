package expense

import (
	"encoding/json"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Filter narrows expense.get. Dates are YYYY-MM-DD.
type Filter struct {
	InvoiceID     fastbill.ID `json:"INVOICE_ID,omitempty"`
	InvoiceNumber string      `json:"INVOICE_NUMBER,omitempty"`
	StartPaidDate string      `json:"START_PAID_DATE,omitempty"`
	EndPaidDate   string      `json:"END_PAID_DATE,omitempty"`
	StartDate     string      `json:"START_DATE,omitempty"`
	EndDate       string      `json:"END_DATE,omitempty"`
	// Deprecated: FastBill deprecated MONTH and YEAR; use StartDate and
	// EndDate.
	Month int `json:"MONTH,omitempty"`
	// Deprecated: see Month.
	Year int `json:"YEAR,omitempty"`
}

// Expense is an expense as expense.get returns it.
type Expense struct {
	InvoiceID          fastbill.ID     `json:"INVOICE_ID"`
	Organization       string          `json:"ORGANIZATION"`
	InvoiceNumber      string          `json:"INVOICE_NUMBER"`
	InvoiceDate        string          `json:"INVOICE_DATE"`
	ServicePeriodStart string          `json:"SERVICE_PERIOD_START"`
	ServicePeriodEnd   string          `json:"SERVICE_PERIOD_END"`
	DueDate            string          `json:"DUE_DATE"`
	ProjectID          fastbill.ID     `json:"PROJECT_ID"`
	CustomerID         fastbill.ID     `json:"CUSTOMER_ID"`
	SubTotal           fastbill.Number `json:"SUB_TOTAL"`
	VatTotal           fastbill.Number `json:"VAT_TOTAL"`
	Total              fastbill.Number `json:"TOTAL"`
	Note               string          `json:"NOTE"`
	Comment            string          `json:"COMMENT"`
	// Comments are the comments on the expense (not in the current docs,
	// kept from v1).
	Comments         fastbill.List[Comment] `json:"COMMENTS"`
	VatItems         fastbill.List[VatItem] `json:"VAT_ITEMS"`
	Items            fastbill.List[Item]    `json:"ITEMS"`
	PaidDate         string                 `json:"PAID_DATE"`
	CurrencyCode     string                 `json:"CURRENCY_CODE"`
	BaseCurrencyCode string                 `json:"BASE_CURRENCY_CODE"`
	// ExchangeRate converts to BaseCurrencyCode: base amount = amount ÷
	// ExchangeRate. It is 1 if both currencies are the same.
	ExchangeRate fastbill.Number `json:"EXCHANGE_RATE"`
	// Category is kept as sent: FastBill does not document its shape.
	Category    json.RawMessage `json:"CATEGORY"`
	PaymentInfo string          `json:"PAYMENT_INFO"`
	DocumentURL string          `json:"DOCUMENT_URL"`
}

// Comment is a comment on an expense.
type Comment struct {
	Date          string        `json:"DATE"`
	Comment       string        `json:"COMMENT"`
	CommentPublic fastbill.Flag `json:"COMMENT_PUBLIC"`
}

// VatItem sums up one VAT rate of an expense.
type VatItem struct {
	VatPercent  fastbill.Number `json:"VAT_PERCENT"`
	CompleteNet fastbill.Number `json:"COMPLETE_NET"`
	VatValue    fastbill.Number `json:"VAT_VALUE"`
}

// Item is a line of an expense.
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
	Category  json.RawMessage `json:"CATEGORY,omitempty"`
	SortOrder fastbill.Number `json:"SORT_ORDER,omitempty"`
}

// Request is the data of expense.create. InvoiceDate, Organization and
// SubTotal are required.
type Request struct {
	InvoiceDate        string          `json:"INVOICE_DATE,omitempty"`
	ServicePeriodStart string          `json:"SERVICE_PERIOD_START,omitempty"`
	ServicePeriodEnd   string          `json:"SERVICE_PERIOD_END,omitempty"`
	DueDate            string          `json:"DUE_DATE,omitempty"`
	ProjectID          fastbill.ID     `json:"PROJECT_ID,omitempty"`
	CustomerID         fastbill.ID     `json:"CUSTOMER_ID,omitempty"`
	Organization       string          `json:"ORGANIZATION,omitempty"`
	InvoiceNumber      string          `json:"INVOICE_NUMBER,omitempty"`
	Comment            string          `json:"COMMENT,omitempty"`
	SubTotal           fastbill.Number `json:"SUB_TOTAL,omitempty"`
	VatTotal           fastbill.Number `json:"VAT_TOTAL,omitempty"`
	Items              []Item          `json:"ITEMS,omitempty"`
}

// CreateResponse is the answer of expense.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

type getResponse struct {
	Expenses fastbill.List[Expense] `json:"EXPENSES"`
}
