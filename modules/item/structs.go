package item

//Filter available  filter
type Filter struct {
	InvoiceID int `json:"INVOICE_ID"` //A specific invoice ID
}

//Item fastbill item definition
type Item struct {
	InvoiceItemID string  `json:"INVOICE_ITEM_ID" mapstructure:"INVOICE_ITEM_ID"`
	InvoiceID     string  `json:"INVOICE_ID" mapstructure:"INVOICE_ID"`
	CustomerID    string  `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
	ArticleNumber string  `json:"ARTICLE_NUMBER" mapstructure:"ARTICLE_NUMBER"`
	Description   string  `json:"DESCRIPTION" mapstructure:"DESCRIPTION"`
	Quantity      string  `json:"QUANTITY" mapstructure:"QUANTITY"`
	UnitPrice     string  `json:"UNIT_PRICE" mapstructure:"UNIT_PRICE"`
	VatPrice      string  `json:"VAT_PERCENT" mapstructure:"VAT_PERCENT"`
	VatValue      int     `json:"VAT_VALUE" mapstructure:"VAT_VALUE"`
	CompleteNet   float64 `json:"COMPLETE_NET" mapstructure:"COMPLETE_NET"`
	CompleteGross float64 `json:"COMPLETE_GROSS" mapstructure:"COMPLETE_GROSS"`
	CurrencyCode  string  `json:"CURRENCY_CODE" mapstructure:"CURRENCY_CODE"`
	SortOrder     int     `json:"SORT_ORDER" mapstructure:"SORT_ORDER"`
}

type getResponse struct {
	Items []Item `json:"ITEMS"`
}

type deleteRequest struct {
	InvoiceItemID string `json:"INVOICE_ITEM_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
