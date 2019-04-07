package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jpsiyu/treeman/server/conf"

	"github.com/jpsiyu/treeman/server/auth"
	"go.mongodb.org/mongo-driver/bson"
)

// base routing
func HandleHome(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(404)
	w.Write([]byte("404 - " + http.StatusText(404)))
}

func normalServerErr(w http.ResponseWriter, err error) {
	log.Println(err)
	encode, _ := json.Marshal(ServerMsg{2, nil})
	w.Write(encode)
}

// auth
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("user")
	password := r.FormValue("password")
	authUser := auth.AuthUser{
		User:     user,
		Password: password,
	}

	log.Println(user, user == "treeman")
	if user != conf.AccessUser || password != conf.AccessPasswd {
		encode, _ := json.Marshal(ServerMsg{2, "Access denied"})
		w.Write([]byte(encode))
		return
	}

	tokenStr, err := auth.GenTokenStr(&authUser)

	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, tokenStr})
	w.Write([]byte(encode))
}

// person collection
func HandleGetAllPerson(w http.ResponseWriter, r *http.Request) {
	var results []bson.M
	err := GetAllPerson(&results)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, results})
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
		normalServerErr(w, err)
		return
	}
	err = GenPerson(name, gender, age)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}

func HandleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var err error
	name := r.FormValue("name")
	gender, err := strconv.Atoi(r.FormValue("gender"))
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		normalServerErr(w, err)
		return
	}
	err = UpdatePerson(name, gender, age)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}

func HandleDeletePerson(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	err := DeletePerson(id)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}

func HandleFindPersonByName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	var results []bson.M
	err := FindPerson(&results, name)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, &results})
	w.Write([]byte(encode))
}

// record collection

func HandleGetRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	var results []bson.M
	err := GetRecord(&results, id)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, &results})
	w.Write([]byte(encode))
}

func HandleAddRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	detail := r.FormValue("detail")
	comment := r.FormValue("comment")
	err := AddRecord(id, detail, comment)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}

func HandleDeleteRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	err := DeleteRecord(id)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}

func HandleUpdateRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	detail := r.FormValue("detail")
	comment := r.FormValue("comment")
	err := UpdateRecord(id, detail, comment)
	if err != nil {
		normalServerErr(w, err)
		return
	}
	encode, _ := json.Marshal(ServerMsg{0, "ok"})
	w.Write([]byte(encode))
}
