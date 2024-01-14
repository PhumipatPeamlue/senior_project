package video_doc_service

import "mime/multipart"

type VideoDocService interface {
	GetVideoDoc(docID string) (getRes *GetResponse, err error)
	CreateVideoDoc(newDocReq *NewVideoDocRequest) (err error)
	CreateVideoDocWithImage(newDocReq *NewVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	SearchVideoDoc(page int, pageSize int, keyword string) (searchRes *SearchResponse, err error)
	UpdateVideoDoc(updateDocReq *UpdateVideoDocRequest) (err error)
	UpdateVideoDocWithImage(updateDocReq *UpdateVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	DeleteVideoDoc(docID string) (err error)
}
