package webhook

import (
	"strings"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/contact"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/estimate"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/invoice"
)

// TypeURL is the endpoint type (TYPE) for an HTTP endpoint, the only one
// FastBill currently offers.
const TypeURL = "url"

// EventType names the event that triggers a notification.
type EventType string

// Event types a webhook can subscribe to.
const (
	CustomerCreated  EventType = "customer.created"
	CustomerUpdated  EventType = "customer.updated"
	CustomerDeleted  EventType = "customer.deleted"
	InvoiceCreated   EventType = "invoice.created"
	InvoiceCompleted EventType = "invoice.completed"
	InvoiceCanceled  EventType = "invoice.canceled"
	EstimateCreated  EventType = "estimate.created"
	EstimateUpdated  EventType = "estimate.updated"
	ContactCreated   EventType = "contact.created"
	ContactUpdated   EventType = "contact.updated"
	ContactDeleted   EventType = "contact.deleted"
)

// JoinEvents formats events as the comma-separated EVENTS value.
func JoinEvents(events ...EventType) string {
	names := make([]string, len(events))
	for i, e := range events {
		names[i] = string(e)
	}
	return strings.Join(names, ",")
}

// Webhook is a registered webhook as webhook.get returns it.
type Webhook struct {
	WebhookID fastbill.ID `json:"WEBHOOK_ID"`
	Endpoint  string      `json:"ENDPOINT"`
	Type      string      `json:"TYPE"`
	// Events is a comma-separated list of event types.
	Events string `json:"EVENTS"`
}

// Request is the data of webhook.create.
type Request struct {
	// Type is TypeURL.
	Type     string `json:"TYPE"`
	Endpoint string `json:"ENDPOINT"`
	// Events is a comma-separated list of event types; see JoinEvents.
	Events string `json:"EVENTS"`
}

// CreateResponse is the answer of webhook.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	WebhookID fastbill.ID `json:"WEBHOOK_ID"`
}

// Event is a notification FastBill sends to a webhook endpoint. It carries
// the objects the event is about; the others are nil. FastBill sends the
// object fields in lower case, which decode into the same structs as the
// API responses.
type Event struct {
	// ID is the notification ID.
	ID   fastbill.ID `json:"id"`
	Type EventType   `json:"type"`
	// Created is the time of the notification, YYYY-MM-DD hh:mm:ss.
	Created  string             `json:"created"`
	Customer *customer.Customer `json:"customer,omitempty"`
	Contact  *contact.Contact   `json:"contact,omitempty"`
	Invoice  *invoice.Invoice   `json:"invoice,omitempty"`
	Estimate *estimate.Estimate `json:"estimate,omitempty"`
}

type idRequest struct {
	WebhookID fastbill.ID `json:"WEBHOOK_ID"`
}

type getResponse struct {
	Webhooks fastbill.List[Webhook] `json:"WEBHOOKS"`
}
