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
	params, err := decodeAndValidateRequest(*r.Request, apimodel.CreateJournalDto{}, map[string]bool{"Uuid":true})
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseObject, err := json.Marshal(apimodel.JournalDto{}.DtoFromModel(journal))
	w.Write(responseObject)
	return nil
}

func (sc *journalController) PutJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	existingJournal, err := model.GetJournalById(id)

	authorized := services.IsUserAuthorizedForResourceEdit(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	if err != nil{
		return NotFoundError(err)
	}

	params, err := decodeAndValidateRequest(*r.Request, apimodel.JournalDto{}, map[string]bool{"Uuid":true})
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.JournalDto{}.FillDtoFromMap(params)
	if err != nil{
		return InternalServerError(err)
	}

	if existingJournal.OwnerId != r.User.Uuid {
		return UnauthorizedError(nil)
	}

	newJournal := apimodel.JournalDto{}.ModelFromDto(dto);
	existingAsDto := apimodel.JournalDto{}.DtoFromModel(existingJournal)
	err = checkReadonlyFields(existingAsDto, dto)
	if err != nil{
		return BadRequestError(err)
	}

	newJournal, err = model.UpdateJournal(newJournal)
	if err != nil {
		return InternalServerError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseObject, err := json.Marshal(apimodel.JournalDto{}.DtoFromModel(newJournal))
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
	var (
		journals []model.Journal
		err error
	)
	if r.User == nil {
		journals, err = model.GetPublicJournals()
	} else {
		journals, err = model.GetAuthorizedJournalsForUser(r.User.Uuid, true)
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

func (sc *journalController) GetPostsForJournal(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForJournalRead(r.User.Uuid, id)
	if !authorized {
		return UnauthorizedError(nil)
	}

	posts, err := model.GetPostsByJournalId(id)
	responseDtos := apimodel.PostDto{}.DtosFromModels(posts)
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

	resp, _ := json.Marshal(apimodel.PostDto{}.DtoFromModel(post))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return nil
}

