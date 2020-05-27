package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Config contains the values to be used to connect to the database
type Config struct {
	StringConnection string
	Timeout          time.Duration
}

// Open return a new connection to MongoDB
func Open(cfg Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.StringConnection)
	client, error := mongo.NewClient(clientOptions)

	if error != nil {
		return nil, error
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = time.Second * 10
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancelFunc()

	error = client.Connect(ctx)

	if error != nil {
		return nil, error
	}

	return client, error
}
