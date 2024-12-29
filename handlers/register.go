package handlers

import (
	"net/http"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	getHtmlTemplate().ExecuteTemplate(w, "register.html", struct{ Register bool }{Register: r.URL.Path == "/register"})
}
