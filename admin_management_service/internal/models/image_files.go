package models

type ImageFile struct {
	ID       int    `json:"id"`
	DocID    string `json:"doc_id"`
	Filename string `json:"filename"`
	FilePath string `json:"file_path"`
}

type ImageFiles struct {
	Data []ImageFile `json:"data"`
}
