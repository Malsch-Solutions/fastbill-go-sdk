package revenue

import (
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
	"github.com/stretchr/testify/assert"
)

type dummyService struct {
}

func (c *dummyService) DoRequest(fastBillRequest request.Request) (response.Response, error) {

	if fastBillRequest.Service == "revenue.get" {
		return response.Response{
			Response: getResponse{Revenues: []Revenue{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "revenue.create" {
		return response.Response{
			Response: CreateResponse{InvoiceID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "revenue.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "revenue.setpaid" {
		return response.Response{
			Response: SetPaidResponse{Status: "success", InvoiceNumber: "1337"},
		}, nil
	}

	return response.Response{}, errors.New("unknown services")
}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	if fastBillRequest.Service == "revenue.create" {
		return response.Response{
			Response: CreateResponse{InvoiceID: 1, Status: "success"},
		}, nil
	}
	return response.Response{}, errors.New("unknown service")
}

func TestNewRevenueClient(t *testing.T) {
	client := NewRevenueClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestRevenueClientGet(t *testing.T) {
	client := NewRevenueClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Revenue{}, resp)
	assert.Len(t, resp, 1)
}

func TestRevenueClientCreate(t *testing.T) {
	client := NewRevenueClient(&dummyService{})
	resp, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestRevenueClientDelete(t *testing.T) {
	client := NewRevenueClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestRevenueClientSetPaid(t *testing.T) {
	client := NewRevenueClient(&dummyService{})
	resp, err := client.SetPaid(&SetPaidRequest{
		InvoiceID: "1337",
		PaidDate:  time.Now(),
	})
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func (c *dummyErrorService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestRevenueErrorClientGet(t *testing.T) {
	client := NewRevenueClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestRevenueErrorClientCreate(t *testing.T) {
	client := NewRevenueClient(&dummyErrorService{})
	_, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")
	assert.Error(t, err)
}

func TestRevenueErrorClientDelete(t *testing.T) {
	client := NewRevenueClient(&dummyErrorService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestRevenueErrorClientSetPaid(t *testing.T) {
	client := NewRevenueClient(&dummyErrorService{})
	_, err := client.SetPaid(&SetPaidRequest{
		InvoiceID: "1337",
		PaidDate:  time.Now(),
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

func (c *dummyWrongStructService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{
		Response: true,
	}, nil
}

func TestRevenueWrongStructClientGet(t *testing.T) {
	client := NewRevenueClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestRevenueWrongStructClientCreate(t *testing.T) {
	client := NewRevenueClient(&dummyWrongStructService{})
	_, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")
	assert.Error(t, err)
}

func TestRevenueWrongStructClientDelete(t *testing.T) {
	client := NewRevenueClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestRevenueWrongStructClientSetPaid(t *testing.T) {
	client := NewRevenueClient(&dummyWrongStructService{})
	_, err := client.SetPaid(&SetPaidRequest{
		InvoiceID: "1337",
		PaidDate:  time.Now(),
	})
	assert.Error(t, err)
}
