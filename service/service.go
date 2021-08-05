package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
)

const baseURL string = "https://my.fastbill.com/api/1.0/api.php"

//HTTPClient http client interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//NewService creates new fastbill api client
func NewService(email string, apiKey string) Service {
	client := &FastBillService{
		email:  email,
		apiKey: apiKey,
		client: &http.Client{},
	}

	return client
}

//NewServiceWithClient creates new fastbill api client
func NewServiceWithClient(email string, apiKey string, httpClient HTTPClient) Service {
	client := &FastBillService{
		email:  email,
		apiKey: apiKey,
		client: httpClient,
	}

	return client
}

//Service service interface
type Service interface {
	DoRequest(fastBillRequest request.Request) (response.Response, error)
}

//FastBillService the fastbill api client
type FastBillService struct {
	email  string
	apiKey string
	client HTTPClient
}

//DoRequest Executes the api call
func (c *FastBillService) DoRequest(fastBillRequest request.Request) (response.Response, error) {
	var fastBillResponse response.Response

	requestJSON, err := json.Marshal(fastBillRequest)
	if err != nil {
		return fastBillResponse, err
	}

	req, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(requestJSON))

	req.Header.Add("Content-Type", "application/json")

	basicHash := base64.StdEncoding.EncodeToString([]byte(c.email + ":" + c.apiKey))
	req.Header.Add("Authorization", "Basic "+basicHash)

	res, err := c.client.Do(req)
	if err != nil {
		return fastBillResponse, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &fastBillResponse); err != nil {
		return fastBillResponse, err
	}

	var errorResponse response.ErrorResponse
	err = mapstructure.Decode(fastBillResponse.Response, &errorResponse)
	if err == nil && len(errorResponse.Errors) > 0 {
		return response.Response{}, errors.New(strings.Join(errorResponse.Errors, ","))
	}

	return fastBillResponse, nil
}
