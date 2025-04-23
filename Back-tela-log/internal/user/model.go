package user

import (
	"database/sql"
	"net/http"
)

type User struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Api struct {
	Addr    string
	Repo    *Repository
	Service *Service
	Handler *Handler
	Rout    http.Handler
}

func NewApi(Addr string, Db *sql.DB) *Api {
	r := NewRepository(Db)
	s := NewService(r)
	h := NewHandler(s)

	router := http.NewServeMux()

	router.HandleFunc("/cadastro", h.HandleCadastro)

	a := &Api{
		Addr:    Addr,
		Repo:    r,
		Service: s,
		Handler: h,
		Rout:    router,
	}

	return a
}
