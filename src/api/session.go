package api

import (
	"net/http"
	"encoding/json"
)

type sessionController struct {
}

func (sc *sessionController) PostSession(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal("test")
	w.Write(resp)
}

