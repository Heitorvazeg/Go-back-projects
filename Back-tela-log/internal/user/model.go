package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type User struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Log struct {
	Time      time.Time
	Method    string
	Url       string
	Descricao string
}

func NewLog(h *Handler, time time.Time, method string, url *url.URL, descricao http.Header) {
	var headerString string

	for key, values := range descricao {
		for _, value := range values {
			headerString += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	l := &Log{
		Time:      time,
		Method:    method,
		Url:       url.String(),
		Descricao: headerString,
	}

	if err := h.Service.Repo.CreateLog(l); err != nil {
		log.Println("Erro ao criar log: " + err.Error())
	}
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
	router.HandleFunc("/login", h.HandleLogin)

	a := &Api{
		Addr:    Addr,
		Repo:    r,
		Service: s,
		Handler: h,
		Rout:    router,
	}

	return a
}
