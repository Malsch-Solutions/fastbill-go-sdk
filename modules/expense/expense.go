// Package expense wraps the expense.* services: incoming invoices (receipts)
// booked as expenses, optionally with the receipt file attached.
package expense

import (
	"context"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the expense services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns an expense client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists expenses. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Expense, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("expense.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Expenses, nil
}

// Create books an expense and uploads file (the receipt) as fileName. The
// file is optional: with a nil file only the data is sent.
func (c *Client) Create(ctx context.Context, req *Request, file io.Reader, fileName string) (CreateResponse, error) {
	var res CreateResponse
	data := fastbill.DataRequest("expense.create", req)
	var err error
	if file == nil {
		err = c.r.Do(ctx, data, &res)
	} else {
		err = c.r.DoMultipart(ctx, data, file, fileName, &res)
	}
	if err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("expense.create")
}
