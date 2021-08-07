package article

//Filter available article filter
type Filter struct {
	ArticleNumber string `json:"ARTICLE_NUMBER"` //Assigned article number
}

//Article fastbill article definition
type T struct {
	ArticleID     string `json:"ARTICLE_ID" mapstructure:"ARTICLE_ID"`
	ArticleNumber string `json:"ARTICLE_NUMBER" mapstructure:"ARTICLE_NUMBER"`
	Title         string `json:"TITLE" mapstructure:"TITLE"`
	Description   string `json:"DESCRIPTION" mapstructure:"DESCRIPTION"`
	UNIT          string `json:"UNIT" mapstructure:"UNIT"`
	UNITPRICE     string `json:"UNIT_PRICE" mapstructure:"UNIT_PRICE"`
	CURRENCYCODE  string `json:"CURRENCY_CODE" mapstructure:"CURRENCY_CODE"`
	VATPERCENT    string `json:"VAT_PERCENT" mapstructure:"VAT_PERCENT"`
	ISGROSS       int    `json:"IS_GROSS" mapstructure:"IS_GROSS"`
	TAGS          string `json:"TAGS" mapstructure:"TAGS"`
}

//CreateResponse article api response
type CreateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ArticleID int    `json:"ARTICLE_ID" mapstructure:"ARTICLE_ID"`
}

//UpdateResponse article api response
type UpdateResponse struct {
	Status    string `json:"STATUS" mapstructure:"STATUS"`
	ArticleID string `json:"ARTICLE_ID" mapstructure:"ARTICLE_ID"`
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
