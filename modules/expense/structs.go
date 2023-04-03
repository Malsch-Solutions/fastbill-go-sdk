package expense

// Filter available customer filter
type Filter struct {
	InvoiceID     string `json:"INVOICE_ID,omitempty"`     //Invoice ID
	InvoiceNumber string `json:"INVOICE_NUMBER,omitempty"` //Invoice Number
	Month         string `json:"MONTH,omitempty"`          //Month
	Year          string `json:"YEAR,omitempty"`           //Year
}

// Expense fastbill expense
type Expense struct {
	InvoiceID          string      `json:"INVOICE_ID" mapstructure:"INVOICE_ID"`
	Organization       string      `json:"ORGANIZATION" mapstructure:"ORGANIZATION"`
	Category           interface{} `json:"CATEGORY" mapstructure:"CATEGORY"`
	InvoiceNumber      string      `json:"INVOICE_NUMBER" mapstructure:"INVOICE_NUMBER"`
	InvoiceDate        string      `json:"INVOICE_DATE" mapstructure:"INVOICE_DATE"`
	ServicePeriodStart string      `json:"SERVICE_PERIOD_START" mapstructure:"SERVICE_PERIOD_START"`
	ServicePeriodEnd   string      `json:"SERVICE_PERIOD_END" mapstructure:"SERVICE_PERIOD_END"`
	DueDate            string      `json:"DUE_DATE" mapstructure:"DUE_DATE"`
	SubTotal           int         `json:"SUB_TOTAL" mapstructure:"SUB_TOTAL"`
	VatTotal           float64     `json:"VAT_TOTAL" mapstructure:"VAT_TOTAL"`
	Total              float64     `json:"TOTAL" mapstructure:"TOTAL"`
	DocumentURL        string      `json:"DOCUMENT_URL" mapstructure:"DOCUMENT_URL"`
	PaidDate           string      `json:"PAID_DATE" mapstructure:"PAID_DATE"`
	CurrencyCode       string      `json:"CURRENCY_CODE" mapstructure:"CURRENCY_CODE"`
	Note               string      `json:"NOTE" mapstructure:"NOTE"`
	Comments           []Comment   `json:"COMMENTS" mapstructure:"COMMENTS"`
	VatItems           []VatItem   `json:"VAT_ITEMS" mapstructure:"VAT_ITEMS"`
	Items              []Item      `json:"ITEMS" mapstructure:"ITEMS"`
	CustomerID         string      `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
	ProjectID          string      `json:"PROJECT_ID" mapstructure:"PROJECT_ID"`
}

// Comment fastbill api expense comment
type Comment struct {
	Date          string `json:"DATE,omitempty" mapstructure:"DATE"`
	Comment       string `json:"COMMENT,omitempty" mapstructure:"COMMENT"`
	CommentPublic string `json:"COMMENT_PUBLIC,omitempty" mapstructure:"COMMENT_PUBLIC"`
}

// VatItem fastbill api expense vat item
type VatItem struct {
	VatPercent  int     `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	CompleteNet int     `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	VatValue    float64 `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
}

// Item fastbill api expense item
type Item struct {
	InvoiceItemID int           `json:"INVOICE_ITEM_ID,omitempty" mapstructure:"INVOICE_ITEM_ID"`
	ArticleNumber string        `json:"ARTICLE_NUMBER,omitempty" mapstructure:"ARTICLE_NUMBER"`
	Description   string        `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Quantity      int           `json:"QUANTITY,omitempty" mapstructure:"QUANTITY"`
	UnitPrice     int           `json:"UNIT_PRICE,omitempty" mapstructure:"UNIT_PRICE"`
	VatPercent    int           `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	VatValue      float64       `json:"VAT_VALUE,omitempty" mapstructure:"VAT_VALUE"`
	CompleteNet   int           `json:"COMPLETE_NET,omitempty" mapstructure:"COMPLETE_NET"`
	CompleteGross float64       `json:"COMPLETE_GROSS,omitempty" mapstructure:"COMPLETE_GROSS"`
	Category      []interface{} `json:"CATEGORY,omitempty" mapstructure:"CATEGORY"`
	SortOrder     int           `json:"SORT_ORDER,omitempty" mapstructure:"SORT_ORDER"`
}

// Request fastbill create expense request
type Request struct {
	InvoiceDate        string  `json:"INVOICE_DATE,omitempty" mapstructure:"INVOICE_DATE"`
	ServicePeriodStart string  `json:"SERVICE_PERIOD_START,omitempty" mapstructure:"SERVICE_PERIOD_START"`
	ServicePeriodEnd   string  `json:"SERVICE_PERIOD_END,omitempty" mapstructure:"SERVICE_PERIOD_END"`
	DueDate            string  `json:"DUE_DATE,omitempty" mapstructure:"DUE_DATE"`
	ProjectID          string  `json:"PROJECT_ID,omitempty" mapstructure:"PROJECT_ID"`
	CustomerID         string  `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	Organization       string  `json:"ORGANIZATION,omitempty" mapstructure:"ORGANIZATION"`
	InvoiceNumber      string  `json:"INVOICE_NUMBER,omitempty" mapstructure:"INVOICE_NUMBER"`
	Comment            string  `json:"COMMENT,omitempty" mapstructure:"COMMENT"`
	SubTotal           int     `json:"SUB_TOTAL,omitempty" mapstructure:"SUB_TOTAL"`
	VatTotal           float64 `json:"VAT_TOTAL,omitempty" mapstructure:"VAT_TOTAL"`
	Items              []Item  `json:"ITEMS,omitempty" mapstructure:"ITEMS"`
}

type getResponse struct {
	Expenses []Expense `json:"EXPENSES" mapstructure:"EXPENSES"`
}

// CreateResponse fastbill create response
type CreateResponse struct {
	InvoiceID int    `json:"INVOICE_ID" mapstructure:"INVOICE_ID"`
	Status    string `json:"STATUS" mapstructure:"STATUS"`
}
