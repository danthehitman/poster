package api

import (
	"net/http"
	"encoding/json"
	"model"
	"github.com/gorilla/mux"
)

type userController struct {
}

func (sc *userController) PostUser(w http.ResponseWriter, r *ApiRequest) *apiError {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		return BadRequestError(err)
	}
	user, err = model.CreateUser(user)
	resp, _ := json.Marshal(user)
	w.Write(resp)
	return nil
}

func (sc *userController) GetUser(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	user, err := model.GetUserById(id)
	if err != nil{
		return NotFoundError(err)
	}
	resp, _ := json.Marshal(user)
	w.Write(resp)
	return nil
}