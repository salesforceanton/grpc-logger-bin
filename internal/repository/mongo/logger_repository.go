package repository

import (
	"context"

	loggerbin "github.com/salesforceanton/grpc-logger-bin/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const LOGS_COLLECTION_NAME = "logs"

type LoggerRepo struct {
	db *mongo.Database
}

func NewLoggerRepo(db *mongo.Database) *LoggerRepo {
	return &LoggerRepo{db: db}
}

func (r *LoggerRepo) AddLog(ctx context.Context, item loggerbin.LogItem) error {
	_, err := r.db.Collection(LOGS_COLLECTION_NAME).InsertOne(ctx, item)

	return err
}
