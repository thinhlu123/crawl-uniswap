package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"os"
	"time"
)

var ClientDB *mongo.Client
var PairCollection *mongo.Collection

func ConnectDB() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	ClientDB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = ClientDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	setupCollection()
	setupIndexes()

	fmt.Println("Connected to MongoDB!")
	return ClientDB, nil
}

func setupCollection() {
	PairCollection = ClientDB.Database("uniwap_db").Collection("pair")

}

func setupIndexes() {
	setPairIndexes()
}

func setPairIndexes() {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	indexModels := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "url", Value: bsonx.Int32(1)}},
			Options: &options.IndexOptions{
				Background: HelperPtrBool(true),
			},
		},
	}

	// Declare an options object
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := PairCollection.Indexes().CreateMany(ctx, indexModels, opts)

	// Check for the options errors
	if err != nil {
		fmt.Println("Indexes().CreateMany() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		fmt.Println("CreateMany() option:", opts)
	}
}

func HelperPtrBool(field bool) *bool {
	return &field
}

func ToBsonDoc(d interface{}) (bsonDoc bson.M, err error) {
	data, err := bson.Marshal(d)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &bsonDoc)
	return
}
