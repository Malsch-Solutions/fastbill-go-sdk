// Package estimate wraps the estimate.* services: estimates (offers) that
// can be emailed and turned into invoices.
package estimate

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the estimate services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns an estimate client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists estimates. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Estimate, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("estimate.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Estimates, nil
}

// Create creates an estimate.
func (c *Client) Create(ctx context.Context, req *Request) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("estimate.create", req), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("estimate.create")
}

// Delete deletes an estimate.
func (c *Client) Delete(ctx context.Context, estimateID fastbill.ID) error {
	return c.status(ctx, "estimate.delete", idRequest{EstimateID: estimateID})
}

// CreateInvoice writes an invoice from an estimate.
func (c *Client) CreateInvoice(ctx context.Context, estimateID fastbill.ID) (CreateInvoiceResponse, error) {
	var res CreateInvoiceResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("estimate.createinvoice", idRequest{EstimateID: estimateID}), &res); err != nil {
		return res, err
	}
	// FastBill documents only INVOICE_ID here; check STATUS if it is sent.
	if res.Status != "" {
		return res, fastbill.StatusResponse{Status: res.Status}.Err("estimate.createinvoice")
	}
	return res, nil
}

// SendByEmail emails the estimate as PDF.
func (c *Client) SendByEmail(ctx context.Context, req *SendByEmailRequest) error {
	return c.status(ctx, "estimate.sendbyemail", req)
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
