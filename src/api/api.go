package api

import (
	"net/http"
)

var (
	session    *sessionController    = new(sessionController)
)

func Setup() {
	http.HandleFunc("/api/sessions", session.PostSession)
	http.HandleFunc("/api/users", session.PostUser)
}