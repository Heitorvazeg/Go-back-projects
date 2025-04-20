package main

import (
	"encoding/json"
	"net/http"

	"github.com/Heitorvazeg/Go-back-projects/net-http/http-server/api"
	mid "github.com/Heitorvazeg/Go-back-projects/net-http/http-server/middleware"
)

func main() {
	a := &api.Api{Addr: ":8080"}

	router := http.NewServeMux()

	srv := http.Server{
		Addr:    a.Addr,
		Handler: mid.MiddlewareLog(router),
	}

	router.Handle("/", a)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.Response{Message: "Users Page"})
	})

	router.Handle("/profile", mid.MiddlewareAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api.Response{Message: "Profile page protegida"})
	})))

	srv.ListenAndServe()
}
