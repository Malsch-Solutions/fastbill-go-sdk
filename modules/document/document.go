package document

import (
	"fmt"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all document api services
type Client struct {
	client service.Service
}

// NewDocumentClient creates a new document api client
func NewDocumentClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all documents restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) (Response, error) {
	var documentResponse getResponse

	fastBillRequest := request.NewRequestWithFilters("document.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return documentResponse.Items, err
	}

	err = mapstructure.Decode(res.Response, &documentResponse)
	if err != nil {
		return documentResponse.Items, fmt.Errorf("failed to parse document response: %s", err.Error())
	}

	return documentResponse.Items, nil
}

// Create create a document
func (c *Client) Create(document *Document, file io.Reader, fileName string) (CreateResponse, error) {

	var responseDocument CreateResponse

	fastBillRequest := request.NewRequestWithData("document.create", document)
	res, err := c.client.DoMultiPartRequest(fastBillRequest, file, fileName)

	if err != nil {
		return responseDocument, err
	}

	err = mapstructure.Decode(res.Response, &responseDocument)
	if err != nil {
		return responseDocument, fmt.Errorf("failed to parse document response: %s", err.Error())
	}

	return responseDocument, nil
}
