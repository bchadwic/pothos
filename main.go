package main

import (
	"net/http"
	"os"
)

func main() {
	handler := &Handler{}
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}

type Handler struct{}

func (handler *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	body, err := os.ReadFile("./resources/index.html")
	if err != nil {
		panic(err)
	}
	resp.Write(body)
}
