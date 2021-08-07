package webhook

//Webhook fastbill webhook definition
type Webhook struct {
	WebhookID string `json:"WEBHOOK_ID,omitempty" mapstructure:"WEBHOOK_ID"`
	Endpoint  string `json:"ENDPOINT,omitempty" mapstructure:"ENDPOINT"`
	Type      string `json:"TYPE,omitempty" mapstructure:"TYPE"`
	Events    string `json:"EVENS,omitempty" mapstructure:"EVENS"` // customer.created,customer.updated,customer.deleted,invoice.created,invoice.completed,invoice.canceled,estimate.created,estimate.updated,contact.created,contact.updated,contact.deleted
}

//CreateResponse webhook api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	WebhookID int    `json:"WEBHOOK_ID" mapstructure:"WEBHOOK_ID"`
}

type getResponse struct {
	Webhooks []Webhook `json:"WEBHOOKS"`
}

type deleteRequest struct {
	WebhookID string `json:"WEBHOOK_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
