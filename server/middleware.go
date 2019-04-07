package main

import (
	"fmt"
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
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		authUser := auth.AuthUser{}
		err := auth.ParseTokenStr(tokenStr, &authUser)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("%s: %s", http.StatusText(http.StatusUnauthorized), err.Error())))
			return
		}
		next.ServeHTTP(w, r)
	})
}
