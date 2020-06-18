package data

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"mayaleng.org/engine/internal/platform/database"
	"mayaleng.org/engine/internal/platform/types"
)

type testInfo struct {
	envs *types.TestEnvs
	db   *mongo.Client
}

func setupTestInfo() (*testInfo, error) {
	envs := types.TestEnvs{}
	error := envconfig.Process("app", &envs)

	if error != nil {
		return nil, error
	}

	dbConfig := database.Config{
		StringConnection: envs.DatabaseConnection,
		Timeout:          time.Second,
	}

	dbClient, error := database.Open(dbConfig)

	if error != nil {
		return nil, error
	}

	result := testInfo{
		envs: &envs,
		db:   dbClient,
	}

	return &result, nil
}
