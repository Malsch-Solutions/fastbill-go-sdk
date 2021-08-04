package customer

//Filter available customer filter
type Filter struct {
	CustomerID     string `json:"CUSTOMER_ID"`     //A specific customer ID
	CustomerNumber string `json:"CUSTOMER_NUMBER"` //Assigned customer number
	CountryCode    string `json:"COUNTRY_CODE"`    //A country (ISO 3166 ALPHA-2)
	City           string `json:"CITY"`            //City
	Term           string `json:"TERM"`            //Search term in one of the given fields: ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS, ADDRESS_2, ZIPCODE, EMAIL, TAGS.
}

//Customer fastbill customer definition
type Customer struct {
	CustomerID                  string `json:"CUSTOMER_ID"`
	CustomerNumber              string `json:"CUSTOMER_NUMBER"`
	DaysForPayment              string `json:"DAYS_FOR_PAYMENT"`
	Created                     string `json:"CREATED"`
	PaymentType                 string `json:"PAYMENT_TYPE"`
	BankName                    string `json:"BANK_NAME"`
	BankAccountNumber           string `json:"BANK_ACCOUNT_NUMBER"`
	BankCode                    string `json:"BANK_CODE"`
	BankAccountOwner            string `json:"BANK_ACCOUNT_OWNER"`
	BankIban                    string `json:"BANK_IBAN"`
	BankBic                     string `json:"BANK_BIC"`
	BankAccountMandateReference string `json:"BANK_ACCOUNT_MANDATE_REFERENCE"`
	ShowPaymentNotice           string `json:"SHOW_PAYMENT_NOTICE"`
	CustomerAccount             string `json:"CUSTOMER_ACCOUNT"`
	CustomerType                string `json:"CUSTOMER_TYPE"`
	Top                         string `json:"TOP"`
	NewsletterOptIn             string `json:"NEWSLETTER_OPTIN"` //Deprecated
	Organization                string `json:"ORGANIZATION"`
	Position                    string `json:"POSITION"`
	AcademicDegreee             string `json:"ACADEMIC_DEGREE"`
	Salutation                  string `json:"SALUTATION"`
	FirstName                   string `json:"FIRST_NAME"`
	LastName                    string `json:"LAST_NAME"`
	Address                     string `json:"ADDRESS"`
	Address2                    string `json:"ADDRESS_2"`
	ZipCode                     string `json:"ZIPCODE"`
	City                        string `json:"CITY"`
	CountryCode                 string `json:"COUNTRY_CODE"`
	SecondaryAddress            string `json:"SECONDARY_ADDRESS"`
	Phone                       string `json:"PHONE"`
	Phone2                      string `json:"PHONE_2"`
	Fax                         string `json:"FAX"`
	Mobile                      string `json:"MOBILE"`
	Email                       string `json:"EMAIL"`
	Website                     string `json:"WEBSITE"`
	VatID                       string `json:"VAT_ID"`
	CurrencyCode                string `json:"CURRENCY_CODE"`
	LastUpdate                  string `json:"LASTUPDATE"`
	Tags                        string `json:"TAGS"`
	DocumentHistoryURL          string `json:"DOCUMENT_HISTORY_URL"`
}

type getResponse struct {
	Customers []Customer `json:"CUSTOMERS"`
}
