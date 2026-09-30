package article

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Types of articles (TYPE).
const (
	TypeProduct        = "product"
	TypeDigitalProduct = "digital_product"
	TypeService        = "service"
	TypeNone           = "none"
)

// Filter narrows article.get.
type Filter struct {
	ArticleID     fastbill.ID `json:"ARTICLE_ID,omitempty"`
	ArticleNumber string      `json:"ARTICLE_NUMBER,omitempty"`
}

// Article is a product or service as article.get returns it, and the data
// of article.create and article.update.
type Article struct {
	// ArticleID is required for Update and ignored by Create.
	ArticleID     fastbill.ID `json:"ARTICLE_ID,omitempty"`
	ArticleNumber string      `json:"ARTICLE_NUMBER,omitempty"`
	// Type is product, digital_product, service or none.
	Type         string          `json:"TYPE,omitempty"`
	Title        string          `json:"TITLE,omitempty"`
	Description  string          `json:"DESCRIPTION,omitempty"`
	Unit         string          `json:"UNIT,omitempty"`
	UnitPrice    fastbill.Number `json:"UNIT_PRICE,omitempty"`
	CurrencyCode string          `json:"CURRENCY_CODE,omitempty"`
	VatPercent   fastbill.Number `json:"VAT_PERCENT,omitempty"`
	// IsGross (1) means UnitPrice includes VAT.
	IsGross fastbill.Flag `json:"IS_GROSS,omitempty"`
	Tags    string        `json:"TAGS,omitempty"`
}

// CreateResponse is the answer of article.create.
type CreateResponse struct {
	Status    string      `json:"STATUS"`
	ArticleID fastbill.ID `json:"ARTICLE_ID"`
}

type idRequest struct {
	ArticleID fastbill.ID `json:"ARTICLE_ID"`
}

type getResponse struct {
	Articles fastbill.List[Article] `json:"ARTICLES"`
}
