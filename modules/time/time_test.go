package time_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/time"
)

var ctx = context.Background()

// A time entry as FastBill sends it: IDs and text as strings, minutes as
// numbers or strings, no invoice as "0".
const timeJSON = `{
	"TIME_ID": "4410", "TASK_ID": "7", "CUSTOMER_ID": "59", "PROJECT_ID": "21", "INVOICE_ID": "0",
	"DATE": "2026-09-29", "START_TIME": "2026-09-29 09:00:00", "END_TIME": "2026-09-29 10:30:00",
	"MINUTES": 90, "BILLABLE_MINUTES": "75", "COMMENT": "Review"
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"TIMES":[` + timeJSON + `]}` })
	entries, err := time.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100}, &time.Filter{ProjectID: "21", StartDate: "2026-09-01", EndDate: "2026-09-30"})
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("entries = %+v", entries)
	}
	e := entries[0]
	if e.TimeID != "4410" || e.TaskID != "7" || e.InvoiceID != "0" || e.StartTime != "2026-09-29 09:00:00" || e.Minutes != "90" || e.BillableMinutes != "75" || e.Comment != "Review" {
		t.Errorf("entry = %+v", e)
	}
	got := server.Last(t)
	if got.Service != "time.get" || got.Limit != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"PROJECT_ID":"21","START_DATE":"2026-09-01","END_DATE":"2026-09-30"}`)
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"TIMES":{}}` })
	entries, err := time.NewClient(server.Client()).Get(ctx, fastbill.Page{}, nil)
	if err != nil || len(entries) != 0 {
		t.Fatalf("Get = %+v, %v", entries, err)
	}
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","TIME_ID":4411}` })
	res, err := time.NewClient(server.Client()).Create(ctx, &time.Time{
		CustomerID: "59", ProjectID: "21", TaskID: "7",
		StartTime: "2026-09-30 13:00:00", Minutes: fastbill.NewInt(45), Comment: "Call",
	})
	if err != nil || res.TimeID != "4411" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "time.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{
		"CUSTOMER_ID": "59", "PROJECT_ID": "21", "TASK_ID": "7",
		"START_TIME": "2026-09-30 13:00:00", "MINUTES": 45, "COMMENT": "Call"
	}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","TIME_ID":"4411"}` })
	c := time.NewClient(server.Client())

	if err := c.Update(ctx, &time.Time{TimeID: "4411", BillableMinutes: fastbill.NewInt(30)}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "time.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"TIME_ID":"4411","BILLABLE_MINUTES":30}`)

	if err := c.Delete(ctx, "4411"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "time.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"TIME_ID":"4411"}`)
}

// The docs show time.update answering without STATUS.
func TestUpdateWithoutStatus(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"TIME_ID":"4411","COMMENT":"Call"}` })
	if err := time.NewClient(server.Client()).Update(ctx, &time.Time{TimeID: "4411"}); err != nil {
		t.Error(err)
	}
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "time.delete" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Time entry not found"]}`
	})
	c := time.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Time entry not found" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.Create(ctx, &time.Time{}); !errors.As(err, &apiErr) {
		t.Errorf("Create: %v", err)
	}
	if err := c.Delete(ctx, "4411"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
