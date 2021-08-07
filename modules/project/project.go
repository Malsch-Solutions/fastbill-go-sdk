package project

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all project api services
type Client struct {
	client service.Service
}

//NewProjectClient creates a new project api client
func NewProjectClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all projects restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Project, error) {

	fastBillRequest := request.NewRequestWithFilters("project.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Project, 0), err
	}

	var projectResponse getResponse
	err = mapstructure.Decode(res.Response, &projectResponse)
	if err != nil {
		return make([]Project, 0), fmt.Errorf("failed to parse project response: %s", err.Error())
	}

	return projectResponse.Projects, nil
}

//Create create a project
func (c *Client) Create(project *Project) (CreateResponse, error) {

	var responseProject CreateResponse

	fastBillRequest := request.NewRequestWithData("project.create", project)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseProject, err
	}

	err = mapstructure.Decode(res.Response, &responseProject)
	if err != nil {
		return responseProject, fmt.Errorf("failed to parse project response: %s", err.Error())
	}

	return responseProject, nil
}

//Update update a project
func (c *Client) Update(project *Project) (UpdateResponse, error) {

	var responseProject UpdateResponse

	fastBillRequest := request.NewRequestWithData("project.update", project)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseProject, err
	}

	err = mapstructure.Decode(res.Response, &responseProject)
	if err != nil {
		return responseProject, fmt.Errorf("failed to parse project response: %s", err.Error())
	}

	return responseProject, nil
}

//Delete delete a project
func (c *Client) Delete(projectID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("project.delete", deleteRequest{ProjectID: projectID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse project response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
