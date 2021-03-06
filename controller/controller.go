package controller

import (
	"html/template"
	"net/http"
	"config"
)

var (
	login    *loginController    = new(loginController)
)

func Setup(tc *template.Template) {
	SetTemplateCache(tc)
	createResourceServer()

	//http.HandleFunc("/", login.GetLogin)
}

func createResourceServer() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir(config.ResourcePrefix + "resources"))))
}

func SetTemplateCache(tc *template.Template) {
	login.loginTemplate = tc.Lookup("login.html")
}