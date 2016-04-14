package api

import (
	"net/http"
	"encoding/json"
	"github.com/danthehitman/poster/model"
)

type userController struct {
}

func (sc *sessionController) PostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic("Could not deserialize request.")
	}
	user, err = model.CreateUser(user)
	resp, _ := json.Marshal(user)
	w.Write(resp)
}