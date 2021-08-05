package document

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
	"github.com/stretchr/testify/assert"
)

type dummyService struct {
}

func (c *dummyService) DoRequest(fastBillRequest request.Request) (response.Response, error) {

	if fastBillRequest.Service == "document.get" {
		return response.Response{
			Response: Response{
				Documents: []Document{
					{},
				},
				Folders: map[string]Folder{
					"1": {},
				},
			},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	if fastBillRequest.Service == "document.create" {
		return response.Response{
			Response: CreateResponse{DocumentID: 1, Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func TestNewDocumentClient(t *testing.T) {
	client := NewDocumentClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestDocumentClientGet(t *testing.T) {
	client := NewDocumentClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, Response{}, resp)
}

func TestDocumentClientCreate(t *testing.T) {
	client := NewDocumentClient(&dummyService{})

	resp, err := client.Create(&Document{}, strings.NewReader(""), "file.txt")

	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
	assert.NotEmpty(t, resp.DocumentID)
	assert.Equal(t, resp.Status, "success")
}

type dummyErrorService struct {
}

func (c *dummyErrorService) DoRequest(_ request.Request) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func (c *dummyErrorService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestDocumentErrorClientGet(t *testing.T) {
	client := NewDocumentClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestDocumentErrorClientCreate(t *testing.T) {
	client := NewDocumentClient(&dummyErrorService{})
	_, err := client.Create(&Document{}, strings.NewReader(""), "file.txt")
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

func TestDocumentWrongStructClientGet(t *testing.T) {
	client := NewDocumentClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestDocumentWrongStructClientCreate(t *testing.T) {
	client := NewDocumentClient(&dummyWrongStructService{})
	_, err := client.Create(&Document{}, strings.NewReader(""), "file.txt")
	assert.Error(t, err)
}
