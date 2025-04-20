package api

import (
	"fmt"
	"net/http"
)

type Api struct {
	Addr string
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo!")
}

type Response struct {
	Message string `json:"message"`
}
