// Package time wraps the time.* services: work time entries booked on a
// customer's project and task.
package time

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the time services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a time client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists time entries. A nil filter lists the last 10 entries of the
// current month.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Time, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("time.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Times, nil
}

// Create books a time entry. CustomerID, ProjectID and StartTime are
// required.
func (c *Client) Create(ctx context.Context, entry *Time) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("time.create", entry), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("time.create")
}

// Update changes the time entry with entry.TimeID. Only the fields that
// are set are sent.
//
// The docs show time.update answering with TIME_ID and COMMENT but no
// STATUS, so a missing STATUS counts as success; ERRORS still fail.
func (c *Client) Update(ctx context.Context, entry *Time) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("time.update", entry), &res); err != nil {
		return err
	}
	if res.Status == "" {
		return nil
	}
	return res.Err("time.update")
}

// Delete deletes a time entry.
func (c *Client) Delete(ctx context.Context, timeID fastbill.ID) error {
	return c.status(ctx, "time.delete", idRequest{TimeID: timeID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
