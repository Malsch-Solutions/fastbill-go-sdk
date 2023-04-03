package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RequestHandler service interface
type RequestHandler interface {
	ValidateAndGetData() (Event, error)
}

// NewWebhookRequestHandler create new handler which can handle fastbill webhook requests
func NewWebhookRequestHandler(req http.Request) RequestHandler {
	h := &FastbillWebhookRequestHandler{
		req,
	}

	return h
}

// FastbillWebhookRequestHandler can handle fastbill webhook http requests
type FastbillWebhookRequestHandler struct {
	request http.Request
}

// ValidateAndGetData validates the request and returns the event data
func (h *FastbillWebhookRequestHandler) ValidateAndGetData() (Event, error) {
	req := h.request

	var data Event
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

	body, _ := io.ReadAll(req.Body)

	if err := json.Unmarshal(body, &data); err != nil {
		return data, fmt.Errorf("invalid Request, invalid json body: %s", err.Error())
	}

	return data, nil
}
