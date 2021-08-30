package customer

//Filter available customer filter
type Filter struct {
	CustomerID     string `json:"CUSTOMER_ID,omitempty"`     //A specific customer ID
	CustomerNumber string `json:"CUSTOMER_NUMBER,omitempty"` //Assigned customer number
	CountryCode    string `json:"COUNTRY_CODE,omitempty"`    //A country (ISO 3166 ALPHA-2)
	City           string `json:"CITY,omitempty"`            //City
	Term           string `json:"TERM,omitempty"`            //Search term in one of the given fields: ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS, ADDRESS_2, ZIPCODE, EMAIL, TAGS.
}

//Customer fastbill customer definition
type Customer struct {
	CustomerID                  string `json:"CUSTOMER_ID,omitempty" mapstructure:"CUSTOMER_ID"`
	CustomerNumber              string `json:"CUSTOMER_NUMBER,omitempty" mapstructure:"CUSTOMER_NUMBER"`
	DaysForPayment              string `json:"DAYS_FOR_PAYMENT,omitempty" mapstructure:"DAYS_FOR_PAYMENT"`
	Created                     string `json:"CREATED,omitempty" mapstructure:"CREATED"`
	PaymentType                 string `json:"PAYMENT_TYPE,omitempty" mapstructure:"PAYMENT_TYPE"`
	BankName                    string `json:"BANK_NAME,omitempty" mapstructure:"BANK_NAME"`
	BankAccountNumber           string `json:"BANK_ACCOUNT_NUMBER,omitempty" mapstructure:"BANK_ACCOUNT_NUMBER"`
	BankCode                    string `json:"BANK_CODE,omitempty" mapstructure:"BANK_CODE"`
	BankAccountOwner            string `json:"BANK_ACCOUNT_OWNER,omitempty" mapstructure:"BANK_ACCOUNT_OWNER"`
	BankIban                    string `json:"BANK_IBAN,omitempty" mapstructure:"BANK_IBAN"`
	BankBic                     string `json:"BANK_BIC,omitempty" mapstructure:"BANK_BIC"`
	BankAccountMandateReference string `json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" mapstructure:"BANK_ACCOUNT_MANDATE_REFERENCE"`
	ShowPaymentNotice           string `json:"SHOW_PAYMENT_NOTICE,omitempty" mapstructure:"SHOW_PAYMENT_NOTICE"`
	CustomerAccount             string `json:"CUSTOMER_ACCOUNT,omitempty" mapstructure:"CUSTOMER_ACCOUNT"`
	CustomerType                string `json:"CUSTOMER_TYPE,omitempty" mapstructure:"CUSTOMER_TYPE"`
	Top                         string `json:"TOP,omitempty" mapstructure:"TOP"`
	NewsletterOptIn             string `json:"NEWSLETTER_OPTIN,omitempty" mapstructure:"NEWSLETTER_OPTIN"` //Deprecated
	Organization                string `json:"ORGANIZATION,omitempty" mapstructure:"ORGANIZATION"`
	Position                    string `json:"POSITION,omitempty" mapstructure:"POSITION"`
	AcademicDegreee             string `json:"ACADEMIC_DEGREE,omitempty" mapstructure:"ACADEMIC_DEGREE"`
	Salutation                  string `json:"SALUTATION,omitempty" mapstructure:"SALUTATION"`
	FirstName                   string `json:"FIRST_NAME,omitempty" mapstructure:"FIRST_NAME"`
	LastName                    string `json:"LAST_NAME,omitempty" mapstructure:"LAST_NAME"`
	Address                     string `json:"ADDRESS,omitempty" mapstructure:"ADDRESS"`
	Address2                    string `json:"ADDRESS_2,omitempty" mapstructure:"ADDRESS_2"`
	ZipCode                     string `json:"ZIPCODE,omitempty" mapstructure:"ZIPCODE"`
	City                        string `json:"CITY,omitempty" mapstructure:"CITY"`
	CountryCode                 string `json:"COUNTRY_CODE,omitempty" mapstructure:"COUNTRY_CODE"`
	SecondaryAddress            string `json:"SECONDARY_ADDRESS,omitempty" mapstructure:"SECONDARY_ADDRESS"`
	Phone                       string `json:"PHONE,omitempty" mapstructure:"PHONE"`
	Phone2                      string `json:"PHONE_2,omitempty" mapstructure:"PHONE_2"`
	Fax                         string `json:"FAX,omitempty" mapstructure:"FAX"`
	Mobile                      string `json:"MOBILE,omitempty" mapstructure:"MOBILE"`
	Email                       string `json:"EMAIL,omitempty" mapstructure:"EMAIL"`
	Website                     string `json:"WEBSITE,omitempty" mapstructure:"WEBSITE"`
	VatID                       string `json:"VAT_ID,omitempty" mapstructure:"VAT_ID"`
	CurrencyCode                string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	LastUpdate                  string `json:"LASTUPDATE,omitempty" mapstructure:"LASTUPDATE"`
	Tags                        string `json:"TAGS,omitempty" mapstructure:"TAGS"`
	DocumentHistoryURL          string `json:"DOCUMENT_HISTORY_URL,omitempty" mapstructure:"DOCUMENT_HISTORY_URL"`
}

//CreateResponse customer api response
type CreateResponse struct {
	Status     string `json:"STATUS" mapstructure:"STATUS"`
	CustomerID int    `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
}

//UpdateResponse customer api response
type UpdateResponse struct {
	Status     string `json:"STATUS" mapstructure:"STATUS"`
	CustomerID string `json:"CUSTOMER_ID" mapstructure:"CUSTOMER_ID"`
}

type getResponse struct {
	Customers []Customer `json:"CUSTOMERS"`
}

type deleteRequest struct {
	CustomerID string `json:"CUSTOMER_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
