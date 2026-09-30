package article_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/article"
)

var ctx = context.Background()

// An article as FastBill sends it: IDs and text as strings, prices as
// numbers, the gross flag as a number.
const articleJSON = `{
	"ARTICLE_ID": "305", "ARTICLE_NUMBER": "SUP-01", "TYPE": "service", "TITLE": "Support",
	"DESCRIPTION": "Support per hour", "UNIT": "Std.", "UNIT_PRICE": 82.5, "CURRENCY_CODE": "EUR",
	"VAT_PERCENT": "19.00", "IS_GROSS": 0, "TAGS": ""
}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ARTICLES":[` + articleJSON + `]}` })
	articles, err := article.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 20, Offset: 40}, &article.Filter{ArticleNumber: "SUP-01"})
	if err != nil {
		t.Fatal(err)
	}
	if len(articles) != 1 {
		t.Fatalf("articles = %+v", articles)
	}
	a := articles[0]
	if a.ArticleID != "305" || a.ArticleNumber != "SUP-01" || a.Type != article.TypeService || a.UnitPrice != "82.5" || a.VatPercent != "19.00" || a.IsGross != fastbill.No {
		t.Errorf("article = %+v", a)
	}
	got := server.Last(t)
	if got.Service != "article.get" || got.Limit != 20 || got.Offset != 40 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"ARTICLE_NUMBER":"SUP-01"}`)
}

func TestGetEmpty(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"ARTICLES":{}}` })
	articles, err := article.NewClient(server.Client()).Get(ctx, fastbill.Page{}, &article.Filter{ArticleID: "999"})
	if err != nil || len(articles) != 0 {
		t.Fatalf("Get = %+v, %v", articles, err)
	}
	fastbilltest.JSONEqual(t, server.Last(t).Filter, `{"ARTICLE_ID":"999"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","ARTICLE_ID":306}` })
	res, err := article.NewClient(server.Client()).Create(ctx, &article.Article{
		ArticleNumber: "HOST-01", Type: article.TypeService, Title: "Hosting", Unit: "Monat",
		UnitPrice: fastbill.NewNumber(49.9), VatPercent: fastbill.NewInt(19), IsGross: fastbill.Yes,
	})
	if err != nil || res.ArticleID != "306" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "article.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{
		"ARTICLE_NUMBER": "HOST-01", "TYPE": "service", "TITLE": "Hosting", "UNIT": "Monat",
		"UNIT_PRICE": 49.9, "VAT_PERCENT": 19, "IS_GROSS": "1"
	}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	c := article.NewClient(server.Client())

	if err := c.Update(ctx, &article.Article{ArticleID: "306", UnitPrice: fastbill.NewNumber(59.9)}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "article.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"ARTICLE_ID":"306","UNIT_PRICE":59.9}`)

	if err := c.Delete(ctx, "306"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "article.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"ARTICLE_ID":"306"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "article.create" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Article not found"]}`
	})
	c := article.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Article not found" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Delete(ctx, "306"); !errors.As(err, &apiErr) {
		t.Errorf("Delete: %v", err)
	}
	if _, err := c.Create(ctx, &article.Article{Title: "x"}); !errors.As(err, &apiErr) {
		t.Errorf("Create with failed status: %v", err)
	}
}
