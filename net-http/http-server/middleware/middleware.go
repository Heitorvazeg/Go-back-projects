package midd

import (
	"fmt"
	"net/http"
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader != "Bearer 12345" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Não autorizado")
			return
		}
		next.ServeHTTP(w, r)
	})
}
