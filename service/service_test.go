package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	client := NewService("foo", "bar")
	assert.IsType(t, &FastBillService{}, client)
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}

func TestDoRequest(t *testing.T) {
	mockClient := MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {

			assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
			assert.Equal(t, "POST", req.Method)

			defer func() {
				if err := req.Body.Close(); err != nil {
					log.Println(err)
				}
			}()

			body, _ := ioutil.ReadAll(req.Body)

			var fastBillRequest request.Request
			err := json.Unmarshal(body, &fastBillRequest)
			assert.NoError(t, err)

			fastBillResponse := response.Response{
				Request:  nil,
				Response: nil,
			}

			fastBillResponseJSON, _ := json.Marshal(fastBillResponse)

			return &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(fastBillResponseJSON))),
			}, nil
		},
	}

	client := NewServiceWithClient("foo", "bar", &mockClient)
	_, err := client.DoRequest(request.NewRequestWithData("foo.bar", nil))
	assert.NoError(t, err)
}

func TestDoWithNoBodyErrorRequest(t *testing.T) {
	mockClient := MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Body: ioutil.NopCloser(strings.NewReader("")),
			}, nil
		},
	}

	client := NewServiceWithClient("foo", "bar", &mockClient)
	_, err := client.DoRequest(request.NewRequestWithData("foo.bar", nil))
	assert.Error(t, err)
}

func TestDoWithInvalidRequestRequest(t *testing.T) {
	mockClient := MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			fastBillResponse := response.Response{
				Request:  nil,
				Response: nil,
			}

			fastBillResponseJSON, _ := json.Marshal(fastBillResponse)

			return &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(fastBillResponseJSON))),
			}, nil
		},
	}

	client := NewServiceWithClient("foo", "bar", &mockClient)

	req := request.NewRequestWithData("foo.bar", nil)
	req.Filter = make(chan int)
	_, err := client.DoRequest(req)
	assert.Error(t, err)
}

func TestDoWithFastBillErrorRequest(t *testing.T) {
	mockClient := MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			fastBillResponse := response.Response{
				Request: nil,
				Response: response.ErrorResponse{
					Errors: []string{
						"error",
					},
				},
			}

			fastBillResponseJSON, _ := json.Marshal(fastBillResponse)

			return &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(string(fastBillResponseJSON))),
			}, nil
		},
	}

	client := NewServiceWithClient("foo", "bar", &mockClient)
	_, err := client.DoRequest(request.NewRequestWithData("foo.bar", nil))
	assert.Error(t, err)
}

func TestDoWithClientErrorRequest(t *testing.T) {
	mockClient := MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("my awesome client error")
		},
	}

	client := NewServiceWithClient("foo", "bar", &mockClient)
	_, err := client.DoRequest(request.NewRequestWithData("foo.bar", nil))
	assert.Error(t, err)
}
