package ports

import (
	"context"
	"image_storage_service/internal/core/domains"
)

type ImageInfoRepository interface {
	FindByID(ctx context.Context, id string) (imgInfo domains.ImageInfo, err error)
	Save(ctx context.Context, imgInfo domains.ImageInfo) (err error)
	Update(ctx context.Context, info domains.ImageInfo) (err error)
	Delete(ctx context.Context, id string) (err error)
}
