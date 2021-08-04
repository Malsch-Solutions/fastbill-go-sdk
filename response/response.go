package response

//Response fastbill api response
type Response struct {
	Request  interface{} `json:"REQUEST"`
	Response interface{} `json:"RESPONSE"`
}

type ErrorResponse struct {
	Errors []string `json:"ERRORS"`
}
