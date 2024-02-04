package ports

import (
	"context"
	"mime/multipart"
)

type ImageStorageService interface {
	GetURL(ctx context.Context, id string) (imageUrl string, err error)
	Save(ctx context.Context, id string, file *multipart.File, header *multipart.FileHeader) (err error)
	Delete(ctx context.Context, id string) (err error)
}
