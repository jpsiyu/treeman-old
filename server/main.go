package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jpsiyu/treeman/server/conf"
	"github.com/jpsiyu/treeman/server/db"
)

func logInit() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	writers := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writers)
}

func main() {
	logInit()

	r := mux.NewRouter()
	// set middleware
	r.Use(LoggingMiddleware)
	r.Use(AuthMiddleware)
	// set handler
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("dist/"))))
	r.HandleFunc("/", HandleHome).Methods("GET")
	r.HandleFunc("/serverlog", HandleServerLog).Methods("GET")
	r.HandleFunc("/pubapi/login", HandleLogin).Methods("GET")
	r.HandleFunc("/api/allperson", HandleGetAllPerson).Methods("GET")
	r.HandleFunc("/api/genperson", HandleGenPerson).Methods("POST")
	r.HandleFunc("/api/updateperson", HandleUpdatePerson).Methods("PUT")
	r.HandleFunc("/api/deleteperson", HandleDeletePerson).Methods("PUT")
	r.HandleFunc("/api/findperson", HandleFindPersonByName).Methods("GET")
	r.HandleFunc("/api/record", HandleGetRecord).Methods("GET")
	r.HandleFunc("/api/addrecord", HandleAddRecord).Methods("POST")
	r.HandleFunc("/api/deleterecord", HandleDeleteRecord).Methods("PUT")
	r.HandleFunc("/api/updaterecord", HandleUpdateRecord).Methods("PUT")
	r.NotFoundHandler = http.HandlerFunc(HandleHome)

	// connect database
	dbchan := make(chan error)
	go db.SmartConnect(dbchan)
	err := <-dbchan
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Server listening on port %d", conf.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r)
}
