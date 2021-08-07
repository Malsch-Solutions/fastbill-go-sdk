package article

import (
	"fmt"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/service"
	"github.com/mitchellh/mapstructure"
)

//Client includes all article api services
type Client struct {
	client service.Service
}

//NewArticleClient creates a new article api client
func NewArticleClient(c service.Service) *Client {
	cClient := Client{client: c}
	return &cClient
}

//Get get all articles restricted by the given filters
func (c *Client) Get(parameter *parameter.Parameter, filter *Filter) ([]Article, error) {

	fastBillRequest := request.NewRequestWithFilters("article.get", parameter, filter)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return make([]Article, 0), err
	}

	var articleResponse getResponse
	err = mapstructure.Decode(res.Response, &articleResponse)
	if err != nil {
		return make([]Article, 0), fmt.Errorf("failed to parse article response: %s", err.Error())
	}

	return articleResponse.Articles, nil
}

//Create create a article
func (c *Client) Create(article *Article) (CreateResponse, error) {

	var responseArticle CreateResponse

	fastBillRequest := request.NewRequestWithData("article.create", article)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseArticle, err
	}

	err = mapstructure.Decode(res.Response, &responseArticle)
	if err != nil {
		return responseArticle, fmt.Errorf("failed to parse article response: %s", err.Error())
	}

	return responseArticle, nil
}

//Update update a article
func (c *Client) Update(article *Article) (UpdateResponse, error) {

	var responseArticle UpdateResponse

	fastBillRequest := request.NewRequestWithData("article.update", article)
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return responseArticle, err
	}

	err = mapstructure.Decode(res.Response, &responseArticle)
	if err != nil {
		return responseArticle, fmt.Errorf("failed to parse article response: %s", err.Error())
	}

	return responseArticle, nil
}

//Delete delete a article
func (c *Client) Delete(articleID string) (bool, error) {
	fastBillRequest := request.NewRequestWithData("article.delete", deleteRequest{ArticleID: articleID})
	res, err := c.client.DoRequest(fastBillRequest)

	if err != nil {
		return false, err
	}

	var deleteRes deleteResponse

	err = mapstructure.Decode(res.Response, &deleteRes)
	if err != nil {
		return false, fmt.Errorf("failed to parse article response: %s", err.Error())
	}

	return deleteRes.Status == "success", nil
}
