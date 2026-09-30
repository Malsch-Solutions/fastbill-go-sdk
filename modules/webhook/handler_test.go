package webhook_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/webhook"
)

// The invoice.created sample from the FastBill docs, shortened: lower-case
// keys, IDs as strings, amounts as numbers or strings.
const invoiceCreated = `{
	"id": 11,
	"type": "invoice.created",
	"created": "2016-03-02 17:37:13",
	"invoice": {
		"invoice_id": "14566", "type": "outgoing", "customer_id": "2", "customer_number": "2",
		"customer_costcenter_id": "0", "project_id": "0", "currency_code": "EUR",
		"delivery_date": "Leistungszeitraum entspricht Rechnungsdatum.", "invoice_title": "One more time",
		"cash_discount_percent": "0.00", "cash_discount_days": "0",
		"sub_total": 80, "vat_total": 15.2,
		"vat_items": [{"vat_percent": "19.00", "complete_net": 80, "vat_value": 15.2}],
		"items": [{"invoice_item_id": "18", "article_number": "", "description": "Produkt", "quantity": "1.0000",
			"unit_price": "80.00000000", "vat_percent": "19.00", "vat_value": 15.2, "complete_net": 80,
			"complete_gross": 95.2, "sort_order": 1}],
		"total": 95.2, "first_name": "John", "last_name": "Mustemann", "city": "Frankfurt",
		"payment_type": "1", "country_code": "DE", "template_id": "1", "invoice_number": "7",
		"paid_date": "0000-00-00 00:00:00", "is_canceled": "0", "invoice_date": "2016-03-02",
		"due_date": "2016-03-16 00:00:00", "payment_info": "", "lastupdate": "2016-03-02 18:37:10",
		"document_url": "https://my.fastbill.com/download/iGNwBuGMEknEXll"
	},
	"customer": {
		"customer_id": "2", "customer_number": "2", "days_for_payment": "14", "created": "2016-03-02 15:26:57",
		"payment_type": "1", "show_payment_notice": "1", "customer_type": "consumer", "top": "0",
		"newsletter_optin": "1", "first_name": "John", "last_name": "Mustermann", "city": "Frankfurt",
		"country_code": "DE", "currency_code": "EUR", "lastupdate": "2016-03-02 15:58:04", "tags": ""
	}
}`

func newRequest(body string) *http.Request {
	r := httptest.NewRequest(http.MethodPost, "/fastbill", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "FastBill")
	return r
}

func TestParseEvent(t *testing.T) {
	event, err := webhook.ParseEvent(newRequest(invoiceCreated))
	if err != nil {
		t.Fatal(err)
	}
	if event.ID != "11" || event.Type != webhook.InvoiceCreated || event.Created != "2016-03-02 17:37:13" {
		t.Errorf("event = %+v", event)
	}
	if event.Contact != nil || event.Estimate != nil {
		t.Errorf("contact = %+v, estimate = %+v, want nil", event.Contact, event.Estimate)
	}
	inv := event.Invoice
	if inv == nil {
		t.Fatal("invoice = nil")
	}
	if inv.InvoiceID != "14566" || inv.InvoiceNumber != "7" || inv.Total != "95.2" || inv.IsCanceled.Bool() || inv.DocumentURL == "" {
		t.Errorf("invoice = %+v", inv)
	}
	if len(inv.Items) != 1 || inv.Items[0].InvoiceItemID != "18" || inv.Items[0].Quantity != "1.0000" || inv.VatItems[0].VatValue != "15.2" {
		t.Errorf("items = %+v, vat items = %+v", inv.Items, inv.VatItems)
	}
	cust := event.Customer
	if cust == nil {
		t.Fatal("customer = nil")
	}
	if cust.CustomerID != "2" || cust.CustomerType != "consumer" || cust.LastName != "Mustermann" {
		t.Errorf("customer = %+v", cust)
	}
}

func TestParseEventCharset(t *testing.T) {
	r := newRequest(`{"id":"12","type":"contact.deleted","created":"2026-09-30 12:00:00","contact":{"contact_id":"5","customer_id":"2"}}`)
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	event, err := webhook.ParseEvent(r)
	if err != nil {
		t.Fatal(err)
	}
	if event.ID != "12" || event.Type != webhook.ContactDeleted || event.Contact == nil || event.Contact.ContactID != "5" {
		t.Errorf("event = %+v", event)
	}
}

func TestParseEventRejects(t *testing.T) {
	for name, tc := range map[string]struct {
		modify func(*http.Request)
		body   string
		want   string
	}{
		"method": {
			modify: func(r *http.Request) { r.Method = http.MethodPut },
			want:   "method PUT",
		},
		"content type": {
			modify: func(r *http.Request) { r.Header.Set("Content-Type", "application/xml") },
			want:   `Content-Type "application/xml"`,
		},
		"missing content type": {
			modify: func(r *http.Request) { r.Header.Del("Content-Type") },
			want:   `Content-Type ""`,
		},
		"user agent": {
			modify: func(r *http.Request) { r.Header.Set("User-Agent", "curl/8.0") },
			want:   `User-Agent "curl/8.0"`,
		},
		"invalid JSON": {
			body: `{"id": 11, "type": `,
			want: "decode event",
		},
		"wrong field type": {
			body: `{"id": {"x": 1}}`,
			want: "decode event",
		},
		"oversize body": {
			body: `{"id": 11, "type": "invoice.created", "x": "` + strings.Repeat("a", webhook.MaxEventSize) + `"}`,
			want: "larger than",
		},
	} {
		t.Run(name, func(t *testing.T) {
			body := tc.body
			if body == "" {
				body = invoiceCreated
			}
			r := newRequest(body)
			if tc.modify != nil {
				tc.modify(r)
			}
			_, err := webhook.ParseEvent(r)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("err = %v, want it to contain %q", err, tc.want)
			}
		})
	}
}
