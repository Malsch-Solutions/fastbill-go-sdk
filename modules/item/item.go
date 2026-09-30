// Package item wraps the item.* services: the line items of an invoice.
// Items are created with the invoice (invoice.create, ITEMS).
package item

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the item services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns an item client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists the items of an invoice. FastBill requires filter.InvoiceID.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Item, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("item.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Items, nil
}

// Delete deletes an item of an invoice.
func (c *Client) Delete(ctx context.Context, invoiceItemID fastbill.ID) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("item.delete", idRequest{InvoiceItemID: invoiceItemID}), &res); err != nil {
		return err
	}
	return res.Err("item.delete")
}
