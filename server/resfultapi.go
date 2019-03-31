package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func responseError(w http.ResponseWriter, err error) {
	w.Write([]byte(fmt.Sprintf("%s", err)))
}

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

func HandleGenPerson(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var err error
	name := query.Get("name")
	gender, err := strconv.Atoi(query.Get("gender"))
	age, err := strconv.Atoi(query.Get("age"))
	if err != nil {
		responseError(w, err)
		return
	}
	err = GenPerson(name, gender, age)
	if err != nil {
		responseError(w, err)
		return
	}
	w.Write([]byte("ok"))
}

func HandleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	name := r.PostForm.Get("name")
	gender, err := strconv.Atoi(r.FormValue("gender"))
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		responseError(w, err)
		return
	}
	err = UpdatePerson(name, gender, age)
	if err != nil {
		responseError(w, err)
		return
	}
	w.Write([]byte("ok"))
}
