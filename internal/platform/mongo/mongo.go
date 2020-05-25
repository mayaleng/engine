package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Config contains the values to be used to connect to the database
type Config struct {
	StringConnection string
	Database         string
}

// Open return a new connection to MongoDB
func Open(cfg Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.StringConnection)
	client, error := mongo.NewClient(clientOptions)

	if error != nil {
		return nil, error
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	error = client.Connect(ctx)

	if error != nil {
		return nil, error
	}

	return client, error
}
