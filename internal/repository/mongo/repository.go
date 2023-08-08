package repository

import (
	"context"

	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Logger
}

type Logger interface {
	AddLog(ctx context.Context, item loggerbin.LogItem) error
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Logger: NewLoggerRepo(db),
	}
}
