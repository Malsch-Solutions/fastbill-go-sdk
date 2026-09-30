package template

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Template is a template as template.get returns it.
type Template struct {
	// TemplateID changes whenever the template is edited; use TemplateHash
	// to recognize a template over time.
	TemplateID   fastbill.ID `json:"TEMPLATE_ID"`
	TemplateName string      `json:"TEMPLATE_NAME"`
	// TemplateHash is the permanent unique ID of the template.
	TemplateHash string `json:"TEMPLATE_HASH"`
}

type getResponse struct {
	Templates fastbill.List[Template] `json:"TEMPLATES"`
}
