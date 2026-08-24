package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsulanTerpilihController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
