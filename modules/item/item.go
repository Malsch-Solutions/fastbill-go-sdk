package item

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all item api services
type Client struct {
	client service.Service
}

// NewItemClient creates a new item api client
func NewItemClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all items restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Item, error) {

	fastBillRequest := request.NewRequestWithFilters("item.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Item, 0), err
	}

	var itemResponse getResponse
	err = mapstructure.Decode(res.Response, &itemResponse)
	if err != nil {
		return make([]Item, 0), fmt.Errorf("failed to parse item response: %s", err.Error())
	}

	return itemResponse.Items, nil
}

// Delete delete a item
func (c *Client) Delete(itemID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("item.delete", deleteRequest{InvoiceItemID: itemID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse item response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
