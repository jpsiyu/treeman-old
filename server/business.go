package main

import (
	"github.com/jpsiyu/treeman/server/db"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPerson(results *[]bson.M) error {
	filter := bson.M{}
	err := db.Find(results, &filter, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}

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

func DeletePerson(name string) error {
	filter := bson.M{"name": name}
	err := db.Delete(&filter, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}

func GetRecord(results *[]bson.M, id string) error {
	filter := bson.M{"_id": id}
	err := db.Find(results, &filter, db.CollectionRecord)
	if err != nil {
		return err
	}
	return nil
}
