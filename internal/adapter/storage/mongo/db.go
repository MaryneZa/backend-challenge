package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
)

func ConnectMongoDB(config *config.DB) (*mongo.Client, *mongo.Database, error) {

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.User, config.Password, config.Host, config.Port)
	log.Println("url: " + url)
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	db := client.Database(config.Name)

	return client, db, nil
}

func UserCountLogging(db *mongo.Database) {
	ch := time.Tick(10 * time.Second)
	collection := db.Collection("users")
	go func() {
		for range ch {

			count, err := collection.CountDocuments(context.TODO(), bson.D{})
			if err != nil {
				panic(err)
			}

			log.Println("Number of Users in the DB =>", count)
		}
	}()
}

