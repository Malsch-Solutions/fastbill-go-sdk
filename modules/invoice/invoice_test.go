package invoice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/invoice"
)

var ctx = context.Background()

// An invoice as FastBill sends it: IDs and text as strings, amounts as
// numbers, an empty list as {}.
const invoiceJSON = `{
	"INVOICE_ID": "3141", "TYPE": "outgoing", "CUSTOMER_ID": "59", "CUSTOMER_NUMBER": "1042",
	"INVOICE_NUMBER": "00277", "INVOICE_DATE": "2026-09-01", "IS_CANCELED": "0", "IS_GROSS": 0,
	"PAYMENT_TYPE": "1", "SUB_TOTAL": 1234.5, "VAT_TOTAL": 0, "TOTAL": 1234.5, "CURRENCY_CODE": "EUR",
	"ITEMS": [{"INVOICE_ITEM_ID": 7, "DESCRIPTION": "Support", "QUANTITY": 1.5, "UNIT": "Std.", "UNIT_PRICE": 823, "VAT_PERCENT": 0, "COMPLETE_NET": 1234.5, "CATEGORY": []}],
	"VAT_ITEMS": [{"VAT_PERCENT": 0, "COMPLETE_NET": 1234.5, "VAT_VALUE": 0}],
	"PAYMENTS": {},
	"DOCUMENT_URL": "https://my.fastbill.com/download/x.pdf"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"INVOICES":[` + invoiceJSON + `]}` })
	invoices, err := invoice.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100}, &invoice.Filter{Type: invoice.TypeDraft})
	if err != nil {
		t.Fatal(err)
	}
	if len(invoices) != 1 {
		t.Fatalf("invoices = %+v", invoices)
	}
	inv := invoices[0]
	if inv.InvoiceID != "3141" || inv.InvoiceNumber != "00277" || inv.Total != "1234.5" || inv.IsCanceled.Bool() || inv.IsGross != fastbill.No {
		t.Errorf("invoice = %+v", inv)
	}
	if item := inv.Items[0]; item.InvoiceItemID != "7" || item.Quantity != "1.5" || item.Unit != "Std." {
		t.Errorf("item = %+v", item)
	}
	if len(inv.Payments) != 0 || inv.VatItems[0].CompleteNet != "1234.5" {
		t.Errorf("payments = %v, vat items = %v", inv.Payments, inv.VatItems)
	}
	got := server.Last(t)
	if got.Service != "invoice.get" || got.Limit != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"TYPE":"draft"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","INVOICE_ID":3142}` })
	res, err := invoice.NewClient(server.Client()).Create(ctx, &invoice.Request{
		CustomerID: "59",
		VatCase:    "small_business_regulation",
		Items: []invoice.Item{{
			Description: "Support", Quantity: fastbill.NewNumber(1.5), UnitPrice: fastbill.NewInt(80), VatPercent: fastbill.NewInt(0),
		}},
	})
	if err != nil || res.InvoiceID != "3142" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	fastbilltest.JSONEqual(t, server.Last(t).Data, `{
		"CUSTOMER_ID": "59", "VAT_CASE": "small_business_regulation",
		"ITEMS": [{"DESCRIPTION": "Support", "QUANTITY": 1.5, "UNIT_PRICE": 80, "VAT_PERCENT": 0}]
	}`)
}

func TestServicesWithStatus(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		switch r.Service {
		case "invoice.complete":
			return `{"STATUS":"success","INVOICE_NUMBER":"00278"}`
		case "invoice.sendbypost":
			return `{"STATUS":"success","REMAINING_CREDITS":"12"}`
		case "invoice.setpaid":
			return `{"STATUS":"success","INVOICE_NUMBER":"00278"}`
		}
		return `{"STATUS":"success"}`
	})
	c := invoice.NewClient(server.Client())

	if err := c.Update(ctx, &invoice.Request{InvoiceID: "1", DeleteExistingItems: fastbill.Yes}); err != nil {
		t.Error(err)
	}
	fastbilltest.JSONEqual(t, server.Last(t).Data, `{"INVOICE_ID":"1","DELETE_EXISTING_ITEMS":"1"}`)
	if res, err := c.Complete(ctx, "1"); err != nil || res.InvoiceNumber != "00278" {
		t.Errorf("Complete = %+v, %v", res, err)
	}
	for name, call := range map[string]func() error{
		"invoice.delete": func() error { return c.Delete(ctx, "1") },
		"invoice.cancel": func() error { return c.Cancel(ctx, "1") },
		"invoice.lock":   func() error { return c.Lock(ctx, "1") },
	} {
		if err := call(); err != nil {
			t.Errorf("%s: %v", name, err)
		}
		if got := server.Last(t); got.Service != name {
			t.Errorf("service = %s, want %s", got.Service, name)
		}
		fastbilltest.JSONEqual(t, server.Last(t).Data, `{"INVOICE_ID":"1"}`)
	}
	if err := c.SendByEmail(ctx, &invoice.SendByEmailRequest{InvoiceID: "1", Recipient: invoice.Recipient{To: "a@example.com"}}); err != nil {
		t.Error(err)
	}
	fastbilltest.JSONEqual(t, server.Last(t).Data, `{"INVOICE_ID":"1","RECIPIENT":{"TO":"a@example.com"}}`)
	if res, err := c.SendByPost(ctx, "1"); err != nil || res.RemainingCredits != "12" {
		t.Errorf("SendByPost = %+v, %v", res, err)
	}
	if res, err := c.SetPaid(ctx, &invoice.SetPaidRequest{InvoiceID: "1", PaidDate: "2026-09-30"}); err != nil || res.InvoiceNumber != "00278" {
		t.Errorf("SetPaid = %+v, %v", res, err)
	}
	fastbilltest.JSONEqual(t, server.Last(t).Data, `{"INVOICE_ID":"1","PAID_DATE":"2026-09-30"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "invoice.cancel" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Invoice not found"]}`
	})
	c := invoice.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Invoice not found" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Cancel(ctx, "1"); !errors.As(err, &apiErr) {
		t.Errorf("Cancel with failed status: %v", err)
	}
}
