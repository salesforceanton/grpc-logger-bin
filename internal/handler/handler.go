package handler

import (
	"context"

	"github.com/salesforceanton/grpc-logger-bin/internal/service"
	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
)

type Handler struct {
	service *service.Service
	loggerbin.UnimplementedLoggerbinServiceServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddLog(ctx context.Context, req *loggerbin.LogRequest) (*loggerbin.Empty, error) {
	err := h.service.Logger.AddLog(ctx, req)

	return &loggerbin.Empty{}, err
}
