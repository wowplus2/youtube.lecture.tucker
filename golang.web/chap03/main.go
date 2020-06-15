package main

import (
	"net/http"

	"web.chap03/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
