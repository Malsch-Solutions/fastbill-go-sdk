package request

import "github.com/malsch-solutions/fastbill-go-sdk/modules/parameter"

// Request fastbill api request
type Request struct {
	Service string      `json:"SERVICE"`
	Limit   int         `json:"LIMIT,omitempty"`
	Offset  int         `json:"OFFSET,omitempty"`
	Filter  interface{} `json:"FILTER,omitempty"`
	Data    interface{} `json:"DATA,omitempty"`
}

// NewRequestWithFilters creates new fastbill api request with filters, usual used for get requests
func NewRequestWithFilters(service string, parameter *parameter.Parameter, filter interface{}) Request {
	return Request{
		Service: service,
		Limit:   parameter.Limit,
		Filter:  filter,
		Offset:  parameter.Offset,
	}
}

// NewRequestWithData creates new fastbill api request with data, usual used for create or update requests
func NewRequestWithData(service string, data interface{}) Request {
	return Request{
		Service: service,
		Data:    data,
	}
}
