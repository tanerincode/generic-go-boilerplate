package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"tanerincode/generic-go-boilerplate/interval/storage"
)

type Mongo struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

const batchSize = 1500

func New() (storage.Storage, error) {
	credential := options.Credential{
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CON_URI")).SetAuth(credential))
	if err != nil {
		log.Printf("could not start app, %v", err)
		return nil, err
	}

	database := client.Database(os.Getenv("DATABASE_NAME"))
	collection := database.Collection(os.Getenv("COLLECTION_NAME"))

	if err != nil {
		return nil, err
	}

	return &Mongo{
		Client:     client,
		Database:   database,
		Collection: collection,
	}, nil
}

func (m Mongo) CleanUp() error {
	err := m.Collection.Drop(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
func (m Mongo) Disconnect() error {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
