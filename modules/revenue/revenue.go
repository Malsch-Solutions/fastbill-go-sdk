// Package revenue wraps the revenue.* services: income booked without an
// invoice written in FastBill, optionally with a document attached.
package revenue

import (
	"context"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the revenue services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a revenue client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists revenues. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Revenue, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("revenue.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Revenues, nil
}

// Create books a revenue and uploads file as fileName. The file is
// optional: with a nil file only the data is sent.
func (c *Client) Create(ctx context.Context, req *Request, file io.Reader, fileName string) (CreateResponse, error) {
	var res CreateResponse
	data := fastbill.DataRequest("revenue.create", req)
	var err error
	if file == nil {
		err = c.r.Do(ctx, data, &res)
	} else {
		err = c.r.DoMultipart(ctx, data, file, fileName, &res)
	}
	if err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("revenue.create")
}

// SetPaid marks a revenue as paid.
func (c *Client) SetPaid(ctx context.Context, req *SetPaidRequest) (SetPaidResponse, error) {
	var res SetPaidResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("revenue.setpaid", req), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("revenue.setpaid")
}

// Delete deletes a revenue.
func (c *Client) Delete(ctx context.Context, invoiceID fastbill.ID) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("revenue.delete", idRequest{InvoiceID: invoiceID}), &res); err != nil {
		return err
	}
	return res.Err("revenue.delete")
}
