// Package fastbilltest is a fake FastBill API for the SDK's tests.
package fastbilltest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
)

// Received is a request as the fake API saw it.
type Received struct {
	Service  string
	Limit    int
	Offset   int
	Filter   json.RawMessage
	Data     json.RawMessage
	FileName string
	File     string
	Header   http.Header
}

// Server answers every request with the JSON that respond returns as
// RESPONSE, e.g. `{"INVOICES":[{"INVOICE_ID":"1"}]}`.
type Server struct {
	*httptest.Server
	mu       sync.Mutex
	requests []Received
}

// New starts a fake API. It is closed when the test ends.
func New(t *testing.T, respond func(Received) string) *Server {
	t.Helper()
	s := &Server{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if email, key, ok := r.BasicAuth(); !ok || email != "user@example.com" || key != "api-key" {
			t.Errorf("request without credentials")
		}
		received := Received{Header: r.Header.Clone()}
		var body []byte
		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			if err := r.ParseMultipartForm(10 << 20); err != nil {
				t.Errorf("multipart: %v", err)
			}
			body = []byte(r.FormValue("httpbody"))
			if file, header, err := r.FormFile("document"); err == nil {
				content, _ := io.ReadAll(file)
				received.FileName, received.File = header.Filename, string(content)
			}
		} else {
			body, _ = io.ReadAll(r.Body)
		}
		var req struct {
			Service string          `json:"SERVICE"`
			Limit   int             `json:"LIMIT"`
			Offset  int             `json:"OFFSET"`
			Filter  json.RawMessage `json:"FILTER"`
			Data    json.RawMessage `json:"DATA"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("request body %q: %v", body, err)
		}
		received.Service, received.Limit, received.Offset = req.Service, req.Limit, req.Offset
		received.Filter, received.Data = req.Filter, req.Data
		s.mu.Lock()
		s.requests = append(s.requests, received)
		s.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"REQUEST":`+string(body)+`,"RESPONSE":`+respond(received)+`}`)
	}))
	t.Cleanup(s.Close)
	return s
}

// Client returns a client for the fake API.
func (s *Server) Client() *fastbill.Client {
	return fastbill.NewClient("user@example.com", "api-key", fastbill.WithBaseURL(s.URL), fastbill.WithHTTPClient(s.Server.Client()))
}

// Requests returns what the fake API received so far.
func (s *Server) Requests() []Received {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Received(nil), s.requests...)
}

// Last returns the last request, failing the test if there was none.
func (s *Server) Last(t *testing.T) Received {
	t.Helper()
	requests := s.Requests()
	if len(requests) == 0 {
		t.Fatal("no request received")
	}
	return requests[len(requests)-1]
}

// JSONEqual fails the test unless got and want hold the same JSON value.
func JSONEqual(t *testing.T, got json.RawMessage, want string) {
	t.Helper()
	var g, w any
	if err := json.Unmarshal(got, &g); err != nil {
		t.Fatalf("got invalid JSON %q: %v", got, err)
	}
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		t.Fatalf("want invalid JSON %q: %v", want, err)
	}
	gb, _ := json.Marshal(g)
	wb, _ := json.Marshal(w)
	if string(gb) != string(wb) {
		t.Errorf("JSON = %s\nwant   %s", gb, wb)
	}
}
