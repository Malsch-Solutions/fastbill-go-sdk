package expense

import (
	"fmt"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all expense api services
type Client struct {
	client service.Service
}

//NewExpenseClient creates a new expense api client
func NewExpenseClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all expenses restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Expense, error) {

	fastBillRequest := request.NewRequestWithFilters("expense.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Expense, 0), err
	}

	var expenseResponse getResponse
	err = mapstructure.Decode(res.Response, &expenseResponse)
	if err != nil {
		return make([]Expense, 0), fmt.Errorf("failed to parse expense response: %s", err.Error())
	}

	return expenseResponse.Expenses, nil
}

//Create a expense
func (c *Client) Create(req *Request, file io.Reader, fileName string) (CreateResponse, error) {

	var responseDocument CreateResponse

	fastBillRequest := request.NewRequestWithData("expense.create", req)
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
