package template

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

	if fastBillRequest.Service == "template.get" {
		return response.Response{
			Response: getResponse{Templates: []Template{
				{},
			}},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewTemplateClient(t *testing.T) {
	client := NewTemplateClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestTemplateClientGet(t *testing.T) {
	client := NewTemplateClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{})
	assert.NoError(t, err)
	assert.IsType(t, []Template{}, resp)
	assert.Len(t, resp, 1)
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func (c *dummyErrorService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestTemplateErrorClientGet(t *testing.T) {
	client := NewTemplateClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{})
	assert.Error(t, err)
}

type dummyWrongStructService struct {
}

func (c *dummyWrongStructService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func (c *dummyWrongStructService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{
		Response: true,
	}, nil
}

func TestTemplateWrongStructClientGet(t *testing.T) {
	client := NewTemplateClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{})
	assert.Error(t, err)
}
