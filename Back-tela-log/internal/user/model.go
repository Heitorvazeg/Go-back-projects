package user

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type User struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Log struct {
	Time     time.Time
	Method   string
	Url      string
	Status   string
	Response string
}

func NewLog(h *Handler, lrw *loggingResponseWriter, time time.Time,
	method string, url *url.URL) {
	status := strconv.Itoa(lrw.StatusCode)
	l := &Log{
		Time:     time,
		Method:   method,
		Url:      url.String(),
		Status:   status,
		Response: lrw.Body.String(),
	}

	if err := h.Service.Repo.CreateLog(l); err != nil {
		log.Println("Erro ao criar log: " + err.Error())
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       bytes.Buffer
}

func NewLGRW(w http.ResponseWriter) *loggingResponseWriter {
	lrw := &loggingResponseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}

	return lrw
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.Body.Write(b)
	return lrw.ResponseWriter.Write(b)
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
