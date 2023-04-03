package template

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all template api services
type Client struct {
	client service.Service
}

// NewTemplateClient creates a new template api client
func NewTemplateClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all templates restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter) ([]Template, error) {

	fastBillRequest := request.NewRequestWithFilters("template.get", parameter, nil)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Template, 0), err
	}

	var templateResponse getResponse
	err = mapstructure.Decode(res.Response, &templateResponse)
	if err != nil {
		return make([]Template, 0), fmt.Errorf("failed to parse template response: %s", err.Error())
	}

	return templateResponse.Templates, nil
}
