package time

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all time api services
type Client struct {
	client service.Service
}

// NewTimeClient creates a new time api client
func NewTimeClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all times restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Time, error) {

	fastBillRequest := request.NewRequestWithFilters("time.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Time, 0), err
	}

	var timeResponse getResponse
	err = mapstructure.Decode(res.Response, &timeResponse)
	if err != nil {
		return make([]Time, 0), fmt.Errorf("failed to parse time response: %s", err.Error())
	}

	return timeResponse.Times, nil
}

// Create create a time
func (c *Client) Create(time *Time) (CreateResponse, error) {

	var responseTime CreateResponse

	fastBillRequest := request.NewRequestWithData("time.create", time)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseTime, err
	}

	err = mapstructure.Decode(res.Response, &responseTime)
	if err != nil {
		return responseTime, fmt.Errorf("failed to parse time response: %s", err.Error())
	}

	return responseTime, nil
}

// Update update a time
func (c *Client) Update(time *Time) (UpdateResponse, error) {

	var responseTime UpdateResponse

	fastBillRequest := request.NewRequestWithData("time.update", time)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseTime, err
	}

	err = mapstructure.Decode(res.Response, &responseTime)
	if err != nil {
		return responseTime, fmt.Errorf("failed to parse time response: %s", err.Error())
	}

	return responseTime, nil
}

// Delete delete a time
func (c *Client) Delete(timeID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("time.delete", deleteRequest{TimeID: timeID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse time response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
