package api

import (
	"net/http"
	"encoding/json"
	"model"
	"github.com/gorilla/mux"
	"apimodel"
	"services"
)

type journalController struct {
}

func (sc *journalController) PostJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	params, err := decodeAndValidateRequest(*r.Request, apimodel.JournalDto{}, map[string]bool{"Uuid":true})
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.JournalDto{}.FillDtoFromMap(params)
	if err != nil{
		return InternalServerError(err)
	}

	if dto.OwnerId != r.User.Uuid {
		return UnauthorizedError(nil)
	}

	journal := apimodel.JournalDto{}.ModelFromDto(dto);

	journal, err = model.CreateJournal(journal)
	if err != nil {
		return InternalServerError(err)
	}

	w.WriteHeader(http.StatusCreated)
	responseObject, err := json.Marshal(apimodel.JournalDto{}.DtoFromModel(journal))
	w.Write(responseObject)
	return nil
}

func (sc *journalController) GetJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResourceRead(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	journal, err := model.GetJournalById(id)
	if err != nil{
		return NotFoundError(err)
	}

	resp, _ := json.Marshal(apimodel.JournalDto{}.DtoFromModel(journal))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return nil
}

func (sc *journalController) GetJournals(w http.ResponseWriter, r *ApiRequest) *apiError {
	journals, err := model.GetJournalsForUser(r.User.Uuid)
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

func (sc *journalController) GetPostsForJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForJournalRead(r.User.Uuid, id)
	if !authorized {
		return UnauthorizedError(nil)
	}

	posts, err := model.GetPostsByJournalId(id)
	responseDtos := apimodel.PostDtosFromPostModels(posts)
	if (err != nil){
		return InternalServerError(err)
	}

	responseObject, _ := json.Marshal(responseDtos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseObject)
	return nil
}

func (sc *journalController) GetPostForJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	jid := args["jid"]
	pid := args["pid"]

	authorized := services.IsUserAuthorizedForJournalRead(r.User.Uuid, jid)
	if !authorized {
		return UnauthorizedError(nil)
	}

	if found, err:= model.IsPostInJournal(jid, pid); !found {
		if err != nil {
			return InternalServerError(err)
		}
		return NotFoundError(nil)
	}

	post, err := model.GetPostById(pid)
	if err != nil{
		return NotFoundError(err)
	}

	resp, _ := json.Marshal(apimodel.PostDtoFromPostModel(post))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return nil
}

