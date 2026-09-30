package contact_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/contact"
)

var ctx = context.Background()

// A contact as FastBill sends it: IDs and text as strings.
const contactJSON = `{
	"CONTACT_ID": "812", "CUSTOMER_ID": "59", "ORGANIZATION": "Beispiel GmbH", "POSITION": "CFO",
	"ACADEMIC_DEGREE": "Dr.", "SALUTATION": "mrs", "FIRST_NAME": "Erika", "LAST_NAME": "Mustermann",
	"ADDRESS": "Hauptstraße 1", "ZIPCODE": "10115", "CITY": "Berlin", "COUNTRY_CODE": "DE",
	"PHONE": "+49 30 123456", "EMAIL": "erika@example.com", "CURRENCY_CODE": "EUR",
	"CREATED": "2026-01-15 10:00:00", "TAGS": "billing"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"CONTACTS":[` + contactJSON + `]}` })
	contacts, err := contact.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 50}, &contact.Filter{CustomerID: "59"})
	if err != nil {
		t.Fatal(err)
	}
	if len(contacts) != 1 {
		t.Fatalf("contacts = %+v", contacts)
	}
	co := contacts[0]
	if co.ContactID != "812" || co.CustomerID != "59" || co.AcademicDegree != "Dr." || co.LastName != "Mustermann" || co.Created != "2026-01-15 10:00:00" {
		t.Errorf("contact = %+v", co)
	}
	got := server.Last(t)
	if got.Service != "contact.get" || got.Limit != 50 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"CUSTOMER_ID":"59"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","CONTACT_ID":813}` })
	res, err := contact.NewClient(server.Client()).Create(ctx, &contact.Contact{
		CustomerID: "59", FirstName: "Max", LastName: "Mustermann", Email: "max@example.com",
	})
	if err != nil || res.ContactID != "813" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "contact.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"CUSTOMER_ID":"59","FIRST_NAME":"Max","LAST_NAME":"Mustermann","EMAIL":"max@example.com"}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","CONTACT_ID":"813"}` })
	c := contact.NewClient(server.Client())

	if err := c.Update(ctx, &contact.Contact{ContactID: "813", CustomerID: "59", Phone: "+49 30 654321"}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "contact.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"CONTACT_ID":"813","CUSTOMER_ID":"59","PHONE":"+49 30 654321"}`)

	if err := c.Delete(ctx, "813", "59"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "contact.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"CONTACT_ID":"813","CUSTOMER_ID":"59"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "contact.update" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Contact not found"]}`
	})
	c := contact.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Contact not found" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Delete(ctx, "813", "59"); !errors.As(err, &apiErr) {
		t.Errorf("Delete: %v", err)
	}
	if err := c.Update(ctx, &contact.Contact{ContactID: "813"}); !errors.As(err, &apiErr) {
		t.Errorf("Update with failed status: %v", err)
	}
}
