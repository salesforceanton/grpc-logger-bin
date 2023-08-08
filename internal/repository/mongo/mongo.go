package repository

import (
	"context"

	"github.com/salesforceanton/grpc-logger-bin/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(cfg *config.Config) (*mongo.Database, error) {
	// Set params and URI
	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DBUsername,
		Password: cfg.DBPassword,
	})
	opts.ApplyURI(cfg.DBUri)

	// Connect to DB and check connection with ping func
	dbClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	db := dbClient.Database(cfg.DBName)
	return db, nil
}
