package webhook

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all webhook api services
type Client struct {
	client service.Service
}

// NewWebhookClient creates a new webhook api client
func NewWebhookClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all webhooks restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter) ([]Webhook, error) {

	fastBillRequest := request.NewRequestWithFilters("webhook.get", parameter, nil)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Webhook, 0), err
	}

	var webhookResponse getResponse
	err = mapstructure.Decode(res.Response, &webhookResponse)
	if err != nil {
		return make([]Webhook, 0), fmt.Errorf("failed to parse webhook response: %s", err.Error())
	}

	return webhookResponse.Webhooks, nil
}

// Create create a webhook
func (c *Client) Create(webhook *Webhook) (CreateResponse, error) {

	var responseWebhook CreateResponse

	fastBillRequest := request.NewRequestWithData("webhook.create", webhook)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseWebhook, err
	}

	err = mapstructure.Decode(res.Response, &responseWebhook)
	if err != nil {
		return responseWebhook, fmt.Errorf("failed to parse webhook response: %s", err.Error())
	}

	return responseWebhook, nil
}

// Delete delete a webhook
func (c *Client) Delete(webhookID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("webhook.delete", deleteRequest{WebhookID: webhookID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse webhook response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
