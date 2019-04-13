package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const CollectionPerson string = "person" // collection person
const CollectionRecord string = "record" // collection record
const dbname string = "treeman"          // database name
const dbuser string = "root"             // database username
const dbpw string = "88888888"           // database password
var dburls []string = []string{
	"mongodb://localhost:27017",
	"mongodb://mongo:27017",
}

var client *mongo.Client

func SmartConnect(c chan error) {
	var err error
	for i := 0; i < len(dburls); i++ {
		log.Println("connect to db url", dburls[i])
		err = Connect(dburls[i])
		if err != nil {
			c <- err
			return
		}
		for j := 0; j < 10; j++ {
			err = Ping()
			if err != nil {
				log.Println("ping fail", j)
				time.Sleep(1 * time.Second)
			} else {
				log.Println("ping success")
				c <- nil
				return
			}
		}
	}
	c <- errors.New("connect failed")
}

func Connect(url string) error {
	var err error
	cdt := options.Credential{Username: dbuser, Password: dbpw}
	client, err = mongo.NewClient(options.Client().ApplyURI(url).SetAuth(cdt))
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
