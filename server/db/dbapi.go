package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const CollectionPerson string = "person"         // collection person
const CollectionRecord string = "record"         // collection record
const dbname string = "treeman"                  // database name
const dbuser string = "root"                     // database username
const dbpw string = "88888888"                   // database password
const dburl string = "mongodb://localhost:27017" // url for local mongodb
//const dburl string = "mongodb://mongo:27017"     //url for mongodb in docker

var client *mongo.Client

func Connect() error {
	var err error
	cdt := options.Credential{Username: dbuser, Password: dbpw}
	client, err = mongo.NewClient(options.Client().ApplyURI(dburl).SetAuth(cdt))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func Insert(document *interface{}, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database(dbname).Collection(collectionName)
	_, err := collection.InsertOne(ctx, &document)
	if err != nil {
		return err
	}
	return nil
}
