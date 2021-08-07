package webhook_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/contact"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/customer"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/estimate"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/invoice"

	"github.com/stretchr/testify/assert"
)

type InvalidWebhookData struct {
	ID       string             `json:"id"`
	Type     string             `json:"type"`
	Customer *customer.Customer `json:"customer"`
	Contact  *contact.Contact   `json:"contact"`
	Invoice  invoice.Invoice    `json:"invoice"`
	Estimate estimate.Estimate  `json:"estimate"`
	Created  string             `json:"created"`
}

func TestNewHandler(t *testing.T) {
	client := NewWebhookHandler(http.Request{})
	assert.IsType(t, &FastbillWebhookHandler{}, client)
}

func getRequest() http.Request {
	body := WebhookData{
		ID: 1337,
	}

	JSONBody, _ := json.Marshal(body)

	req := http.Request{Body: ioutil.NopCloser(strings.NewReader(string(JSONBody))),
		Header: map[string][]string{}}
	req.Method = "POST"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "FastBill")

	return req
}

func TestValidateValidRequest(t *testing.T) {
	client := NewWebhookHandler(getRequest())
	data, err := client.ValidateAndGetData()

	assert.IsType(t, WebhookData{}, data)
	assert.Equal(t, 1337, data.ID)
	assert.NoError(t, err)
}

func TestValidateInValidMethodRequest(t *testing.T) {
	req := getRequest()
	req.Method = "PUT"
	client := NewWebhookHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidUserAgentRequest(t *testing.T) {
	req := getRequest()
	req.Header.Set("User-Agent", "Wrong")
	client := NewWebhookHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidContentTypeRequest(t *testing.T) {
	req := getRequest()
	req.Header.Set("Content-Type", "application/xml")
	client := NewWebhookHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}

func TestValidateInValidBodyRequest(t *testing.T) {

	body := InvalidWebhookData{
		ID: "1337",
	}

	req := getRequest()

	JSONBody, _ := json.Marshal(body)
	req.Body = ioutil.NopCloser(strings.NewReader(string(JSONBody)))

	client := NewWebhookHandler(req)
	_, err := client.ValidateAndGetData()

	assert.Error(t, err)
}
