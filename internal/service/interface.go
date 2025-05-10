package service

import (
	"context"
	"mime/multipart"
)

type FileService interface {
	UploadFile(ctx context.Context, filename string, file multipart.File) error
}
