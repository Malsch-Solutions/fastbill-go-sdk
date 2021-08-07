package recurring

import (
	"errors"
	"io"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
	"github.com/stretchr/testify/assert"
)

type dummyService struct {
}

func (c *dummyService) DoRequest(fastBillRequest request.Request) (response.Response, error) {

	if fastBillRequest.Service == "recurring.get" {
		return response.Response{
			Response: getResponse{Recurrings: []Recurring{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "recurring.create" {
		return response.Response{
			Response: CreateResponse{InvoiceID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "recurring.update" {
		return response.Response{
			Response: UpdateResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "recurring.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewRecurringClient(t *testing.T) {
	client := NewRecurringClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestRecurringClientGet(t *testing.T) {
	client := NewRecurringClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Recurring{}, resp)
	assert.Len(t, resp, 1)
}

func TestRecurringClientCreate(t *testing.T) {
	client := NewRecurringClient(&dummyService{})
	resp, err := client.Create(&Request{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestRecurringClientUpdate(t *testing.T) {
	client := NewRecurringClient(&dummyService{})
	resp, err := client.Update(&Request{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestRecurringClientDelete(t *testing.T) {
	client := NewRecurringClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func (c *dummyErrorService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestRecurringErrorClientGet(t *testing.T) {
	client := NewRecurringClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestRecurringErrorClientCreate(t *testing.T) {
	client := NewRecurringClient(&dummyErrorService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestRecurringErrorClientUpdate(t *testing.T) {
	client := NewRecurringClient(&dummyErrorService{})
	_, err := client.Update(&Request{})
	assert.Error(t, err)
}

func TestRecurringErrorClientDelete(t *testing.T) {
	client := NewRecurringClient(&dummyErrorService{})
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

func (c *dummyWrongStructService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestRecurringWrongStructClientGet(t *testing.T) {
	client := NewRecurringClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestRecurringWrongStructClientCreate(t *testing.T) {
	client := NewRecurringClient(&dummyWrongStructService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestRecurringWrongStructClientUpdate(t *testing.T) {
	client := NewRecurringClient(&dummyWrongStructService{})
	_, err := client.Update(&Request{})
	assert.Error(t, err)
}

func TestRecurringWrongStructClientDelete(t *testing.T) {
	client := NewRecurringClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
