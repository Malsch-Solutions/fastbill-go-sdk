package expense

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

	if fastBillRequest.Service == "expense.get" {
		return response.Response{
			Response: getResponse{Expenses: []Expense{
				{},
			}},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func (c *dummyService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	if fastBillRequest.Service == "expense.create" {
		return response.Response{
			Response: CreateResponse{InvoiceID: 1, Status: "success"},
		}, nil
	}

	return response.Response{}, errors.New("unknown service")
}

func TestNewExpenseClient(t *testing.T) {
	client := NewExpenseClient(&dummyService{})
	assert.IsType(t, &Client{}, client)
}

func TestExpenseClientGet(t *testing.T) {
	client := NewExpenseClient(&dummyService{})
	resp, err := client.Get(&parameter.Parameter{}, nil)
	assert.NoError(t, err)
	assert.IsType(t, []Expense{}, resp)
	assert.Len(t, resp, 1)
}

func TestExpenseClientCreate(t *testing.T) {
	client := NewExpenseClient(&dummyService{})

	resp, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")

	assert.NoError(t, err)
	assert.IsType(t, CreateResponse{}, resp)
	assert.NotEmpty(t, resp.InvoiceID)
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

func TestExpenseErrorClientGet(t *testing.T) {
	client := NewExpenseClient(&dummyErrorService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestExpenseErrorClientCreate(t *testing.T) {
	client := NewExpenseClient(&dummyErrorService{})
	_, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")
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

func TestExpenseWrongStructClientGet(t *testing.T) {
	client := NewExpenseClient(&dummyWrongStructService{})
	_, err := client.Get(&parameter.Parameter{}, nil)
	assert.Error(t, err)
}

func TestExpenseWrongStructClientCreate(t *testing.T) {
	client := NewExpenseClient(&dummyWrongStructService{})
	_, err := client.Create(&Request{}, strings.NewReader(""), "file.txt")
	assert.Error(t, err)
}
