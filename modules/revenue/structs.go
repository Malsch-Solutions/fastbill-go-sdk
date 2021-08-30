package revenue

import "time"

//Filter available revenue filter
type Filter struct {
	RevenueID     string     `json:"REVENUE_ID,omitempty"`     // Revenue ID
	RevenueNumber string     `json:"REVENUE_NUMBER,omitempty"` // Revenue number
	RevenueTitle  string     `json:"REVENUE_TITLE,omitempty"`  // Revenue title
	CustomerID    string     `json:"CUSTOMER_ID,omitempty"`    // A specific customer ID
	Month         int        `json:"MONTH,omitempty"`          // Month
	Year          int        `json:"YEAR,omitempty"`           // Year
	StartDueDate  *time.Time `json:"START_DUE_DATE,omitempty"` // Revenues that are due from a specific date
	EndDueDate    *time.Time `json:"END_DUE_DATE,omitempty"`   // Revenues that are due up to a specific date
	Type          string     `json:"TYPE,omitempty"`           // Payment type
}

//Revenue fastbill revenue
type Revenue struct {
	InvoiceID            string    `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"`
	Type                 string    `json:"TYPE,omitempty" mapstructure:"TYPE"`
	CustomerID           string    `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerNumber       string    `json:"CUSTOMER_NUMBER,omitempty" mapstructure:"CUSTOMER_NUMBER"`
	CustomerCostCenterID string    `json:"CUSTOMER_COSTCENTER_ID,omitempty" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	ContactID            string    `json:"CONTACT_ID,omitempty" mapstructure:"CONTACT_ID"`
	ProjectID            string    `json:"PROJECT_ID,omitempty" mapstructure:"REVENUE_ID"`
	CurrencyCode         string    `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	DeliveryDate         string    `json:"DELIVERY_DATE,omitempty" mapstructure:"DELIVERY_DATE"`
	InvoiceTitle         string    `json:"INVOICE_TITLE,omitempty" mapstructure:"INVOICE_TITLE"`
	CashDiscountPercent  string    `json:"CASH_DISCOUNT_PERCENT,omitempty" mapstructure:"CASH_DISCOUNT_PERCENT"`
	CashDiscountDays     string    `json:"CASH_DISCOUNT_DAYS,omitempty" mapstructure:"CASH_DISCOUNT_DAYS"`
	SubTotal             int       `json:"SUB_TOTAL,omitempty" mapstructure:"SUB_TOTAL"`
	VatTotal             float64   `json:"VAT_TOTAL,omitempty" mapstructure:"VAT_TOTAL"`
	VatCase              string    `json:"VAT_CASE,omitempty" mapstructure:"VAT_CASE"`
	VatItems             []VatItem `json:"VAT_ITEMS,omitempty" mapstructure:"VAT_ITEMS"`
	Items                []Item    `json:"ITEMS,omitempty" mapstructure:"ITEMS"`
	Total                float64   `json:"TOTAL,omitempty" mapstructure:"TOTAL"`
	Organization         string    `json:"ORGANIZATION,omitempty" mapstructure:"ORGANIZATION"`
	Note                 string    `json:"NOTE,omitempty" mapstructure:"NOTE"`
	Salutation           string    `json:"SALUTATION,omitempty" mapstructure:"SALUTATION"`
	FirstName            string    `json:"FIRST_NAME,omitempty" mapstructure:"FIRST_NAME"`
	LastName             string    `json:"LAST_NAME,omitempty" mapstructure:"LAST_NAME"`
	Address              string    `json:"ADDRESS,omitempty" mapstructure:"ADDRESS"`
	Address2             string    `json:"ADDRESS_2,omitempty" mapstructure:"ADDRESS_2"`
	ZipCode              string    `json:"ZIPCODE,omitempty" mapstructure:"ZIPCODE"`
	City                 string    `json:"CITY,omitempty" mapstructure:"CITY"`
	ServicePeriodStart   string    `json:"SERVICE_PERIOD_START,omitempty" mapstructure:"SERVICE_PERIOD_START"`
	ServicePeriodEnd     string    `json:"SERVICE_PERIOD_END,omitempty" mapstructure:"SERVICE_PERIOD_END"`
	PaymentType          string    `json:"PAYMENT_TYPE,omitempty" mapstructure:"PAYMENT_TYPE"`
	BankName             string    `json:"BANK_NAME,omitempty" mapstructure:"BANK_NAME"`
	BankAccountNumber    string    `json:"BANK_ACCOUNT_NUMBER,omitempty" mapstructure:"BANK_ACCOUNT_NUMBER"`
	BankCode             string    `json:"BANK_CODE,omitempty" mapstructure:"BANK_CODE"`
	BankAccountOwner     string    `json:"BANK_ACCOUNT_OWNER,omitempty" mapstructure:"BANK_ACCOUNT_OWNER"`
	BankIban             string    `json:"BANK_IBAN,omitempty" mapstructure:"BANK_IBAN"`
	BankBic              string    `json:"BANK_BIC,omitempty" mapstructure:"BANK_BIC"`
	CountryCode          string    `json:"COUNTRY_CODE,omitempty" mapstructure:"COUNTRY_CODE"`
	VatID                string    `json:"VAT_ID,omitempty" mapstructure:"VAT_ID"`
	TemplateID           string    `json:"TEMPLATE_ID,omitempty" mapstructure:"TEMPLATE_ID"`
	InvoiceNumber        string    `json:"INVOICE_NUMBER,omitempty" mapstructure:"INVOICE_NUMBER"`
	IntroText            string    `json:"INTROTEXT,omitempty" mapstructure:"INTROTEXT"`
	PaidDate             string    `json:"PAID_DATE,omitempty" mapstructure:"PAID_DATE"`
	IsCanceled           string    `json:"IS_CANCELED,omitempty" mapstructure:"IS_CANCELED"`
	InvoiceDate          string    `json:"INVOICE_DATE,omitempty" mapstructure:"INVOICE_DATE"`
	DueDate              string    `json:"DUE_DATE,omitempty" mapstructure:"DUE_DATE"`
	PaymentInfo          string    `json:"PAYMENT_INFO,omitempty" mapstructure:"PAYMENT_INFO"`
	Payments             []Payment `json:"PAYMENTS,omitempty" mapstructure:"PAYMENTS"`
	LastUpdate           string    `json:"LASTUPDATE,omitempty" mapstructure:"LASTUPDATE"`
	DocumentURL          string    `json:"DOCUMENT_URL,omitempty" mapstructure:"DOCUMENT_URL"`
	Comments             []Comment `json:"COMMENTS" mapstructure:"COMMENTS"`
}

//Comment fastbill api expense comment
type Comment struct {
	Date          string `json:"DATE,omitempty" mapstructure:"DATE"`
	Comment       string `json:"COMMENT,omitempty" mapstructure:"COMMENT"`
	CommentPublic string `json:"COMMENT_PUBLIC,omitempty" mapstructure:"COMMENT_PUBLIC"`
}

//Request revenue create request
type Request struct {
	InvoiceDate   string  `json:"INVOICE_DATE,omitempty" mapstructure:"INVOICE_DATE"`
	DueDate       string  `json:"DUE_DATE,omitempty" mapstructure:"DUE_DATE"`
	CustomerID    string  `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	InvoiceNumber string  `json:"INVOICE_NUMBER,omitempty" mapstructure:"INVOICE_NUMBER"`
	Comment       string  `json:"COMMENT,omitempty" mapstructure:"COMMENT"`
	SubTotal      int     `json:"SUB_TOTAL,omitempty" mapstructure:"SUB_TOTAL"`
	VatTotal      float64 `json:"VAT_TOTAL,omitempty" mapstructure:"VAT_TOTAL"`
}

// Item revenue item
type Item struct {
	InvoiceItemID int           `json:"INVOICE_ITEM_ID,omitempty" mapstructure:"INVOICE_ITEM_ID"`
	ArticleNumber string        `json:"ARTICLE_NUMBER,omitempty" mapstructure:"ARTICLE_NUMBER"`
	Description   string        `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Quantity      int           `json:"QUANTITY,omitempty" mapstructure:"QUANTITY"`
	UnitPrice     float64       `json:"UNIT_PRICE,omitempty" mapstructure:"UNIT_PRICE"`
	VatPercent    int           `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	VatValue      float64       `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
	CompleteNet   int           `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	CompleteGross float64       `json:"COMPLETE_GROSS,omitempty" mapstructure:"COMPLETE_GROSS"`
	Category      []interface{} `json:"CATEGORY,omitempty" mapstructure:"CATEGORY"`
	SortOrder     int           `json:"SORT_ORDER,omitempty" mapstructure:"SORT_ORDER"`
	CurrencyCode  string        `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
}

// Payment revenue payment
type Payment struct {
	PaymentID    string `json:"PAYMENT_ID,omitempty" mapstructure:"PAYMENT_ID"`
	Date         string `json:"DATE,omitempty" mapstructure:"DATE"`
	Amount       string `json:"AMOUNT,omitempty" mapstructure:"AMOUNT"`
	CurrencyCode string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	Note         string `json:"NOTE,omitempty" mapstructure:"NOTE"`
	Type         string `json:"TYPE,omitempty" mapstructure:"TYPE"`
}

//VatItem revenue vat item
type VatItem struct {
	VatPercent  int     `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	CompleteNet float64 `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	VatValue    float64 `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
}

//CreateResponse revenue api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	InvoiceID int    `json:"INVOICE_ID" mapstructure:"INVOICE_ID"`
}

//SetPaidRequest set paid api request
type SetPaidRequest struct {
	InvoiceID string    `json:"INVOICE_ID,omitempty"`
	PaidDate  time.Time `json:"PAID_DATE,omitempty"`
}

//SetPaidResponse revenue api response
type SetPaidResponse struct {
	Status        string `json:"STATUS" mapstructure:"STATUS"`
	InvoiceNumber string `json:"INVOICE_NUMBER" mapstructure:"INVOICE_NUMBER"`
}

type getResponse struct {
	Revenues []Revenue `json:"REVENUES" mapstructure:"REVENUES"`
}

type deleteRequest struct {
	RevenueID string `json:"REVENUE_ID" mapstructure:"REVENUE_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
