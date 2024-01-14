package video_doc_service

import "errors"

var (
	ErrVideoDocNotFound        = errors.New("video document service: video document not found")
	ErrInternalVideoDocService = errors.New("video document service: internal service error")
)
