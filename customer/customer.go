package customer

import (
	"errors"

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

	fastBillRequest := request.NewRequest("customer.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Customer, 0), err
	}

	var customerResponse getResponse
	err = mapstructure.Decode(res.Response, &customerResponse)
	if err != nil {
		return make([]Customer, 0), errors.New("failed to parse customer response")
	}

	return customerResponse.Customers, nil
}
