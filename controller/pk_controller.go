package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PkController interface {
	FindAllPkOpdTahunan(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HubungkanRekin(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HubungkanAtasan(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
