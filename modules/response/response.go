package response

//Response fastbill api response
type Response struct {
	Request  interface{} `json:"REQUEST"`
	Response interface{} `json:"RESPONSE"`
}

//ErrorResponse api error response
type ErrorResponse struct {
	Errors []string `json:"ERRORS"`
}
