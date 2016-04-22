package api

import (
	"net/http"
	"model"
	"apimodel"
	"golang.org/x/crypto/bcrypt"
)

type sessionController struct {
}

func (sc *sessionController) PostSession(w http.ResponseWriter, r *http.Request) *apiError {
	params, err := decodeAndValidateRequest(*r, apimodel.SessionParameters{}, nil)
	if err != nil{
		return BadRequestError(err)
	}

	dto, err := apimodel.FillSessionParameters(params)
	if err != nil{
		return InternalServerError(err)
	}

	var sessionToken string

	user, err := model.GetUserByEmail(dto.Email)
	if err != nil {
		return UnauthorizedError(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return UnauthorizedError(err)
	}

	sessionToken, err = model.CreateSession(user)
	if (err != nil){
		return InternalServerError(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(sessionToken))
	return nil
}




