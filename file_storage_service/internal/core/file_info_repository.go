package core

import (
	"context"
	"file_storage_service/internal/core/domains"
)

type FileInfoRepositoryInterface interface {
	ReadByID(ctx context.Context, id string) (info domains.FileInfo, err error)
	Create(ctx context.Context, info domains.FileInfo) (err error)
	Update(ctx context.Context, info domains.FileInfo) (err error)
	Delete(ctx context.Context, id string) (err error)
}
