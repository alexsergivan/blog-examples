package main

import (
	"bytes"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", myFirstHandler)
	r.With(myMiddleware).Get("/other", myFirstHandler)

	http.ListenAndServe(":3000", r)
}

func myFirstHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a main page"))
}

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (myrw *MyResponseWriter) Write(p []byte) (int, error) {
	return myrw.buf.Write(p)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a response wrapper:
		myResponseWriter := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}
		next.ServeHTTP(myResponseWriter, r)
		myResponseWriter.buf.WriteString(" and some additional modifications")
		if _, err := io.Copy(w, myResponseWriter.buf); err != nil {
			log.Printf("Failed to send out response: %v", err)
		}
	})
}