package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jpsiyu/treeman/server/auth"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		isApi := strings.HasPrefix(r.RequestURI, "/api")
		if isApi == false {
			next.ServeHTTP(w, r)
			return
		}

		if tokenStr == "" {
			encode, _ := json.Marshal(ServerMsg{1, nil})
			w.Write(encode)
			return
		}

		authUser := auth.AuthUser{}
		err := auth.ParseTokenStr(tokenStr, &authUser)
		if err != nil {
			encode, _ := json.Marshal(ServerMsg{1, nil})
			w.Write(encode)
			return
		}
		log.Println("token", tokenStr)
		next.ServeHTTP(w, r)
	})
}
