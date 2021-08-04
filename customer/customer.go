package customer

import (
	"errors"
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all customer api services
type Client struct {
	client service.Service
}

//NewCustomerClient creates a new customer api client
func NewCustomerClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all customers restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Customer, error) {

	fastBillRequest := request.NewRequestWithFilters("customer.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Customer, 0), err
	}

	var customerResponse getResponse
	err = mapstructure.Decode(res.Response, &customerResponse)
	if err != nil {
		return make([]Customer, 0), errors.New(fmt.Sprintf("failed to parse customer response: %s", err.Error()))
	}

	return customerResponse.Customers, nil
}

//Create create a customer
func (c *Client) Create(customer *Customer) (Response, error) {

	var responseCustomer Response

	fastBillRequest := request.NewRequestWithData("customer.create", customer)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseCustomer, err
	}

	err = mapstructure.Decode(res.Response, &responseCustomer)
	if err != nil {
		return responseCustomer, errors.New(fmt.Sprintf("failed to parse customer response: %s", err.Error()))
	}

	return responseCustomer, nil
}

//Update update a customer
func (c *Client) Update(customer *Customer) (Response, error) {

	var responseCustomer Response

	fastBillRequest := request.NewRequestWithData("customer.update", customer)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseCustomer, err
	}

	err = mapstructure.Decode(res.Response, &responseCustomer)
	if err != nil {
		return responseCustomer, errors.New(fmt.Sprintf("failed to parse customer response: %s", err.Error()))
	}

	return responseCustomer, nil
}

//Delete delete a customer
func (c *Client) Delete(customerID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("customer.delete", deleteRequest{CustomerID: customerID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, errors.New(fmt.Sprintf("failed to parse customer response: %s", err.Error()))
	}

	return deleteRes.Status == "success", nil
}
