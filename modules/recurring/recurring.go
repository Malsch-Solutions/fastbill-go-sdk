// Package recurring wraps the recurring.* services: recurring invoices that
// FastBill writes on a schedule.
package recurring

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the recurring invoice services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a recurring invoice client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists recurring invoices. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Recurring, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("recurring.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Recurrings, nil
}

// Create creates a recurring invoice.
func (c *Client) Create(ctx context.Context, req *Request) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("recurring.create", req), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("recurring.create")
}

// Update changes a recurring invoice; req.InvoiceID is required.
func (c *Client) Update(ctx context.Context, req *Request) error {
	return c.status(ctx, "recurring.update", req)
}

// Delete deletes a recurring invoice.
func (c *Client) Delete(ctx context.Context, invoiceID fastbill.ID) error {
	return c.status(ctx, "recurring.delete", idRequest{InvoiceID: invoiceID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
