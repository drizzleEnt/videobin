package service

import (
	"context"
	"mime/multipart"
)

type FileService interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) error
}
