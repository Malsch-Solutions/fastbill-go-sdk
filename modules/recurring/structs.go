package recurring

//Filter available recurring filter
type Filter struct {
	InvoiceID string `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"` //A specific recurring ID}
}

//Recurring fastbill recurring invoice definition
type Recurring struct {
	InvoiceID            string    `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"`
	Type                 string    `json:"TYPE,omitempty" mapstructure:"TYPE"`
	CustomerID           string    `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerNumber       string    `json:"CUSTOMER_NUMBER,omitempty" mapstructure:"CUSTOMER_NUMBER"`
	CustomerCostCenterID string    `json:"CUSTOMER_COSTCENTER_ID,omitempty" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	ContactID            string    `json:"CONTACT_ID,omitempty" mapstructure:"CONTACT_ID"`
	ProjectID            string    `json:"PROJECT_ID,omitempty" mapstructure:"PROJECT_ID"`
	CurrencyCode         string    `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	DeliveryDate         string    `json:"DELIVERY_DATE,omitempty" mapstructure:"DELIVERY_DATE"`
	InvoiceTitle         string    `json:"INVOICE_TITLE,omitempty" mapstructure:"INVOICE_TITLE"`
	CashDiscountPercent  string    `json:"CASH_DISCOUNT_PERCENT,omitempty" mapstructure:"CASH_DISCOUNT_PERCENT"`
	CashDiscountDays     string    `json:"CASH_DISCOUNT_DAYS,omitempty" mapstructure:"CASH_DISCOUNT_DAYS"`
	SubTotal             float64   `json:"SUB_TOTAL,omitempty" mapstructure:"SUB_TOTAL"`
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
	TemplateID           string    `json:"TEMPLATE_ID,omitempty" mapstructure:"TEMPLATE_ID"`
	Occurences           string    `json:"OCCURENCES,omitempty" mapstructure:"OCCURENCES"`
	Frequency            string    `json:"FREQUENCY,omitempty" mapstructure:"FREQUENCY"`
	StartDate            string    `json:"START_DATE,omitempty" mapstructure:"START_DATE"`
	EmailNotify          string    `json:"EMAIL_NOTIFY,omitempty" mapstructure:"EMAIL_NOTIFY"`
	OutputType           string    `json:"OUTPUT_TYPE,omitempty" mapstructure:"OUTPUT_TYPE"`
	IntroText            string    `json:"INTROTEXT,omitempty" mapstructure:"INTROTEXT"`
}

//Request fastbill recurring invoice definition
type Request struct {
	InvoiceID            string `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"`
	DeleteExistingItems  int    `json:"DELETE_EXISTING_ITEMS,omitempty" mapstructure:"DELETE_EXISTING_ITEMS"`
	CustomerID           string `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerCostCenterID string `json:"CUSTOMER_COSTCENTER_ID,omitempty" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	CurrencyCode         string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	TemplateID           string `json:"TEMPLATE_ID,omitempty" mapstructure:"TEMPLATE_ID"`
	TemplateHASH         string `json:"TEMPLATE_HASH,omitempty" mapstructure:"TEMPLATE_HASH"`
	IntroText            string `json:"INTROTEXT,omitempty" mapstructure:"INTROTEXT"`
	Frequency            string `json:"FREQUENCY,omitempty" mapstructure:"FREQUENCY"` // Required	Frequency of invoice run: Weekly | 2 weeks | 4 weeks | monthly | 2 months | 3 months | 6 months | yearly | 2 years
	StartDate            string `json:"START_DATE,omitempty" mapstructure:"START_DATE"`
	Occurences           string `json:"OCCURENCES,omitempty" mapstructure:"OCCURENCES"`
	OutputType           string `json:"OUTPUT_TYPE,omitempty" mapstructure:"OUTPUT_TYPE"` // Type of automatically created invoice: draft | Outgoing
	EmailNotify          string `json:"EMAIL_NOTIFY,omitempty" mapstructure:"EMAIL_NOTIFY"`
	DeliveryDate         string `json:"DELIVERY_DATE,omitempty" mapstructure:"DELIVERY_DATE"`
	CashDiscountPercent  string `json:"CASH_DISCOUNT_PERCENT,omitempty" mapstructure:"CASH_DISCOUNT_PERCENT"`
	CashDiscountDays     string `json:"CASH_DISCOUNT_DAYS,omitempty" mapstructure:"CASH_DISCOUNT_DAYS"`
	VatCase              string `json:"VAT_CASE,omitempty" mapstructure:"VAT_CASE"`
	Items                []Item `json:"ITEMS,omitempty" mapstructure:"ITEMS"`
}

//VatItem fastbill recurring vat item
type VatItem struct {
	VatPercent  int     `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	CompleteNet float64 `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	VatValue    float64 `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
}

//Item fastbill recurring item
type Item struct {
	InvoiceItemID int           `json:"INVOICE_ITEM_ID,omitempty" mapstructure:"INVOICE_ITEM_ID"`
	ArticleNumber string        `json:"ARTICLE_NUMBER,omitempty" mapstructure:"ARTICLE_NUMBER"`
	Description   string        `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Quantity      int           `json:"QUANTITY,omitempty" mapstructure:"QUANTITY"`
	UnitPrice     float64       `json:"UNIT_PRICE,omitempty" mapstructure:"UNIT_PRICE"`
	VatPercent    int           `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	VatValue      float64       `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
	CompleteNet   float64       `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	CompleteGross float64       `json:"COMPLETE_GROSS,omitempty" mapstructure:"COMPLETE_GROSS"`
	Category      []interface{} `json:"CATEGORY,omitempty" mapstructure:"CATEGORY"`
	SortOrder     int           `json:"SORT_ORDER,omitempty" mapstructure:"SORT_ORDER"`
}

//CreateResponse recurring api response
type CreateResponse struct {
	Status    string `json:"STATUS,omitempty" mapstructure:"STATUS"`
	InvoiceID int    `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"`
}

//UpdateResponse recurring api response
type UpdateResponse struct {
	Status string `json:"STATUS,omitempty" mapstructure:"STATUS"`
}

type getResponse struct {
	Recurrings []Recurring `json:"INVOICES,omitempty" mapstructure:"INVOICES"`
}

type deleteRequest struct {
	InvoiceID string `json:"INVOICE_ID,omitempty" mapstructure:"INVOICE_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS,omitempty" mapstructure:"STATUS"`
}
