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

		fmt.Printf("%s|[%s] %s| %d: %s\n", start, r.Method, r.URL, lrw.StatusCode, lrw.Body.String())
	})
}
