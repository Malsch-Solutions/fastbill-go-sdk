package customer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/customer"
)

var ctx = context.Background()

// A customer as FastBill sends it: IDs and text as strings, numbers as
// numbers or strings, a flag as a number.
const customerJSON = `{
	"CUSTOMER_ID": "59", "CUSTOMER_NUMBER": "1042", "CUSTOMER_TYPE": "business",
	"DAYS_FOR_PAYMENT": 14, "PAYMENT_TYPE": "1", "SHOW_PAYMENT_NOTICE": 1, "TOP": "0",
	"ACCOUNT_RECEIVABLE": "10042", "ORGANIZATION": "Beispiel GmbH", "FIRST_NAME": "Erika", "LAST_NAME": "Mustermann",
	"ADDRESS": "Hauptstraße 1", "ZIPCODE": "10115", "CITY": "Berlin", "COUNTRY_CODE": "DE",
	"EMAIL": "erika@example.com", "VAT_ID": "DE123456789", "CURRENCY_CODE": "EUR",
	"BUYER_REFERENCE": "991-12345-67", "GLN": "4012345000009",
	"CREATED": "2026-01-15 10:00:00", "LASTUPDATE": "2026-09-01 12:30:00", "TAGS": "retainer",
	"DOCUMENT_HISTORY_URL": "https://my.fastbill.com/history/abc"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"CUSTOMERS":[` + customerJSON + `]}` })
	customers, err := customer.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100, Offset: 100}, &customer.Filter{CountryCode: "DE", Term: "Beispiel"})
	if err != nil {
		t.Fatal(err)
	}
	if len(customers) != 1 {
		t.Fatalf("customers = %+v", customers)
	}
	cu := customers[0]
	if cu.CustomerID != "59" || cu.CustomerNumber != "1042" || cu.CustomerType != customer.TypeBusiness || cu.Organization != "Beispiel GmbH" {
		t.Errorf("customer = %+v", cu)
	}
	if cu.DaysForPayment != "14" || cu.PaymentType != "1" || !cu.ShowPaymentNotice.Bool() || cu.Top != fastbill.No {
		t.Errorf("payment terms = %+v", cu)
	}
	if cu.AccountReceivable != "10042" || cu.BuyerReference != "991-12345-67" || cu.GLN != "4012345000009" || cu.DocumentHistoryURL == "" {
		t.Errorf("new fields = %+v", cu)
	}
	got := server.Last(t)
	if got.Service != "customer.get" || got.Limit != 100 || got.Offset != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"COUNTRY_CODE":"DE","TERM":"Beispiel"}`)
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"CUSTOMERS":{}}` })
	customers, err := customer.NewClient(server.Client()).Get(ctx, fastbill.Page{}, nil)
	if err != nil || len(customers) != 0 {
		t.Fatalf("Get = %+v, %v", customers, err)
	}
	if got := server.Last(t); got.Filter != nil {
		t.Errorf("filter = %s, want none", got.Filter)
	}
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string {
		return `{"STATUS":"success","CUSTOMER_ID":60,"CUSTOMER_NUMBER":"1043"}`
	})
	res, err := customer.NewClient(server.Client()).Create(ctx, &customer.Customer{
		CustomerType:    customer.TypeConsumer,
		LastName:        "Mustermann",
		DaysForPayment:  fastbill.NewInt(14),
		BuyerReference:  "991-12345-67",
		NewsletterOptIn: fastbill.No,
	})
	if err != nil || res.CustomerID != "60" || res.CustomerNumber != "1043" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "customer.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{
		"CUSTOMER_TYPE": "consumer", "LAST_NAME": "Mustermann", "DAYS_FOR_PAYMENT": 14,
		"BUYER_REFERENCE": "991-12345-67", "NEWSLETTER_OPTIN": "0"
	}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","CUSTOMER_ID":"59"}` })
	c := customer.NewClient(server.Client())

	if err := c.Update(ctx, &customer.Customer{CustomerID: "59", Email: "new@example.com", ShowPaymentNotice: fastbill.Yes}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "customer.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"CUSTOMER_ID":"59","EMAIL":"new@example.com","SHOW_PAYMENT_NOTICE":"1"}`)

	if err := c.Delete(ctx, "59"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "customer.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"CUSTOMER_ID":"59"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "customer.delete" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Customer not found"]}`
	})
	c := customer.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Customer not found" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.Create(ctx, &customer.Customer{}); !errors.As(err, &apiErr) {
		t.Errorf("Create: %v", err)
	}
	if err := c.Delete(ctx, "59"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
