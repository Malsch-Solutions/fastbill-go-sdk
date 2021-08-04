package response

//Response fastbill api response
type Response struct {
	Request  interface{} `json:"REQUEST"`
	Response interface{} `json:"RESPONSE"`
	Errors   *[]string    `json:"errors"`
}
