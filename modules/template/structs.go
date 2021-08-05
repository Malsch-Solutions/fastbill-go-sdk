package template

//Filter available customer filter
type Filter struct {
	ContactID      string `json:"CONTACT_ID"`      //A specific contact ID
	CustomerID     string `json:"CUSTOMER_ID"`     //A specific customer ID
	CustomerNumber string `json:"CUSTOMER_NUMBER"` //Assigned customer number
	Term           string `json:"TERM"`            //Search term in one of the given fields: ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS, ADDRESS_2, ZIPCODE, EMAIL, TAGS.
}

//Template fastbill template
type Template struct {
	TemplateID   int    `json:"TEMPLATE_ID" mapstructure:"TEMPLATE_ID"`
	TemplateName string `json:"TEMPLATE_NAME" mapstructure:"TEMPLATE_NAME"`
	TemplateHash string `json:"TEMPLATE_HASH" mapstructure:"TEMPLATE_HASH"`
}

type getResponse struct {
	Templates []Template `json:"TEMPLATES"`
}
