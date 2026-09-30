package project

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Filter narrows project.get.
type Filter struct {
	ProjectID  fastbill.ID `json:"PROJECT_ID,omitempty"`
	CustomerID fastbill.ID `json:"CUSTOMER_ID,omitempty"`
}

// Project is a project as project.get returns it, and the data of
// project.create and project.update.
type Project struct {
	// ProjectID is required for Update and ignored by Create.
	ProjectID   fastbill.ID `json:"PROJECT_ID,omitempty"`
	ProjectName string      `json:"PROJECT_NAME,omitempty"`
	// ProjectNumber is not in the current FastBill docs.
	ProjectNumber        string          `json:"PROJECT_NUMBER,omitempty"`
	CustomerID           fastbill.ID     `json:"CUSTOMER_ID,omitempty"`
	CustomerCostCenterID fastbill.ID     `json:"CUSTOMER_COSTCENTER_ID,omitempty"`
	HourPrice            fastbill.Number `json:"HOUR_PRICE,omitempty"`
	CurrencyCode         string          `json:"CURRENCY_CODE,omitempty"`
	VatPercent           fastbill.Number `json:"VAT_PERCENT,omitempty"`
	// StartDate and EndDate are YYYY-MM-DD.
	StartDate string `json:"START_DATE,omitempty"`
	EndDate   string `json:"END_DATE,omitempty"`
	// Tasks are returned by Get; Create and Update do not take them.
	Tasks fastbill.List[Task] `json:"TASKS,omitempty"`
}

// Task is a task of a project.
type Task struct {
	TaskID       fastbill.ID     `json:"TASK_ID"`
	TaskNumber   string          `json:"TASK_NUMBER"`
	TaskName     string          `json:"TASK_NAME"`
	Description  string          `json:"DESCRIPTION"`
	Status       string          `json:"STATUS"`
	Priority     string          `json:"PRIORITY"`
	HourPrice    fastbill.Number `json:"HOUR_PRICE"`
	CurrencyCode string          `json:"CURRENCY_CODE"`
	VatPercent   fastbill.Number `json:"VAT_PERCENT"`
}

// CreateResponse is the answer of project.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	ProjectID fastbill.ID `json:"PROJECT_ID"`
}

type idRequest struct {
	ProjectID fastbill.ID `json:"PROJECT_ID"`
}

type getResponse struct {
	Projects fastbill.List[Project] `json:"PROJECTS"`
}
