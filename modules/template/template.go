// Package template wraps the template.* services: the invoice and estimate
// layouts of the account.
package template

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the template services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a template client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists the templates. FastBill has no filter for template.get.
func (c *Client) Get(ctx context.Context, page fastbill.Page) ([]Template, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest[struct{}]("template.get", page, nil), &res); err != nil {
		return nil, err
	}
	return res.Templates, nil
}
