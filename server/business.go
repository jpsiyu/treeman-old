package main

import (
	"fmt"
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

func UpdatePerson(id string, name string, gender, age int) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"name": name, "gender": gender, "age": age}}
	err = db.Update(&filter, &update, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	err = db.Delete(&filter, db.CollectionPerson)
	if err != nil {
		return err
	}
	return nil
}

func FindPerson(results *[]bson.M, name string) error {
	regexStr := fmt.Sprintf(".*%s.*", name)
	filter := bson.M{"name": bson.M{"$regex": regexStr}}
	//filter := bson.M{"name": name}
	err := db.Find(results, &filter, db.CollectionPerson)
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

func UpdateRecord(id, detail, comment string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"detail": detail, "comment": comment}}
	err = db.Update(&filter, &update, db.CollectionRecord)
	if err != nil {
		return err
	}
	return nil

}
