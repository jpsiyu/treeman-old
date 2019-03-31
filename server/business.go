package main

import (
	"time"

	"github.com/jpsiyu/treeman/server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	filter := bson.M{"uid": id}
	err := db.Find(results, &filter, db.CollectionRecord)
	if err != nil {
		return err
	}
	return nil
}

func AddRecord(id, detail, comment string) error {
	timestamp := time.Now().Unix()
	dbrecord := db.DBRecord{
		Uid:       id,
		Detail:    detail,
		Comment:   comment,
		Timestamp: timestamp,
	}
	err := db.Insert(&dbrecord, db.CollectionRecord)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	err = db.Delete(&filter, db.CollectionRecord)
	if err != nil {
		return err
	}
	return nil

}
