package time

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Filter narrows time.get. Dates are YYYY-MM-DD.
type Filter struct {
	CustomerID fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	ProjectID  fastbill.ID `json:"PROJECT_ID,omitempty"`
	TaskID     fastbill.ID `json:"TASK_ID,omitempty"`
	TimeID     fastbill.ID `json:"TIME_ID,omitempty"`
	StartDate  string      `json:"START_DATE,omitempty"`
	EndDate    string      `json:"END_DATE,omitempty"`
	Date       string      `json:"DATE,omitempty"`
}

// Time is a time entry as time.get returns it, and the data of time.create
// and time.update.
type Time struct {
	// TimeID is required for Update and ignored by Create.
	TimeID     fastbill.ID `json:"TIME_ID,omitempty"`
	TaskID     fastbill.ID `json:"TASK_ID,omitempty"`
	CustomerID fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	ProjectID  fastbill.ID `json:"PROJECT_ID,omitempty"`
	// InvoiceID is the invoice the time was billed with; Get only.
	InvoiceID fastbill.ID `json:"INVOICE_ID,omitempty"`
	// Date is YYYY-MM-DD.
	Date string `json:"DATE,omitempty"`
	// StartTime and EndTime are YYYY-MM-DD hh:mm:ss.
	StartTime       string          `json:"START_TIME,omitempty"`
	EndTime         string          `json:"END_TIME,omitempty"`
	Minutes         fastbill.Number `json:"MINUTES,omitempty"`
	BillableMinutes fastbill.Number `json:"BILLABLE_MINUTES,omitempty"`
	Comment         string          `json:"COMMENT,omitempty"`
}

// CreateResponse is the answer of time.create.
type CreateResponse struct {
	Status string      `json:"STATUS"`
	TimeID fastbill.ID `json:"TIME_ID"`
}

type idRequest struct {
	TimeID fastbill.ID `json:"TIME_ID"`
}

type getResponse struct {
	Times fastbill.List[Time] `json:"TIMES"`
}
