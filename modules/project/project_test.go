package project

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

	if fastBillRequest.Service == "project.get" {
		return response.Response{
			Response: getResponse{Projects: []Project{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "project.create" {
		return response.Response{
			Response: CreateResponse{ProjectID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "project.update" {
		return response.Response{
			Response: UpdateResponse{Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "project.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewProjectClient(t *testing.T) {
	client := NewProjectClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestProjectClientGet(t *testing.T) {
	client := NewProjectClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Project{}, resp)
	assert.Len(t, resp, 1)
}

func TestProjectClientCreate(t *testing.T) {
	client := NewProjectClient(&dummyService{})
	resp, err := client.Create(&Project{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestProjectClientUpdate(t *testing.T) {
	client := NewProjectClient(&dummyService{})
	resp, err := client.Update(&Project{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestProjectClientDelete(t *testing.T) {
	client := NewProjectClient(&dummyService{})
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

func TestProjectErrorClientGet(t *testing.T) {
	client := NewProjectClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestProjectErrorClientCreate(t *testing.T) {
	client := NewProjectClient(&dummyErrorService{})
	_, err := client.Create(&Project{})
	assert.Error(t, err)
}

func TestProjectErrorClientUpdate(t *testing.T) {
	client := NewProjectClient(&dummyErrorService{})
	_, err := client.Update(&Project{})
	assert.Error(t, err)
}

func TestProjectErrorClientDelete(t *testing.T) {
	client := NewProjectClient(&dummyErrorService{})
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

func TestProjectWrongStructClientGet(t *testing.T) {
	client := NewProjectClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestProjectWrongStructClientCreate(t *testing.T) {
	client := NewProjectClient(&dummyWrongStructService{})
	_, err := client.Create(&Project{})
	assert.Error(t, err)
}

func TestProjectWrongStructClientUpdate(t *testing.T) {
	client := NewProjectClient(&dummyWrongStructService{})
	_, err := client.Update(&Project{})
	assert.Error(t, err)
}

func TestProjectWrongStructClientDelete(t *testing.T) {
	client := NewProjectClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
