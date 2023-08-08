package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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

	logrus.Info(fmt.Sprintf("SERVER STARTED: %s", time.Now().Local().String()))
	go func() {
		if err := server.ListenAndServe(cfg.ServerPort); err != nil {
			logrus.WithFields(logrus.Fields{
				"handler": "main",
				"problem": "Error with a running server/or force closing server",
			}).Error(err)
			return
		}

	}()

	// Gracefull shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := db.Client().Disconnect(context.Background()); err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "main",
			"problem": "Error with closing DB connection",
		}).Error(err)
		return
	}

	logrus.Info("Server shutdown successfully")
}
