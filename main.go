package main

import (
	"net/http"
)

func main() {
	handler := &Handler{}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}

type Handler struct{}

func (handler *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello!"))
}
