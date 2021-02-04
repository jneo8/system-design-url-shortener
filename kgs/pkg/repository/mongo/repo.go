package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repo struct {
	Config        Config
	Client        *mongo.Client
	Logger        *log.Logger
	KeyCollection string
}

func (r *repo) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (r *repo) Init() error {
	if err := r.createIndexes(); err != nil {
		return err
	}
	return nil
}

func (r *repo) createIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mod := mongo.IndexModel{
		Keys: bson.M{
			"key": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := r.keyCollection().Indexes().CreateOne(ctx, mod)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) KeyBatchInsert(keys []string) (int, error) {
	models := []mongo.WriteModel{}
	for _, k := range keys {
		models = append(
			models,
			mongo.NewUpdateOneModel().SetFilter(
				bson.D{
					primitive.E{
						Key:   "key",
						Value: k,
					},
				},
			).SetUpdate(
				bson.D{
					primitive.E{
						Key: "$set",
						Value: bson.M{
							"used":     false,
							"expireAt": 0,
						},
					},
				},
			).SetUpsert(true),
		)
	}
	opts := options.BulkWrite().SetOrdered(false)
	res, err := r.keyCollection().BulkWrite(context.TODO(), models, opts)
	if err != nil {
		return 0, err
	}
	return int(res.UpsertedCount), nil
}

func (r *repo) keyCollection() *mongo.Collection {
	return r.Client.Database(r.Config.DB).Collection(r.Config.KeyCollection)
}
