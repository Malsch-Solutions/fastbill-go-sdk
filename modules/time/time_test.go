package time

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

	if fastBillRequest.Service == "time.get" {
		return response.Response{
			Response: getResponse{Times: []Time{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "time.create" {
		return response.Response{
			Response: CreateResponse{TimeID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "time.update" {
		return response.Response{
			Response: UpdateResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "time.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewTimeClient(t *testing.T) {
	client := NewTimeClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestTimeClientGet(t *testing.T) {
	client := NewTimeClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Time{}, resp)
	assert.Len(t, resp, 1)
}

func TestTimeClientCreate(t *testing.T) {
	client := NewTimeClient(&dummyService{})
	resp, err := client.Create(&Time{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestTimeClientUpdate(t *testing.T) {
	client := NewTimeClient(&dummyService{})
	resp, err := client.Update(&Time{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestTimeClientDelete(t *testing.T) {
	client := NewTimeClient(&dummyService{})
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

func TestTimeErrorClientGet(t *testing.T) {
	client := NewTimeClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestTimeErrorClientCreate(t *testing.T) {
	client := NewTimeClient(&dummyErrorService{})
	_, err := client.Create(&Time{})
	assert.Error(t, err)
}

func TestTimeErrorClientUpdate(t *testing.T) {
	client := NewTimeClient(&dummyErrorService{})
	_, err := client.Update(&Time{})
	assert.Error(t, err)
}

func TestTimeErrorClientDelete(t *testing.T) {
	client := NewTimeClient(&dummyErrorService{})
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

func TestTimeWrongStructClientGet(t *testing.T) {
	client := NewTimeClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestTimeWrongStructClientCreate(t *testing.T) {
	client := NewTimeClient(&dummyWrongStructService{})
	_, err := client.Create(&Time{})
	assert.Error(t, err)
}

func TestTimeWrongStructClientUpdate(t *testing.T) {
	client := NewTimeClient(&dummyWrongStructService{})
	_, err := client.Update(&Time{})
	assert.Error(t, err)
}

func TestTimeWrongStructClientDelete(t *testing.T) {
	client := NewTimeClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
