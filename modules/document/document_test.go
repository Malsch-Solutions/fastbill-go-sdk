package document_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/document"
)

var ctx = context.Background()

// Folders keyed by folder ID (PHP), documents as a list, IDs as strings.
const listingJSON = `{
	"FOLDERS": {
		"12": {"FOLDER_ID": "12", "NAME": "Receipts", "PARENTFOLDER_ID": "0", "CREATED": "2026-01-05 10:00:00", "CONTENT_COUNT": "3"},
		"4": {"FOLDER_ID": "4", "NAME": "Contracts", "PARENTFOLDER_ID": "0", "CREATED": "2025-11-02 08:30:00", "CONTENT_COUNT": 0}
	},
	"DOCUMENTS": [
		{"DOCUMENT_ID": "881", "TYPE": "receipt", "TITLE": "Hosting 09/2026", "DATE": "2026-09-01", "NOTE": ""}
	]
}`

func TestGet(t *testing.T) {
	for name, response := range map[string]string{
		"as documented":    listingJSON,
		"wrapped in ITEMS": `{"ITEMS":` + listingJSON + `}`,
	} {
		t.Run(name, func(t *testing.T) {
			server := fastbilltest.New(t, func(fastbilltest.Received) string { return response })
			res, err := document.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 50}, &document.Filter{FolderID: "12"})
			if err != nil {
				t.Fatal(err)
			}
			if len(res.Folders) != 2 || res.Folders[0].FolderID != "4" || res.Folders[1].Name != "Receipts" || res.Folders[1].ContentCount != "3" || res.Folders[0].ContentCount != "0" {
				t.Errorf("folders = %+v", res.Folders)
			}
			if len(res.Documents) != 1 {
				t.Fatalf("documents = %+v", res.Documents)
			}
			if doc := res.Documents[0]; doc.DocumentID != "881" || doc.Type != "receipt" || doc.Title != "Hosting 09/2026" || doc.Date != "2026-09-01" {
				t.Errorf("document = %+v", doc)
			}
			got := server.Last(t)
			if got.Service != "document.get" || got.Limit != 50 {
				t.Errorf("request = %+v", got)
			}
			fastbilltest.JSONEqual(t, got.Filter, `{"FOLDER_ID":"12"}`)
		})
	}
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"FOLDERS":{},"DOCUMENTS":{}}` })
	res, err := document.NewClient(server.Client()).Get(ctx, fastbill.Page{}, nil)
	if err != nil || len(res.Folders) != 0 || len(res.Documents) != 0 {
		t.Errorf("Get = %+v, %v", res, err)
	}
	if got := server.Last(t); got.Filter != nil {
		t.Errorf("filter = %s, want none", got.Filter)
	}
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","DOCUMENT_ID":882}` })
	res, err := document.NewClient(server.Client()).Create(ctx, &document.Request{
		Type: "receipt", Title: "Hosting 10/2026", Date: "2026-10-01",
	}, strings.NewReader("%PDF-1.7"), "hosting.pdf")
	if err != nil || res.DocumentID != "882" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "document.create" || got.FileName != "hosting.pdf" || got.File != "%PDF-1.7" {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"TYPE":"receipt","TITLE":"Hosting 10/2026","DATE":"2026-10-01"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "document.create" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Folder not found"]}`
	})
	c := document.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, &document.Filter{FolderID: "99"}); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Folder not found" {
		t.Errorf("Get: %v", err)
	}
	if _, err := c.Create(ctx, &document.Request{}, strings.NewReader("x"), "x.txt"); !errors.As(err, &apiErr) {
		t.Errorf("Create with failed status: %v", err)
	}
}
