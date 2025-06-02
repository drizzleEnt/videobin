package filesrv

import (
	"context"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"videobin/internal/service"
)

var _ service.FileService = (*fileService)(nil)

type fileService struct {
}

func New() *fileService {
	return &fileService{}
}

// UploadFile implements FileService.
func (f *fileService) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) error {
	ext := filepath.Ext(fileHeader.Filename)

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		file, err := fileHeader.Open()
		if err != nil {
			return fmt.Errorf("service.UploadFile %w", err)
		}
		defer file.Close()

		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			return fmt.Errorf("service.UploadFile %w", err)
		}

		mimeType = http.DetectContentType(buffer)
	}

	fmt.Printf("mimeType: %v\n", mimeType)

	return nil
}
