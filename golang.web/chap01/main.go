package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bar!")
}

func main() {
	// HandleFunc를 이용한 라우터 구현
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// HandleFunc와 Handler함수가 분리된 형태
	http.HandleFunc("/bar", barHandler)

	// Handler를 이용한 라우터 구현
	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
