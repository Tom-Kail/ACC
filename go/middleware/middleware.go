// middleware
package main

import (
	"net/http"
	"net/http/httptest"
)

type Middleware struct {
	handler      http.Handler
	handler2     http.Handler
	handlerPath  string
	handler2Path string
}

func (this *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == this.handler2Path {
		rec := httptest.NewRecorder()
		this.handler2.ServeHTTP(rec, r)
		w.Write(rec.Body.Bytes())
		w.Write([]byte("youyoyoyo！\n"))
		//		this.handler2.ServeHTTP(w, r)
	} else if r.URL.Path == this.handlerPath {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from handler2"))
}

func main() {
	middle := Middleware{
		handler:      http.HandlerFunc(myHandler),
		handler2:     http.HandlerFunc(myHandler2),
		handlerPath:  "/1",
		handler2Path: "/2",
	}
	http.ListenAndServe(":8080", &middle)
}
