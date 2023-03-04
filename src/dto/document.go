package dto

// CreateDocumentRequest represents the request body for document creation
type CreateDocumentRequest struct {
    Filename string `json:"filename"`
    Data     []byte `json:"data"`
}