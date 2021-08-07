package time

//Filter available time filter
type Filter struct {
	CustomerID string `json:"CUSTOMER_ID"`
	ProjectID  string `json:"PROJECT_ID"`
	TaskID     string `json:"TASK_ID"`
	TimeID     string `json:"TIME_ID"`
	StartDate  string `json:"START_DATE"`
	EndDate    string `json:"END_DATE"`
	Date       string `json:"DATE"`
}

//Time fastbill time
type Time struct {
	TimeID          string `json:"TIME_ID" mapstructure:"TIME_ID"`
	TaskID          string `json:"TASK_ID" mapstructure:"TASK_ID"`
	CustomerID      string `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
	ProjectID       string `json:"PROJECT_ID" mapstructure:"PROJECT_ID"`
	Date            string `json:"DATE" mapstructure:"DATE"`
	StartTime       string `json:"START_TIME" mapstructure:"START_TIME"`
	EndTime         string `json:"END_TIME" mapstructure:"END_TIME"`
	Minutes         string `json:"MINUTES" mapstructure:"MINUTES"`
	BillableMinutes string `json:"BILLABLE_MINUTES" mapstructure:"BILLABLE_MINUTES"`
	Comment         string `json:"COMMENT" mapstructure:"COMMENT"`
}

//CreateResponse time api response
type CreateResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
	TimeID int    `json:"TIME_ID" mapstructure:"TIME_ID"`
}

//UpdateResponse time api response
type UpdateResponse struct {
	TimeID string `json:"TIME_ID" mapstructure:"TIME_ID"`
	Status string `json:"STATUS" mapstructure:"STATUS"`
}

type getResponse struct {
	Times []Time `json:"TIMES"`
}

type deleteRequest struct {
	TimeID string `json:"TIME_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
