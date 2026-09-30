// Package document wraps the document.* services: the document inbox and
// its folders.
package document

import (
	"context"
	"io"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the document services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns a document client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists the folders and documents of the document inbox. A nil filter
// lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) (GetResponse, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("document.get", page, filter), &res); err != nil {
		return GetResponse{}, err
	}
	if res.Items != nil {
		return *res.Items, nil
	}
	return res.GetResponse, nil
}

// Create uploads file to the document inbox under fileName.
func (c *Client) Create(ctx context.Context, req *Request, file io.Reader, fileName string) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.DoMultipart(ctx, fastbill.DataRequest("document.create", req), file, fileName, &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("document.create")
}
