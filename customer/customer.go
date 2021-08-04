package customer

import (
	"errors"

	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/request"
	"github.com/malsch-solutions/fastbill-go-sdk/session"
)

//Client includes all customer api services
type Client struct {
	client *session.Session
}

//NewCustomerClient creates a new customer api client
func NewCustomerClient(c *session.Session) *Client {
	cClient := Client{client: c}
	return &cClient
}

type getResponse struct {
	Customers []Customer `json:"CUSTOMERS"`
}

//Get get all customers restricted by the given filters
func (c Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Customer, error) {

	fastBillRequest := request.NewRequest("customer.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Customer, 0), err
	}

	customerResponse, ok := res.Response.(getResponse)

	if !ok {
		return make([]Customer, 0), errors.New("failed to parse customer response")
	}

	return customerResponse.Customers, nil
}
