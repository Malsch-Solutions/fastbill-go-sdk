package invoice

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all invoice api services
type Client struct {
	client service.Service
}

//NewInvoiceClient creates a new invoice api client
func NewInvoiceClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all invoices restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Invoice, error) {

	fastBillRequest := request.NewRequestWithFilters("invoice.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Invoice, 0), err
	}

	var invoiceResponse getResponse
	err = mapstructure.Decode(res.Response, &invoiceResponse)
	if err != nil {
		return make([]Invoice, 0), fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return invoiceResponse.Invoices, nil
}

//Create create a invoice
func (c *Client) Create(invoice *Request) (CreateResponse, error) {

	var responseInvoice CreateResponse

	fastBillRequest := request.NewRequestWithData("invoice.create", invoice)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseInvoice, err
	}

	err = mapstructure.Decode(res.Response, &responseInvoice)
	if err != nil {
		return responseInvoice, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return responseInvoice, nil
}

//Update update a invoice
func (c *Client) Update(invoice *Request) (UpdateResponse, error) {

	var responseInvoice UpdateResponse

	fastBillRequest := request.NewRequestWithData("invoice.update", invoice)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseInvoice, err
	}

	err = mapstructure.Decode(res.Response, &responseInvoice)
	if err != nil {
		return responseInvoice, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return responseInvoice, nil
}

//Delete delete a invoice
func (c *Client) Delete(invoiceID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("invoice.delete", deleteRequest{InvoiceID: invoiceID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}

//Cancel cancel an invoice
func (c *Client) Cancel(invoiceID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("invoice.cancel", cancelRequest{InvoiceID: invoiceID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes cancelResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}

//Complete complete a invoice
func (c *Client) Complete(invoiceID string) (CompleteResponse, error) {
	var completeResponse CompleteResponse
	fastBillRequest := request.NewRequestWithData("invoice.complete", completeRequest{InvoiceID: invoiceID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return completeResponse, err
	}

	err = mapstructure.Decode(res.Response, &completeResponse)
	if err != nil {
		return completeResponse, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return completeResponse, nil
}

//SendByEmail send an invoice by email
func (c *Client) SendByEmail(sendByMailRequest *SendByMailRequest) (bool, error) {
	fastBillRequest := request.NewRequestWithData("invoice.sendbyemail", sendByMailRequest)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var sendByMailRes sendByMailResponse

	err = mapstructure.Decode(res.Response, &sendByMailRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return sendByMailRes.Status == "success", nil
}

//SendByPost send an invoice by post
func (c *Client) SendByPost(invoiceID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("invoice.sendbypost", sendByPostRequest{InvoiceID: invoiceID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var sendByPostRes sendByPostResponse

	err = mapstructure.Decode(res.Response, &sendByPostRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse invoice response: %s", err.Error())
	}

	return sendByPostRes.Status == "success", nil
}

//SetPaid set an invoice paid
func (c *Client) SetPaid(setPaidRequest *SetPaidRequest) (SetPaidResponse, error) {

	var setPaidResponse SetPaidResponse

	fastBillRequest := request.NewRequestWithData("invoice.setpaid", setPaidRequest)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return setPaidResponse, err
	}

	err = mapstructure.Decode(res.Response, &setPaidResponse)
	if err != nil {
		return setPaidResponse, fmt.Errorf("failed to parse setPaidRequest response: %s", err.Error())
	}

	return setPaidResponse, nil
}
