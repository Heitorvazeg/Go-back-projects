package main

import (
	"log"
	"net/http"

	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/db"
	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/user"
	mid "github.com/Heitorvazeg/Go-back-projects/Back-tela-log/pkg/middleware"
)

func main() {
	db := db.Connect()
	defer db.Close()

	a := user.NewApi(":8081", db)

	srv := http.Server{
		Addr:    a.Addr,
		Handler: mid.CORS(mid.MidLog(a.Rout)),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
