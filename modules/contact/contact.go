package contact

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all contact api services
type Client struct {
	client service.Service
}

//NewContactClient creates a new contact api client
func NewContactClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all contacts restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Contact, error) {

	fastBillRequest := request.NewRequestWithFilters("contact.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Contact, 0), err
	}

	var contactResponse getResponse
	err = mapstructure.Decode(res.Response, &contactResponse)
	if err != nil {
		return make([]Contact, 0), fmt.Errorf("failed to parse contact response: %s", err.Error())
	}

	return contactResponse.Contacts, nil
}

//Create create a contact
func (c *Client) Create(contact *Contact) (CreateResponse, error) {

	var responseContact CreateResponse

	fastBillRequest := request.NewRequestWithData("contact.create", contact)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseContact, err
	}

	err = mapstructure.Decode(res.Response, &responseContact)
	if err != nil {
		return responseContact, fmt.Errorf("failed to parse contact response: %s", err.Error())
	}

	return responseContact, nil
}

//Update update a contact
func (c *Client) Update(contact *Contact) (UpdateResponse, error) {

	var responseContact UpdateResponse

	fastBillRequest := request.NewRequestWithData("contact.update", contact)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseContact, err
	}

	err = mapstructure.Decode(res.Response, &responseContact)
	if err != nil {
		return responseContact, fmt.Errorf("failed to parse contact response: %s", err.Error())
	}

	return responseContact, nil
}

//Delete delete a contact
func (c *Client) Delete(contactID string, customerID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("contact.delete", deleteRequest{ContactID: contactID, CustomerID: customerID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse contact response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
