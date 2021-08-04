package customer

import (
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/request"
	"github.com/malsch-solutions/fastbill-go-sdk/response"
	"github.com/stretchr/testify/assert"
)

type dummyService struct {
}

func (c *dummyService) DoRequest(fastBillRequest request.Request) (response.Response, error) {

	if fastBillRequest.Service == "customer.get" {
		return response.Response{
			Response: getResponse{Customers: []Customer{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "customer.create" {
		return response.Response{
			Response: Response{CustomerID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "customer.update" {
		return response.Response{
			Response: Response{CustomerID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "customer.delete" {
		return response.Response{
			Response: deleteResponse{ Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func TestNewCustomerClient(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestCustomerClientGet(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Customer{}, resp)
	assert.Len(t, resp, 1)
}

func TestCustomerClientCreate(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Create(&Customer{})
	assert.NoError(t, err)
	assert.IsType(t, Response{}, resp)
}

func TestCustomerClientUpdate(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Update(&Customer{})
	assert.NoError(t, err)
	assert.IsType(t, Response{}, resp)
}

func TestCustomerClientDelete(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}
