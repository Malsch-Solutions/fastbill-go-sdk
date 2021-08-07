package webhook

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

	if fastBillRequest.Service == "webhook.get" {
		return response.Response{
			Response: getResponse{Webhooks: []Webhook{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "webhook.create" {
		return response.Response{
			Response: CreateResponse{WebhookID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "webhook.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewWebhookClient(t *testing.T) {
	client := NewWebhookClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestWebhookClientGet(t *testing.T) {
	client := NewWebhookClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{})
	assert.NoError(t, err)
	assert.IsType(t, []Webhook{}, resp)
	assert.Len(t, resp, 1)
}

func TestWebhookClientCreate(t *testing.T) {
	client := NewWebhookClient(&dummyService{})
	resp, err := client.Create(&Webhook{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestWebhookClientDelete(t *testing.T) {
	client := NewWebhookClient(&dummyService{})
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

func TestWebhookErrorClientGet(t *testing.T) {
	client := NewWebhookClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{})
	assert.Error(t, err)
}

func TestWebhookErrorClientCreate(t *testing.T) {
	client := NewWebhookClient(&dummyErrorService{})
	_, err := client.Create(&Webhook{})
	assert.Error(t, err)
}

func TestWebhookErrorClientDelete(t *testing.T) {
	client := NewWebhookClient(&dummyErrorService{})
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

func TestWebhookWrongStructClientGet(t *testing.T) {
	client := NewWebhookClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{})
	assert.Error(t, err)
}

func TestWebhookWrongStructClientCreate(t *testing.T) {
	client := NewWebhookClient(&dummyWrongStructService{})
	_, err := client.Create(&Webhook{})
	assert.Error(t, err)
}

func TestWebhookWrongStructClientDelete(t *testing.T) {
	client := NewWebhookClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
