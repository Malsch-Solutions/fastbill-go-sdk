package webhook_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/webhook"
)

var ctx = context.Background()

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string {
		return `{"WEBHOOKS":[
			{"WEBHOOK_ID": "15", "ENDPOINT": "https://example.com/hook", "TYPE": "url", "EVENTS": "customer.created,customer.updated"},
			{"WEBHOOK_ID": "16", "ENDPOINT": "https://example.com/hook2", "TYPE": "url", "EVENTS": "invoice.created"}
		]}`
	})
	webhooks, err := webhook.NewClient(server.Client()).Get(ctx, fastbill.Page{})
	if err != nil {
		t.Fatal(err)
	}
	if len(webhooks) != 2 {
		t.Fatalf("webhooks = %+v", webhooks)
	}
	if w := webhooks[0]; w.WebhookID != "15" || w.Endpoint != "https://example.com/hook" || w.Type != webhook.TypeURL || w.Events != "customer.created,customer.updated" {
		t.Errorf("webhook = %+v", w)
	}
	if got := server.Last(t); got.Service != "webhook.get" || got.Filter != nil {
		t.Errorf("request = %+v", got)
	}
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"WEBHOOKS":{}}` })
	webhooks, err := webhook.NewClient(server.Client()).Get(ctx, fastbill.Page{})
	if err != nil || len(webhooks) != 0 {
		t.Errorf("Get = %+v, %v", webhooks, err)
	}
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","WEBHOOK_ID":15}` })
	res, err := webhook.NewClient(server.Client()).Create(ctx, &webhook.Request{
		Type:     webhook.TypeURL,
		Endpoint: "https://example.com/hook",
		Events:   webhook.JoinEvents(webhook.CustomerCreated, webhook.InvoiceCompleted),
	})
	if err != nil || res.WebhookID != "15" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "webhook.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"TYPE":"url","ENDPOINT":"https://example.com/hook","EVENTS":"customer.created,invoice.completed"}`)
}

func TestDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	if err := webhook.NewClient(server.Client()).Delete(ctx, "15"); err != nil {
		t.Fatal(err)
	}
	got := server.Last(t)
	if got.Service != "webhook.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"WEBHOOK_ID":"15"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		switch r.Service {
		case "webhook.delete":
			return `{"STATUS":"failed"}`
		case "webhook.create":
			return `{"ERRORS":["Invalid endpoint"]}`
		}
		return `{"ERRORS":["Unauthorized"]}`
	})
	c := webhook.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Unauthorized" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.Create(ctx, &webhook.Request{Type: webhook.TypeURL}); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Invalid endpoint" {
		t.Errorf("Create: %v", err)
	}
	if err := c.Delete(ctx, "15"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
