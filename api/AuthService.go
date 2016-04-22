package api

import (
	"model"
	"net/http"
)

func AuthenticatedUser(r http.Request) *model.User {
	auth := r.Header.Get("auth")
	user := model.GetAuthenticatedUser(auth)
	return user
}
