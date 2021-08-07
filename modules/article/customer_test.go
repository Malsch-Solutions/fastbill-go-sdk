package article

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

	if fastBillRequest.Service == "article.get" {
		return response.Response{
			Response: getResponse{Articles: []Article{
				{},
			}},
		}, nil
	}

	if fastBillRequest.Service == "article.create" {
		return response.Response{
			Response: CreateResponse{ArticleID: 10, Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "article.update" {
		return response.Response{
			Response: UpdateResponse{ArticleID: "10", Status: "success"},
		}, nil
	}

	if fastBillRequest.Service == "article.delete" {
		return response.Response{
			Response: deleteResponse{Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")

}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	return response.Response{}, errors.New("unknown service")
}

func TestNewArticleClient(t *testing.T) {
	client := NewArticleClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestArticleClientGet(t *testing.T) {
	client := NewArticleClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Article{}, resp)
	assert.Len(t, resp, 1)
}

func TestArticleClientCreate(t *testing.T) {
	client := NewArticleClient(&dummyService{})
	resp, err := client.Create(&Article{})
	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
}

func TestArticleClientUpdate(t *testing.T) {
	client := NewArticleClient(&dummyService{})
	resp, err := client.Update(&Article{})
	assert.NoError(t, err)
	assert.IsType(t, UpdateResponse{}, resp)
}

func TestArticleClientDelete(t *testing.T) {
	client := NewArticleClient(&dummyService{})
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

func TestArticleErrorClientGet(t *testing.T) {
	client := NewArticleClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestArticleErrorClientCreate(t *testing.T) {
	client := NewArticleClient(&dummyErrorService{})
	_, err := client.Create(&Article{})
	assert.Error(t, err)
}

func TestArticleErrorClientUpdate(t *testing.T) {
	client := NewArticleClient(&dummyErrorService{})
	_, err := client.Update(&Article{})
	assert.Error(t, err)
}

func TestArticleErrorClientDelete(t *testing.T) {
	client := NewArticleClient(&dummyErrorService{})
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

func TestArticleWrongStructClientGet(t *testing.T) {
	client := NewArticleClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestArticleWrongStructClientCreate(t *testing.T) {
	client := NewArticleClient(&dummyWrongStructService{})
	_, err := client.Create(&Article{})
	assert.Error(t, err)
}

func TestArticleWrongStructClientUpdate(t *testing.T) {
	client := NewArticleClient(&dummyWrongStructService{})
	_, err := client.Update(&Article{})
	assert.Error(t, err)
}

func TestArticleWrongStructClientDelete(t *testing.T) {
	client := NewArticleClient(&dummyWrongStructService{})
	_, err := client.Delete("1337")
	assert.Error(t, err)
}
