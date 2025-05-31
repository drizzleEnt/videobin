package filectrl

import (
	"videobin/internal/api"
	"videobin/internal/service"
)

var _ api.FileController = (*handler)(nil)

type handler struct {
	service service.FileService
}

func New(s service.FileService) *handler {
	return &handler{
		service: s,
	}
}
