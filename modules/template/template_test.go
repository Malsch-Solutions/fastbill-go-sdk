package template_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/template"
)

var ctx = context.Background()

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string {
		return `{"TEMPLATES":[
			{"TEMPLATE_ID": "1", "TEMPLATE_NAME": "Standard", "TEMPLATE_HASH": "a1b2c3"},
			{"TEMPLATE_ID": 27, "TEMPLATE_NAME": "English", "TEMPLATE_HASH": "d4e5f6"}
		]}`
	})
	templates, err := template.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 10, Offset: 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(templates) != 2 {
		t.Fatalf("templates = %+v", templates)
	}
	if tpl := templates[1]; tpl.TemplateID != "27" || tpl.TemplateName != "English" || tpl.TemplateHash != "d4e5f6" {
		t.Errorf("template = %+v", tpl)
	}
	got := server.Last(t)
	if got.Service != "template.get" || got.Limit != 10 || got.Offset != 10 || got.Filter != nil {
		t.Errorf("request = %+v", got)
	}
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"TEMPLATES":{}}` })
	templates, err := template.NewClient(server.Client()).Get(ctx, fastbill.Page{})
	if err != nil || len(templates) != 0 {
		t.Errorf("Get = %+v, %v", templates, err)
	}
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ERRORS":["Unauthorized"]}` })
	var apiErr *fastbill.APIError
	if _, err := template.NewClient(server.Client()).Get(ctx, fastbill.Page{}); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Unauthorized" {
		t.Errorf("Get: %v", err)
	}
}
