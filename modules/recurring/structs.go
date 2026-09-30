package recurring

import (
	"encoding/json"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Frequencies of a recurring invoice (FREQUENCY).
const (
	FrequencyWeekly  = "weekly"
	Frequency2Weeks  = "2 weeks"
	Frequency4Weeks  = "4 weeks"
	FrequencyMonthly = "monthly"
	Frequency2Months = "2 months"
	Frequency3Months = "3 months"
	Frequency6Months = "6 months"
	FrequencyYearly  = "yearly"
)

// Filter narrows recurring.get.
type Filter struct {
	InvoiceID fastbill.ID   `json:"INVOICE_ID,omitempty"`
	IsGross   fastbill.Flag `json:"IS_GROSS,omitempty"`
}

// Recurring is a recurring invoice as recurring.get returns it.
type Recurring struct {
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
	InvoiceNumber       string                 `json:"INVOICE_NUMBER"`
	InvoiceDate         string                 `json:"INVOICE_DATE"`
	DueDate             string                 `json:"DUE_DATE"`
	PaidDate            string                 `json:"PAID_DATE"`
	IsCanceled          fastbill.Flag          `json:"IS_CANCELED"`
	IsGross             fastbill.Flag          `json:"IS_GROSS"`
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
	PaymentType       fastbill.Number `json:"PAYMENT_TYPE"`
	BankName          string          `json:"BANK_NAME"`
	BankAccountNumber string          `json:"BANK_ACCOUNT_NUMBER"`
	BankCode          string          `json:"BANK_CODE"`
	BankAccountOwner  string          `json:"BANK_ACCOUNT_OWNER"`
	BankIBAN          string          `json:"BANK_IBAN"`
	BankBIC           string          `json:"BANK_BIC"`
	TemplateID        fastbill.ID     `json:"TEMPLATE_ID"`
	// Occurrences is the number of invoices to write, 0 for unlimited.
	Occurrences fastbill.Number `json:"OCCURENCES"`
	// Frequency is one of the Frequency constants.
	Frequency   string        `json:"FREQUENCY"`
	StartDate   string        `json:"START_DATE"`
	EmailNotify fastbill.Flag `json:"EMAIL_NOTIFY"`
	// OutputType is draft or outgoing.
	OutputType string `json:"OUTPUT_TYPE"`
	IntroText  string `json:"INTROTEXT"`
}

// Item is a line of a recurring invoice.
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

// VatItem sums up one VAT rate of a recurring invoice.
type VatItem struct {
	VatPercent  fastbill.Number `json:"VAT_PERCENT"`
	CompleteNet fastbill.Number `json:"COMPLETE_NET"`
	VatValue    fastbill.Number `json:"VAT_VALUE"`
}

// Request is the data of recurring.create and recurring.update. Create
// requires CustomerID, StartDate, Frequency, OutputType and Items.
type Request struct {
	// InvoiceID is required for Update.
	InvoiceID fastbill.ID `json:"INVOICE_ID,omitempty"`
	// DeleteExistingItems (1) replaces the items on Update.
	DeleteExistingItems  fastbill.Flag `json:"DELETE_EXISTING_ITEMS,omitempty"`
	CustomerID           fastbill.ID   `json:"CUSTOMER_ID,omitempty"`
	CustomerCostCenterID fastbill.ID   `json:"CUSTOMER_COSTCENTER_ID,omitempty"`
	CurrencyCode         string        `json:"CURRENCY_CODE,omitempty"`
	TemplateID           fastbill.ID   `json:"TEMPLATE_ID,omitempty"`
	TemplateHash         string        `json:"TEMPLATE_HASH,omitempty"`
	IntroText            string        `json:"INTROTEXT,omitempty"`
	// StartDate is the first run, YYYY-MM-DD.
	StartDate string `json:"START_DATE,omitempty"`
	// Frequency is one of the Frequency constants.
	Frequency string `json:"FREQUENCY,omitempty"`
	// Occurrences is the number of invoices to write, 0 for unlimited.
	Occurrences fastbill.Number `json:"OCCURENCES,omitempty"`
	// OutputType is draft or outgoing.
	OutputType string `json:"OUTPUT_TYPE,omitempty"`
	// EmailNotify (1) sends each invoice to the account's own address.
	EmailNotify         fastbill.Flag   `json:"EMAIL_NOTIFY,omitempty"`
	DeliveryDate        string          `json:"DELIVERY_DATE,omitempty"`
	CashDiscountPercent fastbill.Number `json:"CASH_DISCOUNT_PERCENT,omitempty"`
	CashDiscountDays    fastbill.Number `json:"CASH_DISCOUNT_DAYS,omitempty"`
	// VatCase is e.g. standard or small_business_regulation.
	VatCase string `json:"VAT_CASE,omitempty"`
	Items   []Item `json:"ITEMS,omitempty"`
}

// CreateResponse is the answer of recurring.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

type idRequest struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID"`
}

type getResponse struct {
	Recurrings fastbill.List[Recurring] `json:"INVOICES"`
}
