package fastbill

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultBaseURL is the FastBill API endpoint.
const DefaultBaseURL = "https://my.fastbill.com/api/1.0/api.php"

// MaxLimit is the most records FastBill returns per call.
const MaxLimit = 100

const (
	defaultTimeout  = 60 * time.Second
	maxResponseSize = 20 << 20
	maxDocumentSize = 100 << 20
	userAgent       = "fastbill-go-sdk/v2"
)

// Requester sends requests to the FastBill API. *Client implements it; the
// module clients take it so tests can replace it.
type Requester interface {
	// Do sends req and decodes the RESPONSE object into out (if not nil).
	Do(ctx context.Context, req Request, out any) error
	// DoMultipart sends req together with a file, for services that
	// upload documents.
	DoMultipart(ctx context.Context, req Request, file io.Reader, fileName string, out any) error
}

// Client talks to the FastBill API with the account's email and API key.
type Client struct {
	baseURL string
	email   string
	apiKey  string
	http    *http.Client
}

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient replaces the HTTP client (default: 60 s timeout).
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.http = h } }

// WithBaseURL replaces the API endpoint, e.g. for tests.
func WithBaseURL(u string) Option { return func(c *Client) { c.baseURL = u } }

// NewClient returns a client that authenticates as email with apiKey.
func NewClient(email, apiKey string, opts ...Option) *Client {
	c := &Client{
		baseURL: DefaultBaseURL,
		email:   email,
		apiKey:  apiKey,
		http:    &http.Client{Timeout: defaultTimeout},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Request is the body FastBill expects: the service name plus a filter
// (reading) or data (writing).
type Request struct {
	Service string `json:"SERVICE"`
	Limit   int    `json:"LIMIT,omitempty"`
	Offset  int    `json:"OFFSET,omitempty"`
	Filter  any    `json:"FILTER,omitempty"`
	Data    any    `json:"DATA,omitempty"`
}

// Page selects a slice of a listing. FastBill returns 10 records without a
// limit and at most MaxLimit.
type Page struct {
	Limit  int
	Offset int
}

// GetRequest builds a request for a `*.get` service. A nil filter is left
// out instead of being sent as null.
func GetRequest[F any](service string, page Page, filter *F) Request {
	req := Request{Service: service, Limit: page.Limit, Offset: page.Offset}
	if filter != nil {
		req.Filter = filter
	}
	return req
}

// DataRequest builds a request that sends data (create, update, …).
func DataRequest(service string, data any) Request {
	return Request{Service: service, Data: data}
}

// APIError is the ERRORS list FastBill returns for a failed call.
type APIError struct {
	Service  string
	Messages []string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("fastbill %s: %s", e.Service, strings.Join(e.Messages, "; "))
}

// StatusError is a non-2xx HTTP answer.
type StatusError struct {
	StatusCode int
	Body       string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("fastbill: HTTP %d: %s", e.StatusCode, e.Body)
}

// Do sends req and decodes the RESPONSE object into out.
func (c *Client) Do(ctx context.Context, req Request, out any) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("fastbill %s: encode request: %w", req.Service, err)
	}
	return c.send(ctx, req.Service, bytes.NewReader(payload), "application/json", out)
}

// DoMultipart sends req as the part "httpbody" and the file as the part
// "document", which is how FastBill takes uploads.
func (c *Client) DoMultipart(ctx context.Context, req Request, file io.Reader, fileName string, out any) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("fastbill %s: encode request: %w", req.Service, err)
	}
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	part, err := w.CreateFormField("httpbody")
	if err != nil {
		return err
	}
	if _, err := part.Write(payload); err != nil {
		return err
	}
	part, err = w.CreateFormFile("document", fileName)
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("fastbill %s: read file: %w", req.Service, err)
	}
	if err := w.Close(); err != nil {
		return err
	}
	return c.send(ctx, req.Service, &body, w.FormDataContentType(), out)
}

func (c *Client) send(ctx context.Context, service string, body io.Reader, contentType string, out any) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, body)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", contentType)
	httpReq.Header.Set("User-Agent", userAgent)
	httpReq.SetBasicAuth(c.email, c.apiKey)

	raw, _, err := c.fetch(httpReq, maxResponseSize)
	if err != nil {
		return fmt.Errorf("fastbill %s: %w", service, err)
	}
	var envelope struct {
		Response json.RawMessage `json:"RESPONSE"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return fmt.Errorf("fastbill %s: decode response: %w", service, err)
	}
	var failed struct {
		Errors List[string] `json:"ERRORS"`
	}
	// RESPONSE is not always an object (e.g. an empty list); only an
	// object can carry ERRORS.
	if bytes.HasPrefix(bytes.TrimSpace(envelope.Response), []byte("{")) {
		if err := json.Unmarshal(envelope.Response, &failed); err == nil && len(failed.Errors) > 0 {
			return &APIError{Service: service, Messages: failed.Errors}
		}
	}
	// PHP encodes an empty RESPONSE as [].
	if trimmed := bytes.TrimSpace(envelope.Response); out == nil || len(trimmed) == 0 || string(trimmed) == "[]" || string(trimmed) == "null" {
		return nil
	}
	if err := json.Unmarshal(envelope.Response, out); err != nil {
		return fmt.Errorf("fastbill %s: decode response: %w", service, err)
	}
	return nil
}

// ErrNotDocument means a download returned a web page instead of a file,
// e.g. a sign-in or error page.
var ErrNotDocument = errors.New("fastbill: download returned an HTML page, not a document")

// Download fetches a DOCUMENT_URL (invoice PDF, receipt file). The URL
// carries its own access token, so no credentials are sent with it. It
// returns the body and its detected content type.
func (c *Client) Download(ctx context.Context, documentURL string) ([]byte, string, error) {
	u, err := url.Parse(documentURL)
	if err != nil || u.Scheme != "https" || u.Host == "" {
		return nil, "", fmt.Errorf("fastbill: not a document URL: %q", documentURL)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, documentURL, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("User-Agent", userAgent)
	body, header, err := c.fetch(req, maxDocumentSize)
	if err != nil {
		return nil, "", fmt.Errorf("fastbill: download %s: %w", u.Path, err)
	}
	if len(body) == 0 {
		return nil, "", fmt.Errorf("fastbill: download %s: empty body", u.Path)
	}
	contentType := http.DetectContentType(body)
	if strings.HasPrefix(contentType, "text/html") {
		return nil, "", ErrNotDocument
	}
	if contentType == "application/octet-stream" && header.Get("Content-Type") != "" {
		contentType = header.Get("Content-Type")
	}
	return body, contentType, nil
}

func (c *Client) fetch(req *http.Request, limit int64) ([]byte, http.Header, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, limit+1))
	if err != nil {
		return nil, nil, err
	}
	if int64(len(body)) > limit {
		return nil, nil, fmt.Errorf("response larger than %d bytes", limit)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		text := string(body)
		if len(text) > 300 {
			text = text[:300] + "…"
		}
		return nil, nil, &StatusError{StatusCode: resp.StatusCode, Body: text}
	}
	return body, resp.Header, nil
}
