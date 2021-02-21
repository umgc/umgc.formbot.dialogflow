package main

import (
	"net/http"
)

type test_struct struct {
	Name    string
	Address string
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Authenticated!"))
	return
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}