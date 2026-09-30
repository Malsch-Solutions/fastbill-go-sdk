// Package article wraps the article.* services: the products and services
// of the product catalog, which invoice items can refer to by number.
package article

import (
	"context"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Client calls the article services.
type Client struct {
	r fastbill.Requester
}

// NewClient returns an article client.
func NewClient(r fastbill.Requester) *Client {
	return &Client{r: r}
}

// Get lists articles. A nil filter lists all.
func (c *Client) Get(ctx context.Context, page fastbill.Page, filter *Filter) ([]Article, error) {
	var res getResponse
	if err := c.r.Do(ctx, fastbill.GetRequest("article.get", page, filter), &res); err != nil {
		return nil, err
	}
	return res.Articles, nil
}

// Create creates an article. ArticleNumber, Title and UnitPrice are
// required.
func (c *Client) Create(ctx context.Context, article *Article) (CreateResponse, error) {
	var res CreateResponse
	if err := c.r.Do(ctx, fastbill.DataRequest("article.create", article), &res); err != nil {
		return res, err
	}
	return res, fastbill.StatusResponse{Status: res.Status}.Err("article.create")
}

// Update changes the article with article.ArticleID. Only the fields that
// are set are sent.
func (c *Client) Update(ctx context.Context, article *Article) error {
	return c.status(ctx, "article.update", article)
}

// Delete deletes an article.
func (c *Client) Delete(ctx context.Context, articleID fastbill.ID) error {
	return c.status(ctx, "article.delete", idRequest{ArticleID: articleID})
}

func (c *Client) status(ctx context.Context, service string, data any) error {
	var res fastbill.StatusResponse
	if err := c.r.Do(ctx, fastbill.DataRequest(service, data), &res); err != nil {
		return err
	}
	return res.Err(service)
}
