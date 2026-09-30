// Package contact wraps the contact.* services: the contact persons of a
// customer.
package contact

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the contact services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a contact client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists contacts. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Contact, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("contact.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Contacts, nil
}

// Create adds a contact to the customer contact.CustomerID.
func (c *Client) Create(ctx context.Context, contact *Contact) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("contact.create", contact), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("contact.create")
}

// Update changes the contact with contact.ContactID of the customer
// contact.CustomerID. Only the fields that are set are sent.
func (c *Client) Update(ctx context.Context, contact *Contact) error {
	return c.status(ctx, "contact.update", contact)
}

// Delete deletes a contact of a customer.
func (c *Client) Delete(ctx context.Context, contactID, customerID fastbill.ID) error {
	return c.status(ctx, "contact.delete", idRequest{ContactID: contactID, CustomerID: customerID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
