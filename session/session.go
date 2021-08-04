package session

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"

	"github.com/malsch-solutions/fastbill-go-sdk/request"
	"github.com/malsch-solutions/fastbill-go-sdk/response"
)

const baseURL string = "https://my.fastbill.com/api/1.0/api.php"

//NewSession creates new fastbill api client
func NewSession(email string, apiKey string) Session {
	client := &FastBillSession{
		email:  email,
		apiKey: apiKey,
		client: &http.Client{},
	}

	return client
}

//Session session interface
type Session interface {
	DoRequest(fastBillRequest request.Request) (response.Response, error)
}

//FastBillSession the fastbill api client
type FastBillSession struct {
	email  string
	apiKey string
	client *http.Client
}

//DoRequest Executes the api call
func (c *FastBillSession) DoRequest(fastBillRequest request.Request) (response.Response, error) {
	var fastBillResponse response.Response

	requestJSON, err := json.Marshal(fastBillRequest)

	if err != nil {
		return fastBillResponse, err
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return fastBillResponse, err
	}

	req.Header.Add("Content-Type", "application/json")

	basicHash := base64.StdEncoding.EncodeToString([]byte(c.email + ":" + c.apiKey))
	req.Header.Add("Authorization", "Basic "+basicHash)

	res, err := c.client.Do(req)
	if err != nil {
		return fastBillResponse, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println(err)
		}
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
