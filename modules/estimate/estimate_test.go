package estimate

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

	if fastBillRequest.Service == "estimate.get" {
		return response.Response{
			Response: getResponse{Estimates: []Estimate{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "estimate.create" {
		return response.Response{
			Response: CreateResponse{EstimateID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "estimate.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "estimate.createinvoice" {
		return response.Response{
			Response: CreateInvoiceResponse{InvoiceID: 1337},
		}, nil
	}

	if fastBillRequest.Service == "estimate.sendbyemail" {
		return response.Response{
			Response: sendByMailResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func TestNewEstimateClient(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestEstimateClientGet(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Estimate{}, resp)
	assert.Len(t, resp, 1)
}

func TestEstimateClientCreate(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	resp, err := client.Create(&Request{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestEstimateClientDelete(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestEstimateClientCreateInvoice(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	resp, err := client.CreateInvoice("1337")
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.InvoiceID)
}

func TestEstimateClientSendByMail(t *testing.T) {
	client := NewEstimateClient(&dummyService{})
	resp, err := client.SendByEmail(&SendByMailRequest{
		EstimateID: "1337",
	})
	assert.NoError(t, err)
	assert.True(t, resp)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestEstimateErrorClientGet(t *testing.T) {
	client := NewEstimateClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestEstimateErrorClientCreate(t *testing.T) {
	client := NewEstimateClient(&dummyErrorService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestEstimateErrorClientDelete(t *testing.T) {
	client := NewEstimateClient(&dummyErrorService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestEstimateErrorClientCreateInvoice(t *testing.T) {
	client := NewEstimateClient(&dummyErrorService{})
	_, err := client.CreateInvoice("1337")
	assert.Error(t, err)
}

func TestEstimateErrorClientSendByMail(t *testing.T) {
	client := NewEstimateClient(&dummyErrorService{})
	_, err := client.SendByEmail(&SendByMailRequest{
		EstimateID: "1337",
	})
	assert.Error(t, err)
}

type dummyWrongStructService struct {
}

func (c *dummyWrongStructService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{
		Response: true,
	}, nil
}

func TestEstimateWrongStructClientGet(t *testing.T) {
	client := NewEstimateClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestEstimateWrongStructClientCreate(t *testing.T) {
	client := NewEstimateClient(&dummyWrongStructService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestEstimateWrongStructClientDelete(t *testing.T) {
	client := NewEstimateClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestEstimateWrongStructClientCreateInvoice(t *testing.T) {
	client := NewEstimateClient(&dummyWrongStructService{})
	_, err := client.CreateInvoice("1337")
	assert.Error(t, err)
}

func TestEstimateWrongStructClientSendByMail(t *testing.T) {
	client := NewEstimateClient(&dummyWrongStructService{})
	_, err := client.SendByEmail(&SendByMailRequest{
		EstimateID: "1337",
	})
	assert.Error(t, err)
}
