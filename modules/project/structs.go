package project

// Filter available project filter
type Filter struct {
	ProjectID  string `json:"PROJECT_ID,omitempty"`  //project id
	CustomerID string `json:"CUSTOMER_ID,omitempty"` //customer id
}

// Project fastbill project definition
type Project struct {
	ProjectID            string `json:"PROJECT_ID,omitempty" mapstructure:"PROJECT_ID"`
	ProjectName          string `json:"PROJECT_NAME,omitempty" mapstructure:"PROJECT_NAME"`
	ProjectNumber        string `json:"PROJECT_NUMBER,omitempty" mapstructure:"PROJECT_NUMBER"`
	CustomerID           string `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerCostCenterID string `json:"CUSTOMER_COSTCENTER_ID,omitempty" mapstructure:"CUSTOMER_COSTCENTER_ID"`
	HourPrice            string `json:"HOUR_PRICE,omitempty" mapstructure:"HOUR_PRICE"`
	CurrencyCode         string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	VatPercent           string `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	StartDate            string `json:"START_DATE,omitempty" mapstructure:"START_DATE"`
	EndDate              string `json:"END_DATE,omitempty" mapstructure:"END_DATE"`
	Tasks                []Task `json:"TASKS,omitempty" mapstructure:"TASKS"`
}

// Task fastbill project task definition
type Task struct {
	TaskID       string `json:"TASK_ID,omitempty" mapstructure:"TASK_ID"`
	TaskNumber   string `json:"TASK_NUMBER,omitempty" mapstructure:"TASK_NUMBER"`
	TaskName     string `json:"TASK_NAME,omitempty" mapstructure:"TASK_NAME"`
	Description  string `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Status       string `json:"STATUS,omitempty" mapstructure:"STATUS"`
	Priority     string `json:"PRIORITY,omitempty" mapstructure:"PRIORITY"`
	HourPrice    string `json:"HOUR_PRICE,omitempty" mapstructure:"HOUR_PRICE"`
	CurrencyCode string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	VatPercent   int    `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
}

// CreateResponse project api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ProjectID int    `json:"PROJECT_ID" mapstructure:"PROJECT_ID"`
}

// UpdateResponse project api response
type UpdateResponse struct {
	ProjectID string `json:"PROJECT_ID" mapstructure:"PROJECT_ID"`
	Status    string `json:"STATUS" mapstructure:"STATUS"`
}

type getResponse struct {
	Projects []Project `json:"PROJECTS"`
}

type deleteRequest struct {
	ProjectID string `json:"PROJECT_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
