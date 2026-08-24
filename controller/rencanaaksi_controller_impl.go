package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/rencanaaksi"
	"ekak_kab_sleman/service"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type RencanaAksiControllerImpl struct {
	RencanaAksiService service.RencanaAksiService
}

func NewRencanaAksiControllerImpl(rencanaAksiService service.RencanaAksiService) *RencanaAksiControllerImpl {
	return &RencanaAksiControllerImpl{
		RencanaAksiService: rencanaAksiService,
	}
}

func (controller *RencanaAksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaKinerjaId := params.ByName("rekin_id")

	rencanaAksiCreateRequest := rencanaaksi.RencanaAksiCreateRequest{}
	helper.ReadFromRequestBody(request, &rencanaAksiCreateRequest)

	rencanaAksiCreateRequest.RencanaKinerjaId = rencanaKinerjaId

	rencanaAksiResponse, err := controller.RencanaAksiService.Create(request.Context(), rencanaAksiCreateRequest)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Create Rencana Aksi",
		Data:   rencanaAksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaAksiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaAksiUpdateRequest := rencanaaksi.RencanaAksiUpdateRequest{}
	helper.ReadFromRequestBody(request, &rencanaAksiUpdateRequest)

	rencanaAksiId := params.ByName("rencanaaksiId")
	rencanaAksiUpdateRequest.Id = rencanaAksiId

	rencanaAksiResponse, err := controller.RencanaAksiService.Update(request.Context(), rencanaAksiUpdateRequest)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Update Rencana Aksi",
		Data:   rencanaAksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaAksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaKinerjaId := params.ByName("rencana_kinerja_id")

	rencanaAksiResponses, err := controller.RencanaAksiService.FindAll(request.Context(), rencanaKinerjaId)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Rencana Aksi",
		Data:   rencanaAksiResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaAksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaAksiId := params.ByName("rencanaaksiId")

	rencanaAksiResponse, err := controller.RencanaAksiService.FindById(request.Context(), rencanaAksiId)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Rencana Aksi",
		Data:   rencanaAksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaAksiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaAksiId := params.ByName("rencanaaksiId")

	err := controller.RencanaAksiService.Delete(request.Context(), rencanaAksiId)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Rencana Aksi",
		Data:   "Rencana Aksi berhasil dihapus",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RencanaAksiControllerImpl) FindAllByRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaKinerjaId := params.ByName("rencana_kinerja_id")

	rencanaAksiResponses, err := controller.RencanaAksiService.FindAll(request.Context(), rencanaKinerjaId)
	if err != nil {
		webResponse := web.WebRencanaAksiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	buttonActions := []web.ActionButton{
		{
			NameAction: "Create Rencana Aksi",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/rencana_aksi/create/rencanaaksi/:rekin_id", host, port),
		},
	}

	webResponse := web.WebRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Rencana Aksi",
		Action: buttonActions,
		Data:   rencanaAksiResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
