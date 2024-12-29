package handlers

import (
	"encoding/json"
	"net/http"
)

func ErrorPage(w http.ResponseWriter,page string, status int, err error) {
	w.WriteHeader(status)
	getHtmlTemplate().ExecuteTemplate(w, page, map[string]interface{}{
		"StatuCode":    status,
		"MessageError": err,
	})
}

func ErrorJs(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
