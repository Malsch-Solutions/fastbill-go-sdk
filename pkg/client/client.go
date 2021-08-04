package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/malsch-solutions/fastbill-go-sdk/pkg/request"
	"github.com/malsch-solutions/fastbill-go-sdk/pkg/response"
)

const baseURL string = "https://my.fastbill.com/api/1.0/api.php"

//NewClient creates new fastbill api client
func NewClient(email string, apiKey string) *Client {
	client := &Client{
		email:  email,
		apiKey: apiKey,
		client: &http.Client{},
	}

	return client
}

//Client the fastbill api client
type Client struct {
	email  string
	apiKey string
	client *http.Client
}

//DoRequest Executes the api call
func (c Client) DoRequest(fastBillRequest request.Request) (response.Response, error) {
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

	return fastBillResponse, nil
}
