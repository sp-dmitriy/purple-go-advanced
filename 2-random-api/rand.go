package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

type RandHandler struct{}

func NewRandHandler(router *http.ServeMux) {
	handler := &RandHandler{}
	router.HandleFunc("/rand", handler.Rand())
}

func (handler *RandHandler) Rand() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		randomNumber := rand.IntN(6) + 1
		w.Write([]byte(fmt.Sprintf("%d", randomNumber)))
	}
}
