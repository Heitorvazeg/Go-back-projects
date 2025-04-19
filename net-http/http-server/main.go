package main

import (
	"fmt"
	"net/http"

	"github.com/Heitorvazeg/Go-back-projects/net-http/http-server/api"
	midd "github.com/Heitorvazeg/Go-back-projects/net-http/http-server/middleware"
)

func main() {
	a := &api.Api{Addr: ":8080"}

	router := http.NewServeMux()

	srv := http.Server{
		Addr:    a.Addr,
		Handler: midd.MiddlewareLog(router),
	}

	router.Handle("/", a)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Users page")
	})

	router.Handle("/profile", midd.MiddlewareAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profile page protegida")
	})))

	srv.ListenAndServe()
}
