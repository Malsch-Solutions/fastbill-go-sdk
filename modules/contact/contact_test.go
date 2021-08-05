package contact

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

	if fastBillRequest.Service == "contact.get" {
		return response.Response{
			Response: getResponse{Contacts: []Contact{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "contact.create" {
		return response.Response{
			Response: CreateResponse{ContactID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "contact.update" {
		return response.Response{
			Response: UpdateResponse{ContactID: "10", Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "contact.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func TestNewContactClient(t *testing.T) {
	client := NewContactClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestContactClientGet(t *testing.T) {
	client := NewContactClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Contact{}, resp)
	assert.Len(t, resp, 1)
}

func TestContactClientCreate(t *testing.T) {
	client := NewContactClient(&dummyService{})
	resp, err := client.Create(&Contact{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestContactClientUpdate(t *testing.T) {
	client := NewContactClient(&dummyService{})
	resp, err := client.Update(&Contact{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestContactClientDelete(t *testing.T) {
	client := NewContactClient(&dummyService{})
	resp, err := client.Delete("1337", "17")
	assert.NoError(t, err)
	assert.True(t, resp)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestContactErrorClientGet(t *testing.T) {
	client := NewContactClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestContacErrorClientCreate(t *testing.T) {
	client := NewContactClient(&dummyErrorService{})
	_, err := client.Create(&Contact{})
	assert.Error(t, err)
}

func TestContacErrorClientUpdate(t *testing.T) {
	client := NewContactClient(&dummyErrorService{})
	_, err := client.Update(&Contact{})
	assert.Error(t, err)
}

func TestContacErrorClientDelete(t *testing.T) {
	client := NewContactClient(&dummyErrorService{})
	_, err := client.Delete("1337", "17")
	assert.Error(t, err)
}

type dummyWrongStructService struct {
}

func (c *dummyWrongStructService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{
		Response: true,
	}, nil
}

func TestContactWrongStructClientGet(t *testing.T) {
	client := NewContactClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestContactWrongStructClientCreate(t *testing.T) {
	client := NewContactClient(&dummyWrongStructService{})
	_, err := client.Create(&Contact{})
	assert.Error(t, err)
}

func TestContactWrongStructClientUpdate(t *testing.T) {
	client := NewContactClient(&dummyWrongStructService{})
	_, err := client.Update(&Contact{})
	assert.Error(t, err)
}

func TestContactWrongStructClientDelete(t *testing.T) {
	client := NewContactClient(&dummyWrongStructService{})
	_, err := client.Delete("1337", "17")
	assert.Error(t, err)
}
