package service

import (
	"context"

	repository "github.com/salesforceanton/grpc-logger-bin/internal/repository/mongo"
	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
)

type Service struct {
	Logger
}

type Logger interface {
	AddLog(ctx context.Context, req *loggerbin.LogRequest) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Logger: NewLoggerService(repos.Logger),
	}
}
