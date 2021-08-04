package request

import "github.com/malsch-solutions/fastbill-go-sdk/parameter"

//Request fastbill api request
type Request struct {
	Service string      `json:"SERVICE"`
	Limit   int        `json:"LIMIT,omitempty"`
	Offset  int        `json:"OFFSET,omitempty"`
	Filter  interface{} `json:"FILTER,omitempty"`
	Data    interface{} `json:"DATA,omitempty"`
}

//NewRequest creates new fastbill api request
func NewRequest(service string, parameter *parameter.Parameter, filter interface{}) Request {
	return Request{
		Service: service,
		Limit:   parameter.Limit,
		Filter:  filter,
		Offset:  parameter.Offset,
	}

}
