package invoice

import (
	"encoding/json"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Types of invoices (TYPE).
const (
	TypeOutgoing = "outgoing"
	TypeDraft    = "draft"
	TypeCredit   = "credit"
)

// Filter narrows invoice.get. Dates are YYYY-MM-DD.
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
	// Month and Year are deprecated by FastBill; use StartDate and EndDate.
	Month int `json:"MONTH,omitempty"`
	Year  int `json:"YEAR,omitempty"`
}

// Invoice is an invoice as invoice.get returns it.
type Invoice struct {
	InvoiceID              fastbill.ID            `json:"INVOICE_ID"`
	Type                   string                 `json:"TYPE"`
	Subtype                string                 `json:"SUBTYPE"`
	PreviousInvoiceID      fastbill.ID            `json:"PREVIOUS_INVOICE_ID"`
	NextInvoiceID          fastbill.ID            `json:"NEXT_INVOICE_ID"`
	CustomerID             fastbill.ID            `json:"CUSTOMER_ID"`
	CustomerNumber         string                 `json:"CUSTOMER_NUMBER"`
	CustomerCostCenterID   fastbill.ID            `json:"CUSTOMER_COSTCENTER_ID"`
	ContactID              fastbill.ID            `json:"CONTACT_ID"`
	ProjectID              fastbill.ID            `json:"PROJECT_ID"`
	SubscriptionID         fastbill.ID            `json:"SUBSCRIPTION_ID"`
	TemplateID             fastbill.ID            `json:"TEMPLATE_ID"`
	Organization           string                 `json:"ORGANIZATION"`
	Salutation             string                 `json:"SALUTATION"`
	FirstName              string                 `json:"FIRST_NAME"`
	LastName               string                 `json:"LAST_NAME"`
	Address                string                 `json:"ADDRESS"`
	Address2               string                 `json:"ADDRESS_2"`
	ZipCode                string                 `json:"ZIPCODE"`
	City                   string                 `json:"CITY"`
	CountryCode            string                 `json:"COUNTRY_CODE"`
	VatID                  string                 `json:"VAT_ID"`
	Comment                string                 `json:"COMMENT_"`
	Note                   string                 `json:"NOTE"`
	PaymentType            fastbill.Number        `json:"PAYMENT_TYPE"`
	DaysForPayment         fastbill.Number        `json:"DAYS_FOR_PAYMENT"`
	BankName               string                 `json:"BANK_NAME"`
	BankAccountNumber      string                 `json:"BANK_ACCOUNT_NUMBER"`
	BankCode               string                 `json:"BANK_CODE"`
	BankAccountOwner       string                 `json:"BANK_ACCOUNT_OWNER"`
	BankIBAN               string                 `json:"BANK_IBAN"`
	BankBIC                string                 `json:"BANK_BIC"`
	CustomerAccountingNote string                 `json:"CUSTOMER_ACCOUNTING_NOTE"`
	ContractReference      string                 `json:"CONTRACT_REFERENCE"`
	CustomerOrderReference string                 `json:"CUSTOMER_ORDER_REFERENCE"`
	OrderReference         string                 `json:"ORDER_REFERENCE"`
	Affiliate              string                 `json:"AFFILIATE"`
	CurrencyCode           string                 `json:"CURRENCY_CODE"`
	BaseCurrencyCode       string                 `json:"BASE_CURRENCY_CODE"`
	ExchangeRate           fastbill.Number        `json:"EXCHANGE_RATE"`
	InvoiceNumber          string                 `json:"INVOICE_NUMBER"`
	InvoiceTitle           string                 `json:"INVOICE_TITLE"`
	IntroText              string                 `json:"INTROTEXT"`
	InvoiceDate            string                 `json:"INVOICE_DATE"`
	DueDate                string                 `json:"DUE_DATE"`
	DeliveryDate           string                 `json:"DELIVERY_DATE"`
	ServicePeriodStart     string                 `json:"SERVICE_PERIOD_START"`
	ServicePeriodEnd       string                 `json:"SERVICE_PERIOD_END"`
	PaidDate               string                 `json:"PAID_DATE"`
	IsCanceled             fastbill.Flag          `json:"IS_CANCELED"`
	IsGross                fastbill.Flag          `json:"IS_GROSS"`
	State                  string                 `json:"STATE"`
	VatCase                string                 `json:"VAT_CASE"`
	CashDiscountPercent    fastbill.Number        `json:"CASH_DISCOUNT_PERCENT"`
	CashDiscountDays       fastbill.Number        `json:"CASH_DISCOUNT_DAYS"`
	SubTotal               fastbill.Number        `json:"SUB_TOTAL"`
	VatTotal               fastbill.Number        `json:"VAT_TOTAL"`
	Total                  fastbill.Number        `json:"TOTAL"`
	VatItems               fastbill.List[VatItem] `json:"VAT_ITEMS"`
	Items                  fastbill.List[Item]    `json:"ITEMS"`
	Payments               fastbill.List[Payment] `json:"PAYMENTS"`
	PaymentInfo            string                 `json:"PAYMENT_INFO"`
	LastUpdate             string                 `json:"LASTUPDATE"`
	DocumentURL            string                 `json:"DOCUMENT_URL"`
	DetailsURL             string                 `json:"DETAILS_URL"`
}

// Item is a line of an invoice.
type Item struct {
	InvoiceItemID fastbill.ID     `json:"INVOICE_ITEM_ID,omitempty"`
	ArticleNumber string          `json:"ARTICLE_NUMBER,omitempty"`
	Description   string          `json:"DESCRIPTION,omitempty"`
	Quantity      fastbill.Number `json:"QUANTITY,omitempty"`
	Unit          string          `json:"UNIT,omitempty"`
	UnitPrice     fastbill.Number `json:"UNIT_PRICE,omitempty"`
	VatPercent    fastbill.Number `json:"VAT_PERCENT,omitempty"`
	VatValue      fastbill.Number `json:"VAT_VALUE,omitempty"`
	CompleteNet   fastbill.Number `json:"COMPLETE_NET,omitempty"`
	CompleteGross fastbill.Number `json:"COMPLETE_GROSS,omitempty"`
	// Category is kept as sent; FastBill does not document its shape
	// (v1 saw a list).
	Category  json.RawMessage `json:"CATEGORY,omitempty"`
	SortOrder fastbill.Number `json:"SORT_ORDER,omitempty"`
}

// VatItem sums up one VAT rate of an invoice.
type VatItem struct {
	VatPercent  fastbill.Number `json:"VAT_PERCENT"`
	CompleteNet fastbill.Number `json:"COMPLETE_NET"`
	VatValue    fastbill.Number `json:"VAT_VALUE"`
}

// Payment is a payment booked on an invoice.
type Payment struct {
	PaymentID    fastbill.ID     `json:"PAYMENT_ID"`
	Date         string          `json:"DATE"`
	Amount       fastbill.Number `json:"AMOUNT"`
	CurrencyCode string          `json:"CURRENCY_CODE"`
	Note         string          `json:"NOTE"`
	Type         string          `json:"TYPE"`
}

// Request is the data of invoice.create and invoice.update.
type Request struct {
	// InvoiceID is required for Update.
	InvoiceID fastbill.ID `json:"INVOICE_ID,omitempty"`
	// DeleteExistingItems (1) replaces the items on Update.
	DeleteExistingItems  fastbill.Flag   `json:"DELETE_EXISTING_ITEMS,omitempty"`
	CustomerID           fastbill.ID     `json:"CUSTOMER_ID,omitempty"`
	CustomerCostCenterID fastbill.ID     `json:"CUSTOMER_COSTCENTER_ID,omitempty"`
	ContactID            fastbill.ID     `json:"CONTACT_ID,omitempty"`
	ProjectID            fastbill.ID     `json:"PROJECT_ID,omitempty"`
	Subtype              string          `json:"SUBTYPE,omitempty"`
	PreviousInvoiceID    fastbill.ID     `json:"PREVIOUS_INVOICE_ID,omitempty"`
	CurrencyCode         string          `json:"CURRENCY_CODE,omitempty"`
	TemplateID           fastbill.ID     `json:"TEMPLATE_ID,omitempty"`
	TemplateHash         string          `json:"TEMPLATE_HASH,omitempty"`
	IntroText            string          `json:"INTROTEXT,omitempty"`
	InvoiceTitle         string          `json:"INVOICE_TITLE,omitempty"`
	InvoiceDate          string          `json:"INVOICE_DATE,omitempty"`
	DeliveryDate         string          `json:"DELIVERY_DATE,omitempty"`
	ServicePeriodStart   string          `json:"SERVICE_PERIOD_START,omitempty"`
	ServicePeriodEnd     string          `json:"SERVICE_PERIOD_END,omitempty"`
	CashDiscountPercent  fastbill.Number `json:"CASH_DISCOUNT_PERCENT,omitempty"`
	CashDiscountDays     fastbill.Number `json:"CASH_DISCOUNT_DAYS,omitempty"`
	// VatCase is e.g. standard or small_business_regulation.
	VatCase                string        `json:"VAT_CASE,omitempty"`
	IsGross                fastbill.Flag `json:"IS_GROSS,omitempty"`
	CustomerAccountingNote string        `json:"CUSTOMER_ACCOUNTING_NOTE,omitempty"`
	ContractReference      string        `json:"CONTRACT_REFERENCE,omitempty"`
	CustomerOrderReference string        `json:"CUSTOMER_ORDER_REFERENCE,omitempty"`
	OrderReference         string        `json:"ORDER_REFERENCE,omitempty"`
	Items                  []Item        `json:"ITEMS,omitempty"`
}

// CreateResponse is the answer of invoice.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

// CompleteResponse is the answer of invoice.complete.
type CompleteResponse struct {
	Status        string `json:"STATUS"`
	InvoiceNumber string `json:"INVOICE_NUMBER"`
}

// SendByEmailRequest is the data of invoice.sendbyemail.
type SendByEmailRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
	Recipient Recipient   `json:"RECIPIENT"`
	Subject   string      `json:"SUBJECT,omitempty"`
	Message   string      `json:"MESSAGE,omitempty"`
	// ReceiptConfirmation (1) asks for a read receipt.
	ReceiptConfirmation fastbill.Flag `json:"RECEIPT_CONFIRMATION,omitempty"`
}

// Recipient are the addresses of an email.
type Recipient struct {
	To  string `json:"TO,omitempty"`
	Cc  string `json:"CC,omitempty"`
	Bcc string `json:"BCC,omitempty"`
}

// SendByPostResponse is the answer of invoice.sendbypost.
type SendByPostResponse struct {
	Status           string          `json:"STATUS"`
	RemainingCredits fastbill.Number `json:"REMAINING_CREDITS"`
}

// SetPaidRequest is the data of invoice.setpaid.
type SetPaidRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
	// PaidDate is YYYY-MM-DD; FastBill uses today when it is empty.
	PaidDate      string `json:"PAID_DATE,omitempty"`
	PaymentMethod string `json:"PAYMENT_METHOD,omitempty"`
}

// SetPaidResponse is the answer of invoice.setpaid.
type SetPaidResponse struct {
	Status        string `json:"STATUS"`
	InvoiceNumber string `json:"INVOICE_NUMBER"`
}

type idRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

type getResponse struct {
	Invoices fastbill.List[Invoice] `json:"INVOICES"`
}
