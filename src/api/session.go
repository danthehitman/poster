package api

import (
	"net/http"
	"encoding/json"
	"model"
)

type sessionController struct {
}

type sessionParameters struct {
	Email string
	Password string
}

func (sc *sessionController) PostSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params sessionParameters
	err := decoder.Decode(&params)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	var sessionToken string
	sessionToken, err = model.CreateSession(params.Email, params.Password)
	if (err != nil){
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(sessionToken))
}

