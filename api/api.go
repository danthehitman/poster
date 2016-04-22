package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"model"
)

var (
	sessionCon = new(sessionController)
	userCon = new(userController)
	postCon = new(postController)
)

type ApiRequest struct {
	*http.Request
	User *model.User
}

// Use this handler when the request does not need to be secured (e.g. registration)
type appHandler func(http.ResponseWriter, *http.Request) *apiError

// Use this handler when the request will be secured with an auth token.
type secureAppHandler func(http.ResponseWriter, *ApiRequest) *apiError

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

func Init() {
	r := mux.NewRouter()

	// Register unsecured routes
	r.Handle("/api/sessions", appHandler(sessionCon.PostSession))

	// Register secured routes
	r.Handle("/api/users", secureAppHandler(userCon.PostUser))
	r.Handle("/api/users/{id}", secureAppHandler(userCon.GetUser))
	r.Handle("/api/posts", secureAppHandler(postCon.PostPost))
	r.Handle("/api/posts/{id}", secureAppHandler(postCon.GetPost))

	http.Handle("/", r)
}