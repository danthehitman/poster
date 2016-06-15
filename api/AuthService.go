package api

import (
	"model"
	"net/http"
)

func AuthenticatedUser(r http.Request) *model.User {
	auth := r.Header.Get("auth")
	var user *model.User = nil
	if auth != "" {
		user = model.GetAuthenticatedUser(auth)
	}
	return user
}
