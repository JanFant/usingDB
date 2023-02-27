package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"usingDB/internal/model/config"
)

func ConnectMongo(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb:"+config.GlobalConfig.MongoConfig.URL))
	if err != nil {
		return nil, err
	}
	return client, nil
}
