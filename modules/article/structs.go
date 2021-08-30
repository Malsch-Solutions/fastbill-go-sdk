package article

//Filter available article filter
type Filter struct {
	ArticleNumber string `json:"ARTICLE_NUMBER,omitempty"` //Assigned article number
}

//Article fastbill article definition
type Article struct {
	ArticleID     string `json:"ARTICLE_ID,omitempty" mapstructure:"ARTICLE_ID"`
	ArticleNumber string `json:"ARTICLE_NUMBER,omitempty" mapstructure:"ARTICLE_NUMBER"`
	Title         string `json:"TITLE,omitempty" mapstructure:"TITLE"`
	Description   string `json:"DESCRIPTION,omitempty" mapstructure:"DESCRIPTION"`
	Unit          string `json:"UNIT,omitempty" mapstructure:"UNIT"`
	UnitPrice     string `json:"UNIT_PRICE,omitempty" mapstructure:"UNIT_PRICE"`
	CurrencyCode  string `json:"CURRENCY_CODE,omitempty" mapstructure:"CURRENCY_CODE"`
	VatPercent    string `json:"VAT_PERCENT,omitempty" mapstructure:"VAT_PERCENT"`
	IsGross       int    `json:"IS_GROSS,omitempty" mapstructure:"IS_GROSS"`
	Tags          string `json:"TAGS,omitempty" mapstructure:"TAGS"`
}

//CreateResponse article api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ArticleID int    `json:"ARTICLE_ID" mapstructure:"ARTICLE_ID"`
}

//UpdateResponse article api response
type UpdateResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}

type getResponse struct {
	Articles []Article `json:"ARTICLES"`
}

type deleteRequest struct {
	ArticleID string `json:"ARTICLE_ID"`
}

type deleteResponse struct {
	Status string `json:"STATUS" mapstructure:"STATUS"`
}
