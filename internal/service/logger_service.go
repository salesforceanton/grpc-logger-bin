package service

import (
	"context"

	repository "github.com/salesforceanton/grpc-logger-bin/internal/repository/mongo"
	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
)

type LoggerService struct {
	repo repository.Logger
}

func NewLoggerService(repo repository.Logger) *LoggerService {
	return &LoggerService{repo: repo}
}

func (s *LoggerService) AddLog(ctx context.Context, req *loggerbin.LogRequest) error {
	item := loggerbin.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.AddLog(ctx, item)
}
