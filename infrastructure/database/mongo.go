package database

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongo(address string, databaseName string) *mongo.Database {
	connection, err := Open(address)
	if err != nil {
		log.Fatal("Fatal error", err)
	}

	db := connection.Database(databaseName)
	log.Println("Database loaded")

	return db
}

func Open(uri string) (*mongo.Client, error) {
	conf := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, conf)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "timeout: could not connect to mongo")
	}

	return client, err
}
