package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"model"
)

var (
	session = new(sessionController)
	user = new(userController)
)

type ApiRequest struct {
	*http.Request
	User *model.User
}

type appHandler func(http.ResponseWriter, *http.Request) *apiError

type secureAppHandler func(http.ResponseWriter, *ApiRequest) * apiError

func (fn secureAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := ApiRequest{r, nil}
	user := AuthenticatedUser(*r)
	if user == nil {
		http.Error(w, "", http.StatusUnauthorized)
	} else {
		request.User = user
		if e := fn(w, &request); e != nil {
			http.Error(w, e.Message, e.Code)
		}
	}
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		http.Error(w, e.Message, e.Code)
	}
}

func Setup() {
	r := mux.NewRouter()

	// Register unsecured routes
	r.Handle("/api/sessions", appHandler(session.PostSession))

	// Register secured routes
	r.Handle("/api/users", secureAppHandler(user.PostUser))
	r.Handle("/api/users/{id}", secureAppHandler(user.GetUser))

	http.Handle("/", r)
}