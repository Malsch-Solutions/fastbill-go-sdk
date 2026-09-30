// Package customer wraps the customer.* services: the customers invoices,
// estimates and projects are made out to. Contacts of a customer are in
// package contact.
package customer

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the customer services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a customer client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists customers. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Customer, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("customer.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Customers, nil
}

// Create creates a customer. CustomerType is required, and Organization
// (business) or LastName (consumer) depending on it.
func (c *Client) Create(ctx context.Context, customer *Customer) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("customer.create", customer), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("customer.create")
}

// Update changes the customer with customer.CustomerID. Only the fields
// that are set are sent.
func (c *Client) Update(ctx context.Context, customer *Customer) error {
	return c.status(ctx, "customer.update", customer)
}

// Delete deletes a customer.
func (c *Client) Delete(ctx context.Context, customerID fastbill.ID) error {
	return c.status(ctx, "customer.delete", idRequest{CustomerID: customerID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
