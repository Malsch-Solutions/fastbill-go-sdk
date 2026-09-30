package item_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/item"
)

var ctx = context.Background()

// Items as FastBill sends them: IDs and text as strings, amounts as numbers.
const itemsJSON = `[
	{"INVOICE_ITEM_ID": "7", "INVOICE_ID": "3141", "CUSTOMER_ID": "59", "ARTICLE_NUMBER": "SUP-1", "DESCRIPTION": "Support",
	 "QUANTITY": 1.5, "UNIT_PRICE": 823.04, "VAT_PERCENT": 19, "VAT_VALUE": 234.57, "COMPLETE_NET": 1234.56, "COMPLETE_GROSS": 1469.13,
	 "CURRENCY_CODE": "EUR", "SORT_ORDER": "1"},
	{"INVOICE_ITEM_ID": 8, "INVOICE_ID": 3141, "DESCRIPTION": "Travel", "QUANTITY": "1", "UNIT_PRICE": "0", "SORT_ORDER": 2}
]`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ITEMS":` + itemsJSON + `}` })
	items, err := item.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100}, &item.Filter{InvoiceID: "3141"})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("items = %+v", items)
	}
	if it := items[0]; it.InvoiceItemID != "7" || it.Quantity != "1.5" || it.UnitPrice != "823.04" || it.VatPercent != "19" ||
		it.CompleteGross != "1469.13" || it.SortOrder != "1" || it.CurrencyCode != "EUR" {
		t.Errorf("item = %+v", it)
	}
	if it := items[1]; it.InvoiceItemID != "8" || it.InvoiceID != "3141" || it.SortOrder != "2" {
		t.Errorf("item = %+v", it)
	}
	got := server.Last(t)
	if got.Service != "item.get" || got.Limit != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"INVOICE_ID":"3141"}`)
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ITEMS":{}}` })
	items, err := item.NewClient(server.Client()).Get(ctx, fastbill.Page{}, &item.Filter{InvoiceID: "1"})
	if err != nil || len(items) != 0 {
		t.Errorf("Get = %v, %v", items, err)
	}
}

func TestDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	if err := item.NewClient(server.Client()).Delete(ctx, "7"); err != nil {
		t.Fatal(err)
	}
	got := server.Last(t)
	if got.Service != "item.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_ITEM_ID":"7"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "item.delete" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["INVOICE_ID is missing"]}`
	})
	c := item.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "INVOICE_ID is missing" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Delete(ctx, "7"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
