package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct {
}

func (l *Logger) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		currentTime, err := time.Parse(time.RFC822, time.Now().Format(time.RFC822))
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%s]: %s %s %v\n", currentTime, r.Method, r.URL.Path, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

func NewLogger() *Logger {
	return &Logger{}
}
