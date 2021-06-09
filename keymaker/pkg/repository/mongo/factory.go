package mongo

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/keymaker/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/dig"
)

// Config for repo.
type Config struct {
	dig.In
	DSN           string `name:"mongo_dsn"`
	KeyCollection string `name:"mongo_key_collection"`
	DB            string `name:"mongo_db"`
}

// New is a simple factory func for KeyRepository.
func New(logger *log.Logger, config Config) (entity.KeyRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DSN))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &repo{Logger: logger, Config: config, Client: client}, nil
}
