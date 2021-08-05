package invoice

import (
	"errors"
	"testing"
	"time"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/request"
	"github.com/stretchr/testify/assert"
)

type dummyService struct {
}

func (c *dummyService) DoRequest(fastBillRequest request.Request) (response.Response, error) {

	if fastBillRequest.Service == "invoice.get" {
		return response.Response{
			Response: getResponse{Invoices: []Invoice{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "invoice.create" {
		return response.Response{
			Response: CreateResponse{InvoiceID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.update" {
		return response.Response{
			Response: UpdateResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.complete" {
		return response.Response{
			Response: CompleteResponse{Status: "success", InvoiceID: "1337"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.cancel" {
		return response.Response{
			Response: cancelResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.sendbyemail" {
		return response.Response{
			Response: sendByMailResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.sendbypost" {
		return response.Response{
			Response: sendByPostResponse{Status: "success", RemainingCredits: "100"},
		}, nil
	}

	if fastBillRequest.Service == "invoice.setpaid" {
		return response.Response{
			Response: SetPaidResponse{Status: "success", InvoiceNumber: "1337"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func TestNewInvoiceClient(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestInvoiceClientGet(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Invoice{}, resp)
	assert.Len(t, resp, 1)
}

func TestInvoiceClientCreate(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Create(&Request{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestInvoiceClientUpdate(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Update(&Request{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestInvoiceClientDelete(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Delete("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestInvoiceClientComplete(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Complete("1337")
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

func TestInvoiceClientCancel(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.Cancel("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestInvoiceClientSendByMail(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.SendByEmail(&SendByMailRequest{
		InvoiceID: "1337",
	})
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestInvoiceClientSendByPost(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.SendByPost("1337")
	assert.NoError(t, err)
	assert.True(t, resp)
}

func TestInvoiceClientSetPaid(t *testing.T) {
	client := NewInvoiceClient(&dummyService{})
	resp, err := client.SetPaid(&SetPaidRequest{
		InvoiceID:     "1337",
		PaidDate:      time.Now(),
		PaymentMethod: "Cash",
	})
	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestInvoiceErrorClientGet(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestInvoiceErrorClientCreate(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestInvoiceErrorClientUpdate(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Update(&Request{})
	assert.Error(t, err)
}

func TestInvoiceErrorClientDelete(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestInvoiceErrorClientComplete(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Complete("1337")
	assert.Error(t, err)
}

func TestInvoiceErrorClientCancel(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.Cancel("1337")
	assert.Error(t, err)
}

func TestInvoiceErrorClientSendByMail(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.SendByEmail(&SendByMailRequest{
		InvoiceID: "1337",
	})
	assert.Error(t, err)
}

func TestInvoiceErrorClientSendByPost(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.SendByPost("1337")
	assert.Error(t, err)
}

func TestInvoiceErrorClientSetPaid(t *testing.T) {
	client := NewInvoiceClient(&dummyErrorService{})
	_, err := client.SetPaid(&SetPaidRequest{
		InvoiceID:     "1337",
		PaidDate:      time.Now(),
		PaymentMethod: "Cash",
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

func TestInvoiceWrongStructClientGet(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientCreate(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Create(&Request{})
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientUpdate(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Update(&Request{})
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientDelete(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientComplete(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Complete("1337")
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientCancel(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.Cancel("1337")
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientSendByMail(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.SendByEmail(&SendByMailRequest{
		InvoiceID: "1337",
	})
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientSendByPost(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.SendByPost("1337")
	assert.Error(t, err)
}

func TestInvoiceWrongStructClientSetPaid(t *testing.T) {
	client := NewInvoiceClient(&dummyWrongStructService{})
	_, err := client.SetPaid(&SetPaidRequest{
		InvoiceID:     "1337",
		PaidDate:      time.Now(),
		PaymentMethod: "Cash",
	})
	assert.Error(t, err)
}
