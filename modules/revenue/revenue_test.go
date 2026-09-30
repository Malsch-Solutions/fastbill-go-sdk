package revenue_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/revenue"
)

var ctx = context.Background()

// A revenue as FastBill sends it: IDs and text as strings, amounts as
// numbers, an empty list as {}.
const revenueJSON = `{
	"INVOICE_ID": "4501", "TYPE": "outgoing", "CUSTOMER_ID": "59", "CUSTOMER_NUMBER": "1042",
	"PROJECT_ID": "7", "CURRENCY_CODE": "EUR", "BASE_CURRENCY_CODE": "EUR", "EXCHANGE_RATE": 1,
	"INVOICE_NUMBER": "E-17", "INVOICE_DATE": "2026-09-02", "DUE_DATE": "2026-09-16", "PAID_DATE": "2026-09-10",
	"IS_CANCELED": 0, "PAYMENT_TYPE": "1", "CASH_DISCOUNT_DAYS": "0",
	"SUB_TOTAL": 1234.56, "VAT_TOTAL": 234.57, "TOTAL": 1469.13, "ORGANIZATION": "Kunde AG",
	"VAT_ITEMS": [{"VAT_PERCENT": 19, "COMPLETE_NET": 1234.56, "VAT_VALUE": 234.57}],
	"ITEMS": [{"INVOICE_ITEM_ID": "33", "DESCRIPTION": "Beratung", "QUANTITY": 2.5, "UNIT_PRICE": 493.824, "VAT_PERCENT": 19, "CURRENCY_CODE": "EUR"}],
	"PAYMENTS": [{"PAYMENT_ID": "88", "DATE": "2026-09-10", "AMOUNT": "1469.13", "CURRENCY_CODE": "EUR", "TYPE": "1"}],
	"COMMENTS": {},
	"LASTUPDATE": "2026-09-10 12:00:00", "DOCUMENT_URL": "https://my.fastbill.com/download/e.pdf"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"REVENUES":[` + revenueJSON + `]}` })
	revenues, err := revenue.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100}, &revenue.Filter{CustomerID: "59", StartPaidDate: "2026-09-01"})
	if err != nil {
		t.Fatal(err)
	}
	if len(revenues) != 1 {
		t.Fatalf("revenues = %+v", revenues)
	}
	r := revenues[0]
	if r.InvoiceID != "4501" || r.ProjectID != "7" || r.Total != "1469.13" || r.ExchangeRate != "1" ||
		r.IsCanceled != fastbill.No || r.PaymentType != "1" || r.InvoiceNumber != "E-17" {
		t.Errorf("revenue = %+v", r)
	}
	if item := r.Items[0]; item.InvoiceItemID != "33" || item.Quantity != "2.5" || item.UnitPrice != "493.824" {
		t.Errorf("item = %+v", item)
	}
	if len(r.Comments) != 0 || r.Payments[0].Amount != "1469.13" || r.VatItems[0].VatPercent != "19" {
		t.Errorf("comments = %v, payments = %v, vat items = %v", r.Comments, r.Payments, r.VatItems)
	}
	got := server.Last(t)
	if got.Service != "revenue.get" || got.Limit != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"CUSTOMER_ID":"59","START_PAID_DATE":"2026-09-01"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","INVOICE_ID":4502}` })
	res, err := revenue.NewClient(server.Client()).Create(ctx, &revenue.Request{
		InvoiceDate: "2026-09-02",
		CustomerID:  "59",
		SubTotal:    fastbill.NewNumber(1234.56),
		VatTotal:    fastbill.NewNumber(234.57),
	}, strings.NewReader("receipt"), "beleg.pdf")
	if err != nil || res.InvoiceID != "4502" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "revenue.create" || got.FileName != "beleg.pdf" || got.File != "receipt" {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_DATE":"2026-09-02","CUSTOMER_ID":"59","SUB_TOTAL":1234.56,"VAT_TOTAL":234.57}`)
}

func TestSetPaidAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "revenue.setpaid" {
			return `{"STATUS":"success","INVOICE_NUMBER":"E-17"}`
		}
		return `{"STATUS":"success"}`
	})
	c := revenue.NewClient(server.Client())

	if res, err := c.SetPaid(ctx, &revenue.SetPaidRequest{InvoiceID: "4501", PaidDate: "2026-09-30"}); err != nil || res.InvoiceNumber != "E-17" {
		t.Errorf("SetPaid = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "revenue.setpaid" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_ID":"4501","PAID_DATE":"2026-09-30"}`)

	if err := c.Delete(ctx, "4501"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "revenue.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"INVOICE_ID":"4501"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "revenue.delete" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Revenue not found"]}`
	})
	c := revenue.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Revenue not found" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.SetPaid(ctx, &revenue.SetPaidRequest{InvoiceID: "1"}); !errors.As(err, &apiErr) {
		t.Errorf("SetPaid: %v", err)
	}
	if err := c.Delete(ctx, "1"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
