package customer

import (
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
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
			Response: CreateResponse{CustomerID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "customer.update" {
		return response.Response{
			Response: UpdateResponse{CustomerID: "10", Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "customer.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
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
	assert.IsType(t, CreateResponse{}, resp)
}

func TestCustomerClientUpdate(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Update(&Customer{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestCustomerClientDelete(t *testing.T) {
	client := NewCustomerClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestCustomerErrorClientGet(t *testing.T) {
	client := NewCustomerClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestContacErrorClientCreate(t *testing.T) {
	client := NewCustomerClient(&dummyErrorService{})
	_, err := client.Create(&Customer{})
	assert.Error(t, err)
}

func TestContacErrorClientUpdate(t *testing.T) {
	client := NewCustomerClient(&dummyErrorService{})
	_, err := client.Update(&Customer{})
	assert.Error(t, err)
}

func TestContacErrorClientDelete(t *testing.T) {
	client := NewCustomerClient(&dummyErrorService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

type dummyWrongStructService struct {
}

func (c *dummyWrongStructService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{
		Response: true,
	}, nil
}

func TestCustomerWrongStructClientGet(t *testing.T) {
	client := NewCustomerClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestContacWrongStructClientCreate(t *testing.T) {
	client := NewCustomerClient(&dummyWrongStructService{})
	_, err := client.Create(&Customer{})
	assert.Error(t, err)
}

func TestContacWrongStructClientUpdate(t *testing.T) {
	client := NewCustomerClient(&dummyWrongStructService{})
	_, err := client.Update(&Customer{})
	assert.Error(t, err)
}

func TestContacWrongStructClientDelete(t *testing.T) {
	client := NewCustomerClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
