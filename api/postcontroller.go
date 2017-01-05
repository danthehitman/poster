package api

import (
	"net/http"
	"encoding/json"
	"model"
	"github.com/gorilla/mux"
	"apimodel"
	"services"
)

type postController struct {
}

func (sc *postController) PostPost(w http.ResponseWriter, r *ApiRequest) *apiError {
	params, err := decodeAndValidateRequest(*r.Request, apimodel.PostDto{}, map[string]bool{"Uuid":true})
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.PostDto{}.FillDtoFromMap(params)
	if err != nil{
		return InternalServerError(err)
	}

	if dto.OwnerId != r.User.Uuid {
		return UnauthorizedError(nil)
	}

	post := apimodel.PostDto{}.ModelFromDto(dto);

	post, err = model.CreatePost(post)
	if err != nil {
		return InternalServerError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseObject, err := json.Marshal(apimodel.PostDto{}.DtoFromModel(post))
	w.Write(responseObject)
	return nil
}

func (sc *postController) PutPost(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResourceEdit(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	existing, err := model.GetPostById(id)

	if err != nil{
		return NotFoundError(err)
	}

	params, err := decodeAndValidateRequest(*r.Request, apimodel.PostDto{}, map[string]bool{"Uuid":true})
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.PostDto{}.FillDtoFromMap(params)
	if err != nil{
		return InternalServerError(err)
	}

	if existing.OwnerId != r.User.Uuid {
		return UnauthorizedError(nil)
	}

	new := apimodel.PostDto{}.ModelFromDto(dto);
	existingAsDto := apimodel.PostDto{}.DtoFromModel(existing)
	err = checkReadonlyFields(existingAsDto, dto)
	if err != nil{
		return BadRequestError(err)
	}

	new, err = model.UpdatePost(new)
	if err != nil {
		return InternalServerError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseObject, err := json.Marshal(apimodel.PostDto{}.DtoFromModel(new))
	w.Write(responseObject)
	return nil
}

func (sc *postController) GetPost(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResourceRead(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	post, err := model.GetPostById(id)
	if err != nil{
		return NotFoundError(err)
	}

	resp, _ := json.Marshal(apimodel.PostDto{}.DtoFromModel(post))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return nil
}

func (sc *postController) GetPosts(w http.ResponseWriter, r *ApiRequest) *apiError {
	posts, err := model.GetPostsForUser(r.User.Uuid)
	if err != nil{
		return InternalServerError(err)
	}

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

