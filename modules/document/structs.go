package document

import "github.com/malsch-solutions/fastbill-go-sdk/v2"

// Filter narrows document.get.
type Filter struct {
	FolderID fastbill.ID `json:"FOLDER_ID,omitempty"`
}

// GetResponse is the answer of document.get.
type GetResponse struct {
	Folders   fastbill.List[Folder]   `json:"FOLDERS"`
	Documents fastbill.List[Document] `json:"DOCUMENTS"`
}

// Folder is a folder of the document inbox.
type Folder struct {
	FolderID       fastbill.ID `json:"FOLDER_ID"`
	Name           string      `json:"NAME"`
	ParentFolderID fastbill.ID `json:"PARENTFOLDER_ID"`
	Created        string      `json:"CREATED"`
	// ContentCount is the number of items in the folder.
	ContentCount fastbill.Number `json:"CONTENT_COUNT"`
}

// Document is a document as document.get returns it.
type Document struct {
	DocumentID fastbill.ID `json:"DOCUMENT_ID"`
	Type       string      `json:"TYPE"`
	Title      string      `json:"TITLE"`
	// Date is YYYY-MM-DD.
	Date string `json:"DATE"`
	Note string `json:"NOTE"`
}

// Request is the data of document.create.
type Request struct {
	Type  string `json:"TYPE,omitempty"`
	Title string `json:"TITLE,omitempty"`
	// Date is YYYY-MM-DD.
	Date string `json:"DATE,omitempty"`
	Note string `json:"NOTE,omitempty"`
}

// CreateResponse is the answer of document.create.
type CreateResponse struct {
	Status     string      `json:"STATUS"`
	DocumentID fastbill.ID `json:"DOCUMENT_ID"`
}

// getResponse takes FOLDERS and DOCUMENTS directly in RESPONSE, as the
// docs show, or wrapped in ITEMS, as the API has been seen to send them.
type getResponse struct {
	GetResponse
	Items *GetResponse `json:"ITEMS"`
}
