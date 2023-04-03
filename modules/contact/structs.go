package contact

// Filter available customer filter
type Filter struct {
	ContactID      string `json:"CONTACT_ID,omitempty"`      //A specific contact ID
	CustomerID     string `json:"CUSTOMER_ID,omitempty"`     //A specific customer ID
	CustomerNumber string `json:"CUSTOMER_NUMBER,omitempty"` //Assigned customer number
	Term           string `json:"TERM,omitempty"`            //Search term in one of the given fields: ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS, ADDRESS_2, ZIPCODE, EMAIL, TAGS.
}

// Contact fastbill contact
type Contact struct {
	ContactID        string `json:"CONTACT_ID,omitempty" mapstructure:"CONTACT_ID"`
	CustomerID       string `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	Organization     string `json:"ORGANIZATION,omitempty" mapstructure:"ORGANIZATION"`
	Position         string `json:"POSITION,omitempty" mapstructure:"POSITION"`
	AcademicDegree   string `json:"ACADEMIC_DEGREE,omitempty" mapstructure:"ACADEMIC_DEGREE"`
	Salutation       string `json:"SALUTATION,omitempty" mapstructure:"SALUTATION"`
	FirstName        string `json:"FIRST_NAME,omitempty" mapstructure:"FIRST_NAME"`
	LastName         string `json:"LAST_NAME,omitempty" mapstructure:"LAST_NAME"`
	Address          string `json:"ADDRESS,omitempty" mapstructure:"ADDRESS"`
	Address2         string `json:"ADDRESS_2,omitempty" mapstructure:"ADDRESS_2"`
	ZipCode          string `json:"ZIPCODE,omitempty" mapstructure:"ZIPCODE"`
	City             string `json:"CITY,omitempty" mapstructure:"CITY"`
	CountryCode      string `json:"COUNTRY_CODE,omitempty" mapstructure:"COUNTRY_CODE"`
	SecondaryAddress string `json:"SECONDARY_ADDRESS,omitempty" mapstructure:"SECONDARY_ADDRESS"`
	Phone            string `json:"PHONE,omitempty" mapstructure:"PHONE"`
	Phone2           string `json:"PHONE_2,omitempty" mapstructure:"PHONE_2"`
	Fax              string `json:"FAX,omitempty" mapstructure:"FAX"`
	Mobile           string `json:"MOBILE,omitempty" mapstructure:"MOBILE"`
	Email            string `json:"EMAIL,omitempty" mapstructure:"EMAIL"`
	Website          string `json:"WEBSITE,omitempty" mapstructure:"WEBSITE"`
	VatID            string `json:"VAT_ID,omitempty" mapstructure:"VAT_ID"`
	CurrencyCode     string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	Comment          string `json:"COMMENT,omitempty" mapstructure:"COMMENT"`
	CreatedAt        string `json:"CREATED,omitempty" mapstructure:"CREATED"`
	LastUpdate       string `json:"LASTUPDATE,omitempty" mapstructure:"LASTUPDATE"`
	Tags             string `json:"TAGS,omitempty" mapstructure:"TAGS"`
}

// CreateResponse customer api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ContactID int    `json:"CONTACT_ID" mapstructure:"CONTACT_ID"`
}

// UpdateResponse customer api response
type UpdateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ContactID string `json:"CONTACT_ID" mapstructure:"CONTACT_ID"`
}

type getResponse struct {
	Contacts []Contact `json:"CONTACTS"`
}

type deleteRequest struct {
	ContactID  string `json:"CONTACT_ID"`
	CustomerID string `json:"CUSTOMER_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
