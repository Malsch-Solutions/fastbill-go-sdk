package customer

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Types of customers (CUSTOMER_TYPE).
const (
	TypeBusiness = "business"
	TypeConsumer = "consumer"
)

// Filter narrows customer.get.
type Filter struct {
	CustomerID     fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	CustomerNumber string      `json:"CUSTOMER_NUMBER,omitempty"`
	// CountryCode is ISO 3166 alpha-2, e.g. DE.
	CountryCode string `json:"COUNTRY_CODE,omitempty"`
	City        string `json:"CITY,omitempty"`
	// Term searches ORGANIZATION, FIRST_NAME, LAST_NAME, ADDRESS,
	// ADDRESS_2, ZIPCODE, EMAIL and TAGS.
	Term string `json:"TERM,omitempty"`
}

// Customer is a customer as customer.get returns it, and the data of
// customer.create and customer.update.
type Customer struct {
	// CustomerID is required for Update and ignored by Create.
	CustomerID     fastbill.ID `json:"CUSTOMER_ID,omitempty"`
	CustomerNumber string      `json:"CUSTOMER_NUMBER,omitempty"`
	// CustomerType is business or consumer.
	CustomerType   string          `json:"CUSTOMER_TYPE,omitempty"`
	DaysForPayment fastbill.Number `json:"DAYS_FOR_PAYMENT,omitempty"`
	Created        string          `json:"CREATED,omitempty"`
	// PaymentType is 1 transfer, 2 direct debit, 3 cash, 4 PayPal,
	// 5 advance payment or 6 credit card. Direct debit requires BankName
	// and BankIBAN.
	PaymentType                 fastbill.Number `json:"PAYMENT_TYPE,omitempty"`
	BankName                    string          `json:"BANK_NAME,omitempty"`
	BankAccountNumber           string          `json:"BANK_ACCOUNT_NUMBER,omitempty"`
	BankCode                    string          `json:"BANK_CODE,omitempty"`
	BankAccountOwner            string          `json:"BANK_ACCOUNT_OWNER,omitempty"`
	BankIBAN                    string          `json:"BANK_IBAN,omitempty"`
	BankBIC                     string          `json:"BANK_BIC,omitempty"`
	BankAccountMandateReference string          `json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"`
	ShowPaymentNotice           fastbill.Flag   `json:"SHOW_PAYMENT_NOTICE,omitempty"`
	// AccountReceivable is the accounts receivable (debtor) number.
	AccountReceivable string `json:"ACCOUNT_RECEIVABLE,omitempty"`
	// Deprecated: FastBill documents the debtor number as
	// ACCOUNT_RECEIVABLE now; use AccountReceivable.
	CustomerAccount string `json:"CUSTOMER_ACCOUNT,omitempty"`
	// Top marks a top customer. It is not in the current FastBill docs.
	Top fastbill.Flag `json:"TOP,omitempty"`
	// Deprecated: FastBill no longer supports the newsletter option.
	NewsletterOptIn fastbill.Flag `json:"NEWSLETTER_OPTIN,omitempty"`
	Organization    string        `json:"ORGANIZATION,omitempty"`
	Position        string        `json:"POSITION,omitempty"`
	AcademicDegree  string        `json:"ACADEMIC_DEGREE,omitempty"`
	// Salutation is mr, mrs, family or empty.
	Salutation  string `json:"SALUTATION,omitempty"`
	FirstName   string `json:"FIRST_NAME,omitempty"`
	LastName    string `json:"LAST_NAME,omitempty"`
	Address     string `json:"ADDRESS,omitempty"`
	Address2    string `json:"ADDRESS_2,omitempty"`
	ZipCode     string `json:"ZIPCODE,omitempty"`
	City        string `json:"CITY,omitempty"`
	CountryCode string `json:"COUNTRY_CODE,omitempty"`
	// SecondaryAddress is the delivery address.
	SecondaryAddress string `json:"SECONDARY_ADDRESS,omitempty"`
	Phone            string `json:"PHONE,omitempty"`
	Phone2           string `json:"PHONE_2,omitempty"`
	Fax              string `json:"FAX,omitempty"`
	Mobile           string `json:"MOBILE,omitempty"`
	Email            string `json:"EMAIL,omitempty"`
	Website          string `json:"WEBSITE,omitempty"`
	VatID            string `json:"VAT_ID,omitempty"`
	CurrencyCode     string `json:"CURRENCY_CODE,omitempty"`
	// BuyerReference is the buyer reference (Leitweg-ID) for e-invoices.
	BuyerReference string `json:"BUYER_REFERENCE,omitempty"`
	// GLN is the GS1 Global Location Number.
	GLN        string `json:"GLN,omitempty"`
	LastUpdate string `json:"LASTUPDATE,omitempty"`
	Tags       string `json:"TAGS,omitempty"`
	// DocumentHistoryURL links to the customer's documents in the
	// customer center.
	DocumentHistoryURL string `json:"DOCUMENT_HISTORY_URL,omitempty"`
}

// CreateResponse is the answer of customer.create.
type CreateResponse struct {
	Status         string      `json:"STATUS"`
	CustomerID     fastbill.ID `json:"CUSTOMER_ID"`
	CustomerNumber string      `json:"CUSTOMER_NUMBER"`
}

type idRequest struct {
	CustomerID fastbill.ID `json:"CUSTOMER_ID"`
}

type getResponse struct {
	Customers fastbill.List[Customer] `json:"CUSTOMERS"`
}
