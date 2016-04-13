package controller

import (
	"github.com/danthehitman/poster/src/viewmodel"
	"html/template"
	"net/http"
)

type loginController struct {
	loginTemplate *template.Template
}

func (lc *loginController) GetLogin(w http.ResponseWriter, r *http.Request) {
	vmodel := viewmodel.Login{}
	lc.loginTemplate.Execute(w, vmodel)
}
