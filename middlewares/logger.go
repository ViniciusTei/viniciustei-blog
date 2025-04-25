package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	currentTime, err := time.Parse(time.RFC822, time.Now().Format(time.RFC822))
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s]: %s %s %v\n", currentTime, r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}
