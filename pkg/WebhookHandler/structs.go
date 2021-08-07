package webhook_handler

import (
	"github.com/malsch-solutions/fastbill-go-sdk/modules/contact"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/estimate"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/invoice"
)

type WebhookData struct {
	ID       int                `json:"id"`
	Type     string             `json:"type"`
	Customer *customer.Customer `json:"customer"`
	Contact  *contact.Contact   `json:"contact"`
	Invoice  invoice.Invoice    `json:"invoice"`
	Estimate estimate.Estimate  `json:"estimate"`
	Created  string             `json:"created"`
}
