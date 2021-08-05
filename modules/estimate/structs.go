package estimate

//Filter available invoice filter
type Filter struct {
	CustomerID        string `json:"CUSTOMER_ID"`         // Customer ID
	EstimateID        string `json:"ESTIMATE_ID"`         // Estimate ID
	EstimateNumber    string `json:"ESTIMATE_NUMBER"`     // Estimate number
	StartEstimateDate string `json:"START_ESTIMATE_DATE"` // EEstimates from a specific date
	EndEstimateDate   string `json:"END_ESTIMATE_DATE"`   //Estimates until a specific date
}

//Estimate fastbill estimate
type Estimate struct {
	EstimateID           string    `json:"ESTIMATE_ID" mapstructure:"ESTIMATE_ID"`
	State                string    `json:"STATE" mapstructure:"STATE"`
	CustomerID           string    `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
	CustomerNumber       string    `json:"CUSTOMER_NUMBER" mapstructure:"CUSTOMER_NUMBER"`
	CustomerCostCenterID string    `json:"CUSTOMER_COSTCENTER_ID" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	ProjectID            string    `json:"PROJECT_ID" mapstructure:"PROJECT_ID"`
	Organization         string    `json:"ORGANIZATION" mapstructure:"ORGANIZATION"`
	Salutation           string    `json:"SALUTATION" mapstructure:"SALUTATION"`
	FirstName            string    `json:"FIRST_NAME" mapstructure:"FIRST_NAME"`
	LastName             string    `json:"LAST_NAME" mapstructure:"LAST_NAME"`
	Address              string    `json:"ADDRESS" mapstructure:"ADDRESS"`
	Address2             string    `json:"ADDRESS_2" mapstructure:"ADDRESS_2"`
	ZipCode              string    `json:"ZIPCODE" mapstructure:"ZIPCODE"`
	City                 string    `json:"CITY" mapstructure:"CITY"`
	InvoiceTitle         string    `json:"INVOICE_TITLE" mapstructure:"INVOICE_TITLE"`
	PaymentType          string    `json:"PAYMENT_TYPE" mapstructure:"PAYMENT_TYPE"`
	BankName             string    `json:"BANK_NAME" mapstructure:"BANK_NAME"`
	BankAccountNumber    string    `json:"BANK_ACCOUNT_NUMBER" mapstructure:"BANK_ACCOUNT_NUMBER"`
	BankCode             string    `json:"BANK_CODE" mapstructure:"BANK_CODE"`
	BankAccountOwner     string    `json:"BANK_ACCOUNT_OWNER" mapstructure:"BANK_ACCOUNT_OWNER"`
	BankIban             string    `json:"BANK_IBAN" mapstructure:"BANK_IBAN"`
	BankBic              string    `json:"BANK_BIC" mapstructure:"BANK_BIC"`
	CountryCode          string    `json:"COUNTRY_CODE" mapstructure:"COUNTRY_CODE"`
	VatID                string    `json:"VAT_ID" mapstructure:"VAT_ID"`
	CurrencyCode         string    `json:"CURRENCY_CODE" mapstructure:"CURRENCY_CODE"`
	TemplateID           string    `json:"TEMPLATE_ID" mapstructure:"TEMPLATE_ID"`
	EstimateNumber       string    `json:"ESTIMATE_NUMBER" mapstructure:"ESTIMATE_NUMBER"`
	IntroText            string    `json:"INTROTEXT" mapstructure:"INTROTEXT"`
	EstimateDate         string    `json:"ESTIMATE_DATE" mapstructure:"ESTIMATE_DATE"`
	DueDate              string    `json:"DUE_DATE" mapstructure:"DUE_DATE"`
	SubTotal             int       `json:"SUB_TOTAL" mapstructure:"SUB_TOTAL"`
	VatTotal             float64   `json:"VAT_TOTAL" mapstructure:"VAT_TOTAL"`
	VatItems             []VatItem `json:"VAT_ITEMS" mapstructure:"VAT_ITEMS"`
	Items                []Item    `json:"ITEMS" mapstructure:"ITEMS"`
	Total                float64   `json:"TOTAL" mapstructure:"TOTAL"`
	DocumentURL          string    `json:"DOCUMENT_URL" mapstructure:"DOCUMENT_URL"`
}

//Request fastbill estimate request
type Request struct {
	CustomerID           string `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerCostCenterID string `json:"CUSTOMER_COSTCENTER_ID,omitempty" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	TemplateID           string `json:"TEMPLATE_ID,omitempty" mapstructure:"TEMPLATE_ID"`
	TemplateHASH         string `json:"TEMPLATE_HASH,omitempty" mapstructure:"TEMPLATE_HASH"`
	Items                []Item `json:"ITEMS,omitempty" mapstructure:"ITEMS"`
}

//VatItem invoice vat item
type VatItem struct {
	VatPercent  string  `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	CompleteNet float64 `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	VatValue    float64 `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
}

// Item invoice item
type Item struct {
	EstimateItemID string        `json:"ESTIMATE_ITEM_ID,omitempty" mapstructure:"ESTIMATE_ITEM_ID"`
	ArticleNumber  string        `json:"ARTICLE_NUMBER,omitempty" mapstructure:"ARTICLE_NUMBER"`
	Description    string        `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Quantity       string        `json:"QUANTITY,omitempty" mapstructure:"QUANTITY"`
	UnitPrice      string        `json:"UNIT_PRICE,omitempty" mapstructure:"UNIT_PRICE"`
	VatPercent     string        `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	VatValue       float64       `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
	CompleteNet    float64       `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	CompleteGross  float64       `json:"COMPLETE_GROSS,omitempty" mapstructure:"COMPLETE_GROSS"`
	Category       []interface{} `json:"CATEGORY,omitempty" mapstructure:"CATEGORY"`
	SortOrder      int           `json:"SORT_ORDER,omitempty" mapstructure:"SORT_ORDER"`
}

//CreateResponse invoice api response
type CreateResponse struct {
	Status     string `json:"STATUS" mapstructure:"STATUS"`
	EstimateID int    `json:"ESTIMATE_ID" mapstructure:"ESTIMATE_ID"`
}

type createInvoiceRequest struct {
	EstimateID string `json:"ESTIMATE_ID" mapstructure:"ESTIMATE_ID"`
}

//CreateInvoiceResponse invoice api response
type CreateInvoiceResponse struct {
	InvoiceID int `json:"INVOICE_ID" mapstructure:"INVOICE_ID"`
}

type getResponse struct {
	Estimates []Estimate `json:"ESTIMATES" mapstructure:"ESTIMATES"`
}

type deleteRequest struct {
	EstimateID string `json:"ESTIMATE_ID" mapstructure:"ESTIMATE_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}

//SendByMailRequest send by mail api request
type SendByMailRequest struct {
	EstimateID          string               `json:"ESTIMATE_ID,omitempty" mapstructure:"ESTIMATE_ID"`
	Recipient           SendByMailRecipients `json:"RECIPIENT,omitempty" mapstructure:"RECIPIENT"`
	Subject             string               `json:"SUBJECT,omitempty" mapstructure:"SUBJECT"`
	Message             string               `json:"MESSAGE,omitempty" mapstructure:"MESSAGE"`
	ReceiptConfirmation string               `json:"RECEIPT_CONFIRMATION,omitempty" mapstructure:"RECEIPT_CONFIRMATION"`
}

//SendByMailRecipients recipient of the mail
type SendByMailRecipients struct {
	To  string `json:"TO,omitempty" mapstructure:"TO"`
	Cc  string `json:"CC,omitempty" mapstructure:"CC"`
	Bcc string `json:"BCC,omitempty" mapstructure:"BCC"`
}

type sendByMailResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
