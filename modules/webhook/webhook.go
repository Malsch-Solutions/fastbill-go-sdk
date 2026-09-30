// Package webhook wraps the webhook.* services, which register endpoints
// that FastBill notifies about changes, and parses those notifications
// with ParseEvent.
package webhook

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the webhook services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a webhook client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists the registered webhooks. FastBill has no filter for
// webhook.get.
func (c *Client) Get(ctx context.Context, page fastbill.Page) ([]Webhook, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest[struct{}]("webhook.get", page, nil), &res); err != nil {
		return nil, err
	}
	return res.Webhooks, nil
}

// Create registers a webhook.
func (c *Client) Create(ctx context.Context, req *Request) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("webhook.create", req), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("webhook.create")
}

// Delete removes a webhook registration.
func (c *Client) Delete(ctx context.Context, webhookID fastbill.ID) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("webhook.delete", idRequest{WebhookID: webhookID}), &res); err != nil {
		return err
	}
	return res.Err("webhook.delete")
}
