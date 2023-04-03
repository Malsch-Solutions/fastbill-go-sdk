package revenue

import (
	"fmt"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

// Client includes all revenue api services
type Client struct {
	client service.Service
}

// NewRevenueClient creates a new revenue api client
func NewRevenueClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

// Get get all revenues restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Revenue, error) {

	fastBillRequest := request.NewRequestWithFilters("revenue.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Revenue, 0), err
	}

	var revenueResponse getResponse
	err = mapstructure.Decode(res.Response, &revenueResponse)
	if err != nil {
		return make([]Revenue, 0), fmt.Errorf("failed to parse revenue response: %s", err.Error())
	}

	return revenueResponse.Revenues, nil
}

// Create create a revenue
func (c *Client) Create(revenue *Request, file io.Reader, fileName string) (CreateResponse, error) {

	var responseRevenue CreateResponse

	fastBillRequest := request.NewRequestWithData("revenue.create", revenue)
	res, err := c.client.DoMultiPartRequest(fastBillRequest, file, fileName)

	if err != nil {
		return responseRevenue, err
	}

	err = mapstructure.Decode(res.Response, &responseRevenue)
	if err != nil {
		return responseRevenue, fmt.Errorf("failed to parse revenue response: %s", err.Error())
	}

	return responseRevenue, nil
}

// Delete delete a revenue
func (c *Client) Delete(revenueID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("revenue.delete", deleteRequest{RevenueID: revenueID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse revenue response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}

// SetPaid set an revenue paid
func (c *Client) SetPaid(setPaidRequest *SetPaidRequest) (SetPaidResponse, error) {

	var setPaidResponse SetPaidResponse

	fastBillRequest := request.NewRequestWithData("revenue.setpaid", setPaidRequest)
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
