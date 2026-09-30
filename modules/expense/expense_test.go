package expense_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/expense"
)

var ctx = context.Background()

// An expense as FastBill sends it: IDs and text as strings, amounts as
// numbers, an empty list as {}.
const expenseJSON = `{
	"INVOICE_ID": "812", "ORGANIZATION": "Hosting GmbH", "INVOICE_NUMBER": "R-2026-0917",
	"INVOICE_DATE": "2026-09-17", "DUE_DATE": "2026-10-01", "PROJECT_ID": "0", "CUSTOMER_ID": "",
	"SUB_TOTAL": 1234.56, "VAT_TOTAL": 234.57, "TOTAL": 1469.13, "NOTE": "", "COMMENT": "Server",
	"CURRENCY_CODE": "USD", "BASE_CURRENCY_CODE": "EUR", "EXCHANGE_RATE": "1.1712",
	"CATEGORY": "Hosting", "PAYMENT_INFO": "SEPA", "PAID_DATE": "2026-09-20",
	"COMMENTS": {},
	"VAT_ITEMS": [{"VAT_PERCENT": 19, "COMPLETE_NET": 1234.56, "VAT_VALUE": 234.57}],
	"ITEMS": [{"INVOICE_ITEM_ID": 91, "DESCRIPTION": "Server", "QUANTITY": 1, "UNIT_PRICE": 1234.56, "VAT_PERCENT": "19.00", "CATEGORY": [], "SORT_ORDER": 1}],
	"DOCUMENT_URL": "https://my.fastbill.com/download/r.pdf"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"EXPENSES":[` + expenseJSON + `]}` })
	expenses, err := expense.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 50, Offset: 50}, &expense.Filter{StartDate: "2026-09-01", EndDate: "2026-09-30"})
	if err != nil {
		t.Fatal(err)
	}
	if len(expenses) != 1 {
		t.Fatalf("expenses = %+v", expenses)
	}
	e := expenses[0]
	if e.InvoiceID != "812" || e.Organization != "Hosting GmbH" || e.SubTotal != "1234.56" || e.Total != "1469.13" ||
		e.ExchangeRate != "1.1712" || e.BaseCurrencyCode != "EUR" || e.PaymentInfo != "SEPA" || e.Comment != "Server" {
		t.Errorf("expense = %+v", e)
	}
	if string(e.Category) != `"Hosting"` || len(e.Comments) != 0 || e.VatItems[0].VatValue != "234.57" {
		t.Errorf("category = %s, comments = %v, vat items = %v", e.Category, e.Comments, e.VatItems)
	}
	if item := e.Items[0]; item.InvoiceItemID != "91" || item.UnitPrice != "1234.56" || item.VatPercent != "19.00" || string(item.Category) != "[]" {
		t.Errorf("item = %+v", item)
	}
	got := server.Last(t)
	if got.Service != "expense.get" || got.Limit != 50 || got.Offset != 50 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"START_DATE":"2026-09-01","END_DATE":"2026-09-30"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","INVOICE_ID":813}` })
	c := expense.NewClient(server.Client())
	req := &expense.Request{
		InvoiceDate:  "2026-09-17",
		Organization: "Hosting GmbH",
		SubTotal:     fastbill.NewNumber(1234.56),
		VatTotal:     fastbill.NewNumber(234.57),
		Items:        []expense.Item{{Description: "Server", Quantity: fastbill.NewInt(1), UnitPrice: fastbill.NewNumber(1234.56)}},
	}
	res, err := c.Create(ctx, req, strings.NewReader("%PDF-1.7"), "receipt.pdf")
	if err != nil || res.InvoiceID != "813" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "expense.create" || got.FileName != "receipt.pdf" || got.File != "%PDF-1.7" {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Data, `{
		"INVOICE_DATE": "2026-09-17", "ORGANIZATION": "Hosting GmbH", "SUB_TOTAL": 1234.56, "VAT_TOTAL": 234.57,
		"ITEMS": [{"DESCRIPTION": "Server", "QUANTITY": 1, "UNIT_PRICE": 1234.56}]
	}`)

	// Without a file the data is sent as plain JSON.
	if _, err := c.Create(ctx, req, nil, ""); err != nil {
		t.Fatal(err)
	}
	if got := server.Last(t); got.FileName != "" || !strings.HasPrefix(got.Header.Get("Content-Type"), "application/json") {
		t.Errorf("request without file = %+v", got)
	}
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "expense.create" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Invalid filter"]}`
	})
	c := expense.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Invalid filter" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.Create(ctx, &expense.Request{}, strings.NewReader("x"), "x.pdf"); !errors.As(err, &apiErr) {
		t.Errorf("Create with failed status: %v", err)
	}
}
