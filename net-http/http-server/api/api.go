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

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
