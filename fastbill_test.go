package fastbill_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
)

var ctx = context.Background()

func TestDoSendsRequestAndDecodesResponse(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string {
		return `{"STATUS":"success","CUSTOMER_ID":42}`
	})
	var out struct {
		Status     string      `json:"STATUS"`
		CustomerID fastbill.ID `json:"CUSTOMER_ID"`
	}
	err := server.Client().Do(ctx, fastbill.DataRequest("customer.create", map[string]string{"ORGANIZATION": "ACME"}), &out)
	if err != nil {
		t.Fatal(err)
	}
	if out.Status != "success" || out.CustomerID != "42" {
		t.Errorf("out = %+v", out)
	}
	got := server.Last(t)
	if got.Service != "customer.create" || got.Header.Get("User-Agent") != "fastbill-go-sdk/v2" {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"ORGANIZATION":"ACME"}`)
}

func TestDoReturnsAPIErrors(t *testing.T) {
	for _, errors_ := range []string{`["Invalid API key"]`, `{"0":"Invalid API key"}`, `"Invalid API key"`} {
		server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ERRORS":` + errors_ + `}` })
		err := server.Client().Do(ctx, fastbill.Request{Service: "customer.get"}, nil)
		var apiErr *fastbill.APIError
		if !errors.As(err, &apiErr) || apiErr.Service != "customer.get" || len(apiErr.Messages) != 1 || apiErr.Messages[0] != "Invalid API key" {
			t.Errorf("%s: err = %v", errors_, err)
		}
	}
}

func TestDoReturnsStatusErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "maintenance", http.StatusServiceUnavailable)
	}))
	defer server.Close()
	err := fastbill.NewClient("a", "b", fastbill.WithBaseURL(server.URL)).Do(ctx, fastbill.Request{Service: "invoice.get"}, nil)
	var statusErr *fastbill.StatusError
	if !errors.As(err, &statusErr) || statusErr.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("err = %v", err)
	}
}

func TestDoMultipartSendsTheFile(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	req := fastbill.DataRequest("document.create", map[string]string{"TITLE": "Beleg"})
	if err := server.Client().DoMultipart(ctx, req, strings.NewReader("%PDF-1.7"), "beleg.pdf", nil); err != nil {
		t.Fatal(err)
	}
	got := server.Last(t)
	if got.Service != "document.create" || got.FileName != "beleg.pdf" || got.File != "%PDF-1.7" {
		t.Errorf("request = %+v", got)
	}
}

func TestGetRequestLeavesOutANilFilter(t *testing.T) {
	type filter struct {
		ID string `json:"ID,omitempty"`
	}
	b, _ := json.Marshal(fastbill.GetRequest[filter]("invoice.get", fastbill.Page{Limit: 5}, nil))
	if string(b) != `{"SERVICE":"invoice.get","LIMIT":5}` {
		t.Errorf("request = %s", b)
	}
	b, _ = json.Marshal(fastbill.GetRequest("invoice.get", fastbill.Page{}, &filter{ID: "1"}))
	if string(b) != `{"SERVICE":"invoice.get","FILTER":{"ID":"1"}}` {
		t.Errorf("request = %s", b)
	}
}

func TestNumber(t *testing.T) {
	var v struct {
		A, B, C, D fastbill.Number
	}
	if err := json.Unmarshal([]byte(`{"A":119.5,"B":"19.00","C":null,"D":""}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.A != "119.5" || v.B != "19.00" || v.C != "" || v.D != "" {
		t.Errorf("v = %+v", v)
	}
	if f, err := v.B.Float64(); err != nil || f != 19 {
		t.Errorf("Float64 = %v, %v", f, err)
	}
	out, _ := json.Marshal(struct {
		Price fastbill.Number `json:"PRICE"`
		Qty   fastbill.Number `json:"QTY,omitempty"`
		Odd   fastbill.Number `json:"ODD"`
	}{Price: fastbill.NewNumber(19.5), Odd: "n/a"})
	if string(out) != `{"PRICE":19.5,"ODD":"n/a"}` {
		t.Errorf("marshal = %s", out)
	}
	if fastbill.NewInt(3) != "3" {
		t.Error("NewInt")
	}
	if err := json.Unmarshal([]byte(`{"A":[1]}`), &v); err == nil {
		t.Error("a list decoded as Number")
	}
}

func TestID(t *testing.T) {
	var v struct{ A, B, C fastbill.ID }
	if err := json.Unmarshal([]byte(`{"A":"17","B":18,"C":null}`), &v); err != nil {
		t.Fatal(err)
	}
	if v.A != "17" || v.B != "18" || v.C != "" {
		t.Errorf("v = %+v", v)
	}
}

func TestFlag(t *testing.T) {
	var v struct{ A, B, C, D, E fastbill.Flag }
	if err := json.Unmarshal([]byte(`{"A":"1","B":0,"C":true,"D":false,"E":null}`), &v); err != nil {
		t.Fatal(err)
	}
	if !v.A.Bool() || v.B.Bool() || !v.C.Bool() || v.D != fastbill.No || v.E.Bool() {
		t.Errorf("v = %+v", v)
	}
	if b, _ := json.Marshal(fastbill.NewFlag(true)); string(b) != `"1"` {
		t.Errorf("marshal = %s", b)
	}
}

func TestList(t *testing.T) {
	for in, want := range map[string]string{
		`[{"N":"a"},{"N":"b"}]`:                        "a,b",
		`{"10":{"N":"c"},"2":{"N":"b"},"1":{"N":"a"}}`: "a,b,c",
		`{}`:   "",
		`[]`:   "",
		`null`: "",
	} {
		var l fastbill.List[struct{ N string }]
		if err := json.Unmarshal([]byte(in), &l); err != nil {
			t.Errorf("%s: %v", in, err)
			continue
		}
		var names []string
		for _, item := range l {
			names = append(names, item.N)
		}
		if got := strings.Join(names, ","); got != want {
			t.Errorf("%s: got %q, want %q", in, got, want)
		}
	}
}

func TestStatusResponse(t *testing.T) {
	if err := (fastbill.StatusResponse{Status: "success"}).Err("invoice.delete"); err != nil {
		t.Error(err)
	}
	if err := (fastbill.StatusResponse{Status: "failed"}).Err("invoice.delete"); err == nil {
		t.Error("failed status passed")
	}
}

func TestRecordsKeepsEveryField(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string {
		return `{"INVOICES":[{"INVOICE_ID":"1","NEW_FIELD":{"X":1}},{"INVOICE_ID":"2"}]}`
	})
	records, err := fastbill.Records(ctx, server.Client(), "invoice.get", "INVOICES", map[string]string{"TYPE": "draft"}, fastbill.Page{Limit: 100, Offset: 200})
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 2 || string(records[0]) != `{"INVOICE_ID":"1","NEW_FIELD":{"X":1}}` {
		t.Errorf("records = %s", records)
	}
	got := server.Last(t)
	if got.Limit != 100 || got.Offset != 200 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"TYPE":"draft"}`)
}

func TestRecordsOfAnEmptyListing(t *testing.T) {
	for _, response := range []string{`{"INVOICES":{}}`, `{"INVOICES":[]}`, `{}`, `[]`} {
		server := fastbilltest.New(t, func(fastbilltest.Received) string { return response })
		records, err := fastbill.Records(ctx, server.Client(), "invoice.get", "INVOICES", nil, fastbill.Page{})
		if err != nil || len(records) != 0 {
			t.Errorf("%s: %v, %v", response, records, err)
		}
	}
}

func TestAll(t *testing.T) {
	var offsets []int
	all, err := fastbill.All(ctx, 2, func(_ context.Context, p fastbill.Page) ([]int, error) {
		offsets = append(offsets, p.Offset)
		return []int{1, 2, 3, 4, 5}[p.Offset:min(p.Offset+p.Limit, 5)], nil
	})
	if err != nil || len(all) != 5 || len(offsets) != 3 {
		t.Errorf("all = %v, offsets = %v, err = %v", all, offsets, err)
	}
	if _, err := fastbill.All(ctx, 101, func(context.Context, fastbill.Page) ([]int, error) { return nil, nil }); err == nil {
		t.Error("page size above MaxLimit accepted")
	}
}

func TestDownload(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			t.Error("download sent credentials")
		}
		switch r.URL.Path {
		case "/invoice.pdf":
			_, _ = w.Write([]byte("%PDF-1.4 ..."))
		case "/login":
			_, _ = w.Write([]byte("<!DOCTYPE html><html><body>Login</body></html>"))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	client := fastbill.NewClient("a", "b", fastbill.WithHTTPClient(server.Client()))

	body, contentType, err := client.Download(ctx, server.URL+"/invoice.pdf?token=x")
	if err != nil || string(body) != "%PDF-1.4 ..." || contentType != "application/pdf" {
		t.Errorf("Download = %q, %q, %v", body, contentType, err)
	}
	if _, _, err := client.Download(ctx, server.URL+"/login"); !errors.Is(err, fastbill.ErrNotDocument) {
		t.Errorf("HTML page: err = %v", err)
	}
	var statusErr *fastbill.StatusError
	if _, _, err := client.Download(ctx, server.URL+"/gone"); !errors.As(err, &statusErr) {
		t.Errorf("404: err = %v", err)
	}
	if _, _, err := client.Download(ctx, "http://my.fastbill.com/x.pdf"); err == nil {
		t.Error("plain HTTP accepted")
	}
}
