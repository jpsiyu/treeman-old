package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpsiyu/treeman/server/db"
)

func main() {
	// set log config
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := mux.NewRouter()
	// set handler
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("dist/"))))
	r.HandleFunc("/", HandleHome).Methods("GET")
	r.HandleFunc("/api/allperson", HandleGetAllPerson).Methods("GET")
	r.HandleFunc("/api/genperson", HandleGenPerson).Methods("POST")
	r.HandleFunc("/api/updateperson", HandleUpdatePerson).Methods("PUT")
	r.HandleFunc("/api/deleteperson", HandleDeletePerson).Methods("PUT")
	r.HandleFunc("/api/record", HandleGetRecord).Methods("GET")
	r.HandleFunc("/api/addrecord", HandleAddRecord).Methods("POST")
	r.HandleFunc("/api/deleterecord", HandleDeleteRecord).Methods("PUT")
	r.NotFoundHandler = http.HandlerFunc(HandleHome)

	// connect database
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established!")

	log.Println(fmt.Sprintf("Server listening on port %d", Port))
	http.ListenAndServe(fmt.Sprintf(":%d", Port), r)
}
