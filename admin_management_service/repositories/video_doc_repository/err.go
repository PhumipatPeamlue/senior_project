package video_doc_repository

import "errors"

var (
	ErrVideoDocNotFound     = errors.New("video document repository: video document not found")
	ErrInternalVideoDocRepo = errors.New("video document repository: internal repository error")
)
