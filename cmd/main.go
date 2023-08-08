package main

import (
	"fmt"
	"time"

	"github.com/salesforceanton/grpc-logger-bin/internal/config"
	hanler "github.com/salesforceanton/grpc-logger-bin/internal/handler"
	repository "github.com/salesforceanton/grpc-logger-bin/internal/repository/mongo"
	"github.com/salesforceanton/grpc-logger-bin/internal/server"
	"github.com/salesforceanton/grpc-logger-bin/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	// Config initialization
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "main",
			"problem": fmt.Sprintf("Error with config initialization: %s", err.Error()),
		}).Error(err)
		return
	}

	// Connect to DB
	db, err := repository.NewMongoDB(cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "main",
			"problem": fmt.Sprintf("Error with connect to Mongo DB: %s", err.Error()),
		}).Error(err)
		return
	}

	// Set dependenties
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := hanler.NewHandler(service)

	server := server.New(handler)

	fmt.Println("SERVER STARTED", time.Now())

	go func() {
		if err := server.ListenAndServe(cfg.ServerPort); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "main",
				"problem": "Error with a running server/or force closing server",
			}).Error(err)
			return
		}

	}()

	logrus.Info(fmt.Sprintf("SERVER STARTED: %s", time.Now().Local().String()))

}
