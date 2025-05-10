package service

import (
	"context"
	"mime/multipart"
)

var _ FileService = (*fileService)(nil)

type fileService struct {
}

func New() *fileService {
	return &fileService{}
}

// UploadFile implements FileService.
func (f *fileService) UploadFile(ctx context.Context, filename string, file multipart.File) error {
	return nil
}
