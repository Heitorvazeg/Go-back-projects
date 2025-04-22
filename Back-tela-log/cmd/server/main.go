package main

import (
	"net/http"

	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/db"
	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/user"
)

func main() {
	db := db.Connect()
	defer db.Close()

	a := user.NewApi(":8081", db)

	srv := http.Server{
		Addr:    a.Addr,
		Handler: mid.MidLog(a.Rout),
	}

	srv.ListenAndServe()
}
