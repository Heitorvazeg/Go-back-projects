package mid

import (
	"fmt"
	"net/http"
	"time"
)

func MidLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Printf("%s|[%s] %s\n", &start, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
