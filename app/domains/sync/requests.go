package sync

type FileMetadata struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
	MimeType string `json:"mime_type"`
}

type SyncRequest struct {
	Metadata []FileMetadata `json:"metadata"`
}
