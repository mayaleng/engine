package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"mayaleng.org/engine/cmd/httpd/internal/handlers"
	"mayaleng.org/engine/internal/platform/database"
	"mayaleng.org/engine/internal/platform/types"
	"mayaleng.org/engine/version"
)

func main() {
	var envs types.Envs
	envError := envconfig.Process("app", &envs)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	logrus.WithFields(logrus.Fields{
		"version": version.BuildNumber,
		"date":    version.BuildTime,
	}).Info("Build information")

	if envError != nil {
		logrus.Fatal(envError)
	}

	if envs.Env == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// Database setup

	logrus.Info("Initializing database connection")

	database, mongoError := database.Open(database.Config{
		StringConnection: envs.DatabaseConnection,
	})

	if mongoError != nil {
		logrus.Fatal(mongoError)
	}

	defer func() {
		logrus.Info("Closing database connection")
		database.Disconnect(context.Background())
	}()

	// API setup

	logrus.Info("Initializing API")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:    envs.Host,
		Handler: handlers.NewAPI(&envs, database),
	}

	serverErrors := make(chan error, 1)

	go func() {
		logrus.Infof("Listening at http://%s", envs.Host)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case error := <-serverErrors:
		logrus.Fatal(error)
	case signal := <-shutdown:
		logrus.Infof("Exiting the app with %s", signal.String())
	}
}
