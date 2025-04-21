package main

import (
	"net/http"

	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/db"
	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/user"
	mid "github.com/Heitorvazeg/Go-back-projects/Back-tela-log/pkg/middleware"
)

func main() {
	a := &user.Api{
		Addr: ":8081",
	}

	srv := http.Server{
		Addr:    a.Addr,
		Handler: mid.MidLog(user.Handler),
	}

	db := db.Connect()
	defer db.Close()

	router := http.NewServeMux()

	router.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {

	})

	srv.ListenAndServe()
}
