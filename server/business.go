package main

import (
	"github.com/jpsiyu/treeman/server/db"
	"go.mongodb.org/mongo-driver/bson"
)

func GenPerson(name string, gender, age int) error {
	dbperson := db.DBPerson{
		Name:   name,
		Gender: gender,
		Age:    age,
	}
	err := db.Insert(&dbperson, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePerson(name string, gender, age int) error {
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"gender": gender, "age": age}}
	err := db.Update(&filter, &update, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}
