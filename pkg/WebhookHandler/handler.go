package webhook_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//WebhookHandler service interface
type WebhookHandler interface {
	ValidateAndGetData() (WebhookData, error)
}

func NewWebhookHandler(req http.Request) WebhookHandler {
	h := &FastbillWebhookHandler{
		req,
	}

	return h
}

type FastbillWebhookHandler struct {
	request http.Request
}

func (h *FastbillWebhookHandler) ValidateAndGetData() (WebhookData, error) {
	req := h.request

	var data WebhookData
	if req.Method != "POST" {
		return data, fmt.Errorf("invalid Request, wrong http method: got %s, expected post", req.Method)
	}

	if req.UserAgent() != "FastBill" {
		return data, fmt.Errorf("invalid Request, wrong user agent: got %s expected FastBill", req.UserAgent())
	}

	if req.Header.Get("Content-type") != "application/json" {
		return data, fmt.Errorf("invalid Request, wrong Content-Type header: got %s expected application/json", req.Method)
	}

	defer func() {
		_ = req.Body.Close()
	}()

	body, _ := ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(body, &data); err != nil {
		return data, fmt.Errorf("invalid Request, invalid json body: %s", err.Error())
	}

	return data, nil
}
