package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func Insert(document interface{}, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database(dbname).Collection(collectionName)
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}

func Update(filter, update *bson.M, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database(dbname).Collection(collectionName)
	_, err := collection.UpdateOne(ctx, *filter, *update)
	if err != nil {
		return err
	}
	return nil
}

func Find(results *[]bson.M, filter *bson.M, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database(dbname).Collection(collectionName)
	cur, err := collection.Find(ctx, *filter)
	defer cur.Close(context.Background())
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return err
		}
		*results = append(*results, result)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}

func Delete(filter *bson.M, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database(dbname).Collection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
