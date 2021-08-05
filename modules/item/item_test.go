package item

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

	if fastBillRequest.Service == "item.get" {
		return response.Response{
			Response: getResponse{Items: []Item{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "item.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func TestNewItemClient(t *testing.T) {
	client := NewItemClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestItemClientGet(t *testing.T) {
	client := NewItemClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Item{}, resp)
	assert.Len(t, resp, 1)
}

func TestItemClientDelete(t *testing.T) {
	client := NewItemClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestItemErrorClientGet(t *testing.T) {
	client := NewItemClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestCustomerErrorClientDelete(t *testing.T) {
	client := NewItemClient(&dummyErrorService{})
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

func TestItemWrongStructClientGet(t *testing.T) {
	client := NewItemClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestItemWrongStructClientDelete(t *testing.T) {
	client := NewItemClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
