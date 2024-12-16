package handlers

import (
	"encoding/json"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, status int, err error) {
	// TODO: add error Page
	ErrorJs(w, status, err)
}

func ErrorJs(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

const notaUser = 0
