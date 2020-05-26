package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"log"
	"mayaleng.org/engine/cmd/http/internal/handlers"
	"mayaleng.org/engine/internal/envs"
	"mayaleng.org/engine/internal/platform/mongo"
	"mayaleng.org/engine/version"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func all(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("%s %s", r.Method, r.URL)
	w.WriteHeader(404)
	w.Write([]byte("Not found"))
}

func main() {
	var envs envs.ENVs
	envError := envconfig.Process("app", &envs)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	log.Printf("Version %s built at: %s", version.BuildNumber, version.BuildTime)

	if envError != nil {
		logrus.Fatal(envError)
	}

	if envs.ENV == "dev" {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
	}

	logrus.Info("Initializing database connection")

	database, mongoError := mongo.Open(mongo.Config{
		StringConnection: envs.DatabaseConnection,
	})

	if mongoError != nil {
		logrus.Fatal(mongoError)
	}

	defer func() {
		logrus.Info("Closing database connection")
		database.Disconnect(context.Background())
	}()

	logrus.Info("Initializing API")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	api := http.Server{
		Addr:    envs.Host,
		Handler: handlers.NewAPI(database),
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
