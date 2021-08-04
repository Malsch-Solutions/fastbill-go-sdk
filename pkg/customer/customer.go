package customer

import (
	"errors"

	"github.com/malsch-solutions/fastbill-go-sdk/pkg/client"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/request"
)

//Client includes all customer api services
type Client struct {
	client *client.Client
}

//NewCustomerClient creates a new customer api client
func NewCustomerClient(c *client.Client) *Client {
	cClient := Client{client: c}
	return &cClient
}

type getResponse struct {
	Customers []Customer `json:"CUSTOMERS"`
}

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
