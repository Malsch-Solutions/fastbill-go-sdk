package recurring

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all recurring api services
type Client struct {
	client service.Service
}

//NewRecurringClient creates a new recurring api client
func NewRecurringClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all recurrings restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Recurring, error) {

	fastBillRequest := request.NewRequestWithFilters("recurring.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Recurring, 0), err
	}

	var recurringResponse getResponse
	err = mapstructure.Decode(res.Response, &recurringResponse)
	if err != nil {
		return make([]Recurring, 0), fmt.Errorf("failed to parse recurring response: %s", err.Error())
	}

	return recurringResponse.Recurrings, nil
}

//Create create a recurring
func (c *Client) Create(recurring *Request) (CreateResponse, error) {

	var responseRecurring CreateResponse

	fastBillRequest := request.NewRequestWithData("recurring.create", recurring)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseRecurring, err
	}

	err = mapstructure.Decode(res.Response, &responseRecurring)
	if err != nil {
		return responseRecurring, fmt.Errorf("failed to parse recurring response: %s", err.Error())
	}

	return responseRecurring, nil
}

//Update update a recurring
func (c *Client) Update(recurring *Request) (UpdateResponse, error) {

	var responseRecurring UpdateResponse

	fastBillRequest := request.NewRequestWithData("recurring.update", recurring)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseRecurring, err
	}

	err = mapstructure.Decode(res.Response, &responseRecurring)
	if err != nil {
		return responseRecurring, fmt.Errorf("failed to parse recurring response: %s", err.Error())
	}

	return responseRecurring, nil
}

//Delete delete a recurring
func (c *Client) Delete(recurringID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("recurring.delete", deleteRequest{InvoiceID: recurringID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse recurring response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
