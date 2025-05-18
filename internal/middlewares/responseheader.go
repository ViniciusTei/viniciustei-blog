package middlewares

import "net/http"

type ResponseHeader struct {
	headerName  string
	headerValue string
}

// NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{headerName, headerValue}
}

func (rh *ResponseHeader) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(rh.headerName, rh.headerValue)
		next.ServeHTTP(w, r)
	})
}
