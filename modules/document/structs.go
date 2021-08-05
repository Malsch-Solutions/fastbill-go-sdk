package document

//Filter available invoice filter
type Filter struct {
	FolderID string `json:"FOLDER_ID"` // Folder ID
}

//Folder fastbill document folder
type Folder struct {
	FolderID       string `json:"FOLDER_ID" mapstructure:"FOLDER_ID"`
	Name           string `json:"NAME" mapstructure:"NAME"`
	ParentFolderID string `json:"PARENTFOLDER_ID" mapstructure:"PARENTFOLDER_ID"`
	Created        string `json:"CREATED" mapstructure:"CREATED"`
	ContentCount   string `json:"CONTENT_COUNT" mapstructure:"CONTENT_COUNT"`
}

//Document fastbill document
type Document struct {
	DocumentID string `json:"DOCUMENT_ID" mapstructure:"DOCUMENT_ID"`
	Type       string `json:"TYPE" mapstructure:"TYPE"`
	Title      string `json:"TITLE" mapstructure:"TITLE"`
	Date       string `json:"DATE" mapstructure:"DATE"`
	Note       string `json:"NOTE" mapstructure:"NOTE"`
}

type getResponse struct {
	Items Response `json:"ITEMS" mapstructure:"ITEMS"`
}

//Response fastbill document api get response
type Response struct {
	Folders   map[string]Folder `json:"FOLDERS" mapstructure:"FOLDERS"`
	Documents []Document        `json:"DOCUMENTS" mapstructure:"DOCUMENTS"`
}

//CreateResponse fastbill create response
type CreateResponse struct {
	DocumentID int    `json:"DOCUMENT_ID" mapstructure:"DOCUMENT_ID"`
	Status     string `json:"STATUS" mapstructure:"STATUS"`
}

type File struct {
}
