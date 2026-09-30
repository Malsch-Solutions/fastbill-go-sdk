package recurring_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/recurring"
)

var ctx = context.Background()

// A recurring invoice as FastBill sends it: IDs and text as strings,
// amounts as numbers, an empty list as {}.
const recurringJSON = `{
	"INVOICE_ID": "620", "TYPE": "outgoing", "CUSTOMER_ID": "59", "CUSTOMER_COSTCENTER_ID": "0",
	"CURRENCY_CODE": "EUR", "BASE_CURRENCY_CODE": "EUR", "EXCHANGE_RATE": "1.0000",
	"TEMPLATE_ID": "3", "INTROTEXT": "Monthly support", "IS_CANCELED": "0", "IS_GROSS": 1,
	"FREQUENCY": "monthly", "START_DATE": "2026-10-01", "OCCURENCES": "12", "OUTPUT_TYPE": "draft", "EMAIL_NOTIFY": "1",
	"CASH_DISCOUNT_PERCENT": 2.5, "CASH_DISCOUNT_DAYS": 10,
	"SUB_TOTAL": 1234.56, "VAT_TOTAL": 234.57, "TOTAL": 1469.13,
	"VAT_ITEMS": {},
	"ITEMS": {"0": {"INVOICE_ITEM_ID": 71, "DESCRIPTION": "Support", "QUANTITY": 1, "UNIT_PRICE": 1234.56, "VAT_PERCENT": 19, "SORT_ORDER": "1"}}
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"INVOICES":[` + recurringJSON + `]}` })
	list, err := recurring.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 10}, &recurring.Filter{InvoiceID: "620"})
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("recurring = %+v", list)
	}
	r := list[0]
	if r.InvoiceID != "620" || r.Frequency != recurring.FrequencyMonthly || r.Occurrences != "12" || !r.EmailNotify.Bool() ||
		!r.IsGross.Bool() || r.IsCanceled.Bool() || r.CashDiscountPercent != "2.5" || r.Total != "1469.13" || r.ExchangeRate != "1.0000" {
		t.Errorf("recurring = %+v", r)
	}
	if len(r.VatItems) != 0 || len(r.Items) != 1 || r.Items[0].InvoiceItemID != "71" || r.Items[0].UnitPrice != "1234.56" {
		t.Errorf("vat items = %v, items = %+v", r.VatItems, r.Items)
	}
	got := server.Last(t)
	if got.Service != "recurring.get" || got.Limit != 10 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"INVOICE_ID":"620"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","INVOICE_ID":621}` })
	res, err := recurring.NewClient(server.Client()).Create(ctx, &recurring.Request{
		CustomerID:  "59",
		StartDate:   "2026-10-01",
		Frequency:   recurring.FrequencyMonthly,
		Occurrences: fastbill.NewInt(0),
		OutputType:  "draft",
		EmailNotify: fastbill.Yes,
		Items:       []recurring.Item{{Description: "Support", Quantity: fastbill.NewInt(1), UnitPrice: fastbill.NewNumber(99.9), VatPercent: fastbill.NewInt(19)}},
	})
	if err != nil || res.InvoiceID != "621" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "recurring.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{
		"CUSTOMER_ID": "59", "START_DATE": "2026-10-01", "FREQUENCY": "monthly", "OCCURENCES": 0,
		"OUTPUT_TYPE": "draft", "EMAIL_NOTIFY": "1",
		"ITEMS": [{"DESCRIPTION": "Support", "QUANTITY": 1, "UNIT_PRICE": 99.9, "VAT_PERCENT": 19}]
	}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	c := recurring.NewClient(server.Client())

	if err := c.Update(ctx, &recurring.Request{InvoiceID: "620", DeleteExistingItems: fastbill.Yes, Frequency: recurring.Frequency3Months}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "recurring.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_ID":"620","DELETE_EXISTING_ITEMS":"1","FREQUENCY":"3 months"}`)

	if err := c.Delete(ctx, "620"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "recurring.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_ID":"620"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "recurring.update" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Invoice not found"]}`
	})
	c := recurring.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Invoice not found" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Update(ctx, &recurring.Request{InvoiceID: "1"}); !errors.As(err, &apiErr) {
		t.Errorf("Update with failed status: %v", err)
	}
	if err := c.Delete(ctx, "1"); !errors.As(err, &apiErr) {
		t.Errorf("Delete: %v", err)
	}
}
