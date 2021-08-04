package customer

import (
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/request"
	"github.com/malsch-solutions/fastbill-go-sdk/response"
	"github.com/stretchr/testify/assert"
)

type dummySession struct {
}

func (c *dummySession) DoRequest(fastBillRequest request.Request) (response.Response, error) {
	return response.Response{
		Response: getResponse{Customers: []Customer{
			{},
		}},
	}, nil
}

func TestNewCustomerClient(t *testing.T) {
	client := NewCustomerClient(&dummySession{})
	assert.IsType(t, &Client{}, client)
}

func TestCustomerClientGet(t *testing.T) {
	client := NewCustomerClient(&dummySession{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Customer{}, resp)
	assert.Len(t, resp, 1)
}