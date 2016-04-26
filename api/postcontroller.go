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

	dto, err := apimodel.FillPostDtoFromMap(params)
	if err != nil{
		return InternalServerError(err)
	}

	if dto.OwnerId != r.User.Uuid {
		return UnauthorizedError(nil)
	}

	post := apimodel.PostModelFromPostDto(dto);

	post, err = model.CreatePost(post)
	if err != nil {
		return InternalServerError(err)
	}

	w.WriteHeader(http.StatusCreated)
	responseObject, err := json.Marshal(apimodel.PostDtoFromPostModel(post))
	w.Write(responseObject)
	return nil
}

func (sc *postController) GetPost(w http.ResponseWriter, r *ApiRequest) *apiError {
	args := mux.Vars(r.Request)
	id := args["id"]

	authorized := services.IsUserAuthorizedForResource(r.User.Uuid, id);
	if !authorized {
		return UnauthorizedError(nil)
	}

	post, err := model.GetPostById(id)
	if err != nil{
		return NotFoundError(err)
	}

	resp, _ := json.Marshal(apimodel.PostDtoFromPostModel(post))
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

