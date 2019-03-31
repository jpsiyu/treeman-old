package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

// base routing
func HandleHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	indexPath := "dist/index.html"
	data, err := ioutil.ReadFile(indexPath)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	} else {
		w.Write(data)
	}
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.WriteHeader(404)
	w.Write([]byte("404 - " + http.StatusText(404)))
}

// business routing

// person collection
func HandleGetAllPerson(w http.ResponseWriter, r *http.Request) {
	var results []bson.M
	err := GetAllPerson(&results)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	encode, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte(encode))
}

func HandleGenPerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	log.Println(r.PostForm, r.Form)
	name := r.FormValue("name")
	gender, err := strconv.Atoi(r.FormValue("gender"))
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	err = GenPerson(name, gender, age)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
}

func HandleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	name := r.FormValue("name")
	gender, err := strconv.Atoi(r.FormValue("gender"))
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	err = UpdatePerson(name, gender, age)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
}

func HandleDeletePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	err := DeletePerson(name)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
}

// record collection

func HandleGetRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	var results []bson.M
	err := GetRecord(&results, id)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	encode, err := json.Marshal(&results)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte(encode))
}

func HandleAddRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	detail := r.FormValue("detail")
	comment := r.FormValue("comment")
	err := AddRecord(id, detail, comment)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
}

func HandleDeleteRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	err := DeleteRecord(id)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte("ok"))
}
