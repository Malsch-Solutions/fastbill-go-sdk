package estimate

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all estimate api services
type Client struct {
	client service.Service
}

// NewEstimateClient creates a new estimate api client
func NewEstimateClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all estimates restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Estimate, error) {

	fastBillRequest := request.NewRequestWithFilters("estimate.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Estimate, 0), err
	}

	var estimateResponse getResponse
	err = mapstructure.Decode(res.Response, &estimateResponse)
	if err != nil {
		return make([]Estimate, 0), fmt.Errorf("failed to parse estimate response: %s", err.Error())
	}

	return estimateResponse.Estimates, nil
}

// Create create a estimate
func (c *Client) Create(estimate *Request) (CreateResponse, error) {

	var responseEstimate CreateResponse

	fastBillRequest := request.NewRequestWithData("estimate.create", estimate)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseEstimate, err
	}

	err = mapstructure.Decode(res.Response, &responseEstimate)
	if err != nil {
		return responseEstimate, fmt.Errorf("failed to parse estimate response: %s", err.Error())
	}

	return responseEstimate, nil
}

// Delete delete a estimate
func (c *Client) Delete(estimateID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("estimate.delete", deleteRequest{EstimateID: estimateID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse estimate response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}

// CreateInvoice create an invoice out of an estimate
func (c *Client) CreateInvoice(estimateID string) (CreateInvoiceResponse, error) {
	var createInvoiceResponse CreateInvoiceResponse

	fastBillRequest := request.NewRequestWithData("estimate.createinvoice", createInvoiceRequest{EstimateID: estimateID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return createInvoiceResponse, err
	}

	err = mapstructure.Decode(res.Response, &createInvoiceResponse)
	if err != nil {
		return createInvoiceResponse, fmt.Errorf("failed to parse estimate response: %s", err.Error())
	}

	return createInvoiceResponse, nil
}

// SendByEmail send an estimate by email
func (c *Client) SendByEmail(sendByMailRequest *SendByMailRequest) (bool, error) {
	fastBillRequest := request.NewRequestWithData("estimate.sendbyemail", sendByMailRequest)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var sendByMailRes sendByMailResponse

	err = mapstructure.Decode(res.Response, &sendByMailRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse estimate response: %s", err.Error())
	}

	return sendByMailRes.Status == "success", nil
}
