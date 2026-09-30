package project_test

import (
	"context"
	"errors"
	"testing"

	"github.com/malsch-solutions/fastbill-go-sdk/v2"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/internal/fastbilltest"
	"github.com/malsch-solutions/fastbill-go-sdk/v2/modules/project"
)

var ctx = context.Background()

// Projects as FastBill sends them: IDs and text as strings, rates as
// numbers or strings, no tasks as {}.
const projectsJSON = `{"PROJECTS":[
	{
		"PROJECT_ID": "21", "PROJECT_NAME": "Relaunch", "PROJECT_NUMBER": "P-021", "CUSTOMER_ID": "59",
		"CUSTOMER_COSTCENTER_ID": "0", "HOUR_PRICE": 95.5, "CURRENCY_CODE": "EUR", "VAT_PERCENT": "19.00",
		"START_DATE": "2026-09-01", "END_DATE": "0000-00-00",
		"TASKS": [{"TASK_ID": "7", "TASK_NUMBER": "1", "TASK_NAME": "Design", "STATUS": "open", "HOUR_PRICE": 80, "CURRENCY_CODE": "EUR", "VAT_PERCENT": 19}]
	},
	{"PROJECT_ID": "22", "PROJECT_NAME": "Support", "CUSTOMER_ID": "59", "TASKS": {}}
]}`

func TestGet(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return projectsJSON })
	projects, err := project.NewClient(server.Client()).Get(ctx, fastbill.Page{Limit: 100}, &project.Filter{CustomerID: "59"})
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 2 {
		t.Fatalf("projects = %+v", projects)
	}
	p := projects[0]
	if p.ProjectID != "21" || p.ProjectName != "Relaunch" || p.ProjectNumber != "P-021" || p.HourPrice != "95.5" || p.VatPercent != "19.00" || p.StartDate != "2026-09-01" {
		t.Errorf("project = %+v", p)
	}
	if len(p.Tasks) != 1 || p.Tasks[0].TaskID != "7" || p.Tasks[0].TaskName != "Design" || p.Tasks[0].HourPrice != "80" {
		t.Errorf("tasks = %+v", p.Tasks)
	}
	if len(projects[1].Tasks) != 0 {
		t.Errorf("empty tasks = %+v", projects[1].Tasks)
	}
	got := server.Last(t)
	if got.Service != "project.get" || got.Limit != 100 {
		t.Errorf("request = %+v", got)
	}
	fastbilltest.JSONEqual(t, got.Filter, `{"CUSTOMER_ID":"59"}`)
}

func TestCreate(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success","PROJECT_ID":23}` })
	res, err := project.NewClient(server.Client()).Create(ctx, &project.Project{
		ProjectName: "Wartung", CustomerID: "59", HourPrice: fastbill.NewNumber(95.5), VatPercent: fastbill.NewInt(19), StartDate: "2026-10-01",
	})
	if err != nil || res.ProjectID != "23" {
		t.Fatalf("Create = %+v, %v", res, err)
	}
	got := server.Last(t)
	if got.Service != "project.create" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"PROJECT_NAME":"Wartung","CUSTOMER_ID":"59","HOUR_PRICE":95.5,"VAT_PERCENT":19,"START_DATE":"2026-10-01"}`)
}

func TestUpdateAndDelete(t *testing.T) {
	server := fastbilltest.New(t, func(fastbilltest.Received) string { return `{"STATUS":"success"}` })
	c := project.NewClient(server.Client())

	if err := c.Update(ctx, &project.Project{ProjectID: "23", EndDate: "2026-12-31"}); err != nil {
		t.Error(err)
	}
	got := server.Last(t)
	if got.Service != "project.update" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"PROJECT_ID":"23","END_DATE":"2026-12-31"}`)

	if err := c.Delete(ctx, "23"); err != nil {
		t.Error(err)
	}
	got = server.Last(t)
	if got.Service != "project.delete" {
		t.Errorf("service = %s", got.Service)
	}
	fastbilltest.JSONEqual(t, got.Data, `{"PROJECT_ID":"23"}`)
}

func TestErrors(t *testing.T) {
	server := fastbilltest.New(t, func(r fastbilltest.Received) string {
		if r.Service == "project.delete" {
			return `{"STATUS":"failed"}`
		}
		return `{"ERRORS":["Project not found"]}`
	})
	c := project.NewClient(server.Client())
	var apiErr *fastbill.APIError
	if _, err := c.Get(ctx, fastbill.Page{}, nil); !errors.As(err, &apiErr) || apiErr.Messages[0] != "Project not found" {
		t.Errorf("Get: %v", err)
	}
	if err := c.Update(ctx, &project.Project{ProjectID: "23"}); !errors.As(err, &apiErr) {
		t.Errorf("Update: %v", err)
	}
	if err := c.Delete(ctx, "23"); !errors.As(err, &apiErr) {
		t.Errorf("Delete with failed status: %v", err)
	}
}
