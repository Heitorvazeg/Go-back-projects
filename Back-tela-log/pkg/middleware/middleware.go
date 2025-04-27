package mid

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Heitorvazeg/Go-back-projects/Back-tela-log/internal/user"
)

func MidLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := user.NewLGRW(w)
		start := time.Now()

		next.ServeHTTP(lrw, r)

		if lrw.StatusCode >= 200 && lrw.StatusCode <= 204 {
			fmt.Printf("%s|[%s] %s| %d\n", start, r.Method, r.URL, lrw.StatusCode)

		} else {
			fmt.Printf("%s|[%s] %s| %d: %s\n", start, r.Method, r.URL, lrw.StatusCode, lrw.Body.String())
		}
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
