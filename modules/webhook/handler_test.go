package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/contact"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/estimate"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/invoice"

	"github.com/stretchr/testify/assert"
)

type invalidWebhookEvent struct {
	ID       string             `json:"id"`
	Type     string             `json:"type"`
	Customer *customer.Customer `json:"customer"`
	Contact  *contact.Contact   `json:"contact"`
	Invoice  invoice.Invoice    `json:"invoice"`
	Estimate estimate.Estimate  `json:"estimate"`
	Created  string             `json:"created"`
}

func TestNewHandler(t *testing.T) {
	client := NewWebhookRequestHandler(http.Request{})
	assert.IsType(t, &FastbillWebhookRequestHandler{}, client)
}

func getRequest() http.Request {
	body := Event{
		ID: 1337,
	}

	JSONBody, _ := json.Marshal(body)

	req := http.Request{Body: io.NopCloser(strings.NewReader(string(JSONBody))),
		Header: map[string][]string{}}
	req.Method = "POST"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "FastBill")

	return req
}

func TestValidateValidRequest(t *testing.T) {
	client := NewWebhookRequestHandler(getRequest())
	data, err := client.ValidateAndGetData()

	assert.IsType(t, Event{}, data)
	assert.Equal(t, 1337, data.ID)
	assert.NoError(t, err)
}

func TestValidateInValidMethodRequest(t *testing.T) {
	req := getRequest()
	req.Method = "PUT"
	client := NewWebhookRequestHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidUserAgentRequest(t *testing.T) {
	req := getRequest()
	req.Header.Set("User-Agent", "Wrong")
	client := NewWebhookRequestHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidContentTypeRequest(t *testing.T) {
	req := getRequest()
	req.Header.Set("Content-Type", "application/xml")
	client := NewWebhookRequestHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidBodyRequest(t *testing.T) {

	body := invalidWebhookEvent{
		ID: "1337",
	}

	req := getRequest()

	JSONBody, _ := json.Marshal(body)
	req.Body = io.NopCloser(strings.NewReader(string(JSONBody)))

	client := NewWebhookRequestHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}
