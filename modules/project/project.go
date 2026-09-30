// Package project wraps the project.* services: projects of a customer
// with their tasks, which work times are booked on.
package project

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the project services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a project client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists projects. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Project, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("project.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Projects, nil
}

// Create creates a project. ProjectName and CustomerID are required.
func (c *Client) Create(ctx context.Context, project *Project) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("project.create", project), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("project.create")
}

// Update changes the project with project.ProjectID. Only the fields that
// are set are sent.
func (c *Client) Update(ctx context.Context, project *Project) error {
	return c.status(ctx, "project.update", project)
}

// Delete deletes a project.
func (c *Client) Delete(ctx context.Context, projectID fastbill.ID) error {
	return c.status(ctx, "project.delete", idRequest{ProjectID: projectID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
