package item

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Filter narrows item.get. InvoiceID is required.
type Filter struct {
	InvoiceID fastbill.ID `json:"INVOICE_ID,omitempty"`
}

// Item is an invoice line as item.get returns it.
type Item struct {
	InvoiceItemID fastbill.ID     `json:"INVOICE_ITEM_ID"`
	InvoiceID     fastbill.ID     `json:"INVOICE_ID"`
	CustomerID    fastbill.ID     `json:"CUSTOMER_ID"`
	ArticleNumber string          `json:"ARTICLE_NUMBER"`
	Description   string          `json:"DESCRIPTION"`
	Quantity      fastbill.Number `json:"QUANTITY"`
	UnitPrice     fastbill.Number `json:"UNIT_PRICE"`
	VatPercent    fastbill.Number `json:"VAT_PERCENT"`
	VatValue      fastbill.Number `json:"VAT_VALUE"`
	CompleteNet   fastbill.Number `json:"COMPLETE_NET"`
	CompleteGross fastbill.Number `json:"COMPLETE_GROSS"`
	CurrencyCode  string          `json:"CURRENCY_CODE"`
	SortOrder     fastbill.Number `json:"SORT_ORDER"`
}

type idRequest struct {
	InvoiceItemID fastbill.ID `json:"INVOICE_ITEM_ID"`
}

type getResponse struct {
	Items fastbill.List[Item] `json:"ITEMS"`
}
