package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/malsch-solutions/fastbill-go-sdk/modules/request"
	"github.com/malsch-solutions/fastbill-go-sdk/modules/response"
)

const baseURL string = "https://my.fastbill.com/api/1.0/api.php"

// HTTPClient http client interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewService creates new fastbill api client
func NewService(email string, apiKey string) Service {
	client := &FastBillService{
		email:  email,
		apiKey: apiKey,
		client: &http.Client{},
	}

	return client
}

// NewServiceWithClient creates new fastbill api client
func NewServiceWithClient(email string, apiKey string, httpClient HTTPClient) Service {
	client := &FastBillService{
		email:  email,
		apiKey: apiKey,
		client: httpClient,
	}

	return client
}

// Service service interface
type Service interface {
	DoRequest(fastBillRequest request.Request) (response.Response, error)
	DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error)
}

// FastBillService the fastbill api client
type FastBillService struct {
	email  string
	apiKey string
	client HTTPClient
}

// DoMultiPartRequest Executes the api call as with file upload
func (c *FastBillService) DoMultiPartRequest(fastBillRequest request.Request, file io.Reader, fileName string) (response.Response, error) {
	var fastBillResponse response.Response

	requestJSON, err := json.Marshal(fastBillRequest)
	if err != nil {
		return fastBillResponse, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	_ = writer.SetBoundary("AaB03x")

	bodyWriter, _ := writer.CreateFormField("httpbody")
	_, _ = bodyWriter.Write(requestJSON)

	fileWriter, _ := writer.CreateFormFile("document", fileName)
	_, _ = io.Copy(fileWriter, file)

	_ = writer.Close()

	req := getRequest(body, c.email, c.apiKey, writer.FormDataContentType())
	req.Header.Add("Content-Length", strconv.Itoa(body.Len()))
	res, err := c.client.Do(req)
	if err != nil {
		return fastBillResponse, err
	}

	err = parseResponse(res, &fastBillResponse)
	return fastBillResponse, err
}

// DoRequest Executes the api call
func (c *FastBillService) DoRequest(fastBillRequest request.Request) (response.Response, error) {
	var fastBillResponse response.Response

	requestJSON, err := json.Marshal(fastBillRequest)
	if err != nil {
		return fastBillResponse, err
	}

	req := getRequest(bytes.NewBuffer(requestJSON), c.email, c.apiKey, "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return fastBillResponse, err
	}

	err = parseResponse(res, &fastBillResponse)
	return fastBillResponse, err
}

func getRequest(body io.Reader, email string, apiKey string, contentType string) *http.Request {
	req, _ := http.NewRequest("POST", baseURL, body)

	req.Header.Add("Content-Type", contentType)

	basicHash := base64.StdEncoding.EncodeToString([]byte(email + ":" + apiKey))
	req.Header.Add("Authorization", "Basic "+basicHash)

	return req
}

func parseResponse(res *http.Response, fastBillResponse *response.Response) error {
	defer func() {
		_ = res.Body.Close()
	}()

	body, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &fastBillResponse); err != nil {
		return err
	}

	var errorResponse response.ErrorResponse
	err := mapstructure.Decode(fastBillResponse.Response, &errorResponse)
	if err == nil && len(errorResponse.Errors) > 0 {
		return errors.New(strings.Join(errorResponse.Errors, ","))
	}

	return nil
}
