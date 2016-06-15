package api

import (
	"net/http"
	"encoding/json"
	"model"
	"github.com/gorilla/mux"
	"apimodel"
	"services"
)

type userController struct {
}

func (sc *userController) PostUser(w http.ResponseWriter, r *http.Request) *apiError {
	params, err := decodeAndValidateRequest(*r, apimodel.RegisterUserDto{}, nil)
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.FillRegisterUserDto(params)
	if err != nil{
		return InternalServerError(err)
	}

	user, err := apimodel.UserModelFromRegisterDto(dto)
	if (err != nil){
		return InternalServerError(err)
	}

	tx :=model.Db.Begin()
	*user, err = model.CreateUser(*user)
	if (err != nil){
		tx.Rollback()
		return InternalServerError(err)
	}

	err = model.CreateResourceAuthorization(model.ResourceAuthorization{ UserId:user.Uuid, ResourceId:user.Uuid, Action:model.EditResourceAction})
	if (err!= nil){
		tx.Rollback()
		return InternalServerError(err)
	}

	responseDto, err := apimodel.UserDtoFromUserModel(*user)
	if (err != nil){
		tx.Rollback()
		return InternalServerError(err)
	}
	tx.Commit()

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	responseObject, err := json.Marshal(responseDto)
	w.Write(responseObject)
	return nil
}

func (sc *userController) GetUser(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResourceRead(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	user, err := model.GetUserById(id)
	if err != nil{
		return NotFoundError(err)
	}

	responseDto, err := apimodel.UserDtoFromUserModel(user)
	if (err != nil){
		return InternalServerError(err)
	}
	responseObject, _ := json.Marshal(responseDto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseObject)
	return nil
}
//
//func (sc *userController) GetUsers(w http.ResponseWriter, r *ApiRequest) *apiError {
//	user, err := model.GetUsersById(id)
//	if err != nil{
//		return NotFoundError(err)
//	}
//
//	responseDto, err := apimodel.UserDtoFromUserModel(user)
//	if (err != nil){
//		return InternalServerError(err)
//	}
//	responseObject, _ := json.Marshal(responseDto)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write(responseObject)
//	return nil
//}

func (sc *userController) DeleteUser(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResourceRead(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	err := model.DeleteUserById(id)
	if err != nil {
		return InternalServerError(err)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
	return nil
}

func (sc *userController) GetJournalsByUser(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	var (
		journals []model.Journal
		err error
	)
	if r.User == nil || services.IsUserAuthorizedForResourceRead(r.User.Uuid, id) {
		journals, err = model.GetPublicJournalsByUserId(id)
	} else {
		journals, err = model.GetAuthorizedJournalsForUser(r.User.Uuid, false)
	}
	if err != nil{
		return InternalServerError(err)
	}

	responseDtos := apimodel.JournalDto{}.DtosFromModels(journals)
	if (err != nil){
		return InternalServerError(err)
	}

	responseObject, _ := json.Marshal(responseDtos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseObject)
	return nil
}