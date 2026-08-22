package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/rencanaaksi"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PelaksanaanRencanaAksiControllerImpl struct {
	PelaksanaanRencanaAksiService service.PelaksanaanRencanaAksiService
}

func NewPelaksanaanRencanaAksiControllerImpl(pelaksanaanRencanaAksiService service.PelaksanaanRencanaAksiService) *PelaksanaanRencanaAksiControllerImpl {
	return &PelaksanaanRencanaAksiControllerImpl{
		PelaksanaanRencanaAksiService: pelaksanaanRencanaAksiService,
	}
}

func (controller *PelaksanaanRencanaAksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createRequest := rencanaaksi.PelaksanaanRencanaAksiCreateRequest{}
	helper.ReadFromRequestBody(request, &createRequest)

	// Mengambil rencanaAksiId dari query URL
	rencanaAksiId := params.ByName("rencanaAksiId")
	if rencanaAksiId == "" {
		webResponse := web.WebPelaksanaanRencanaAksiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "rencanaAksiId tidak boleh kosong",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	createRequest.RencanaAksiId = rencanaAksiId

	response, err := controller.PelaksanaanRencanaAksiService.Create(request.Context(), createRequest)
	if err != nil {
		webResponse := web.WebPelaksanaanRencanaAksiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebPelaksanaanRencanaAksiResponse{
		Code:   http.StatusCreated,
		Status: "success create pelaksanaan rencana aksi",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PelaksanaanRencanaAksiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateRequest := rencanaaksi.PelaksanaanRencanaAksiUpdateRequest{}
	updateRequest.Id = params.ByName("pelaksanaanRencanaAksiId")
	helper.ReadFromRequestBody(request, &updateRequest)

	response, err := controller.PelaksanaanRencanaAksiService.Update(request.Context(), updateRequest)
	if err != nil {
		if customErr, ok := err.(*web.CustomError); ok {
			webResponse := web.WebPelaksanaanRencanaAksiResponse{
				Code:   customErr.Code,
				Status: http.StatusText(customErr.Code),
				Data:   customErr.Message,
			}
			helper.WriteToResponseBody(writer, webResponse)
		} else {
			webResponse := web.WebPelaksanaanRencanaAksiResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   err.Error(),
			}
			helper.WriteToResponseBody(writer, webResponse)
		}
		return
	}

	webResponse := web.WebPelaksanaanRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PelaksanaanRencanaAksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	response, err := controller.PelaksanaanRencanaAksiService.FindById(request.Context(), id)
	if err != nil {
		webResponse := web.WebPelaksanaanRencanaAksiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebPelaksanaanRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PelaksanaanRencanaAksiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := controller.PelaksanaanRencanaAksiService.Delete(request.Context(), id)
	if err != nil {
		webResponse := web.WebPelaksanaanRencanaAksiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebPelaksanaanRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Pelaksanaan Rencana Aksi berhasil dihapus",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PelaksanaanRencanaAksiControllerImpl) FindByRencanaAksiId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rencanaAksiId := params.ByName("rencanaAksiId")

	responses, err := controller.PelaksanaanRencanaAksiService.FindByRencanaAksiId(request.Context(), rencanaAksiId)
	if err != nil {
		webResponse := web.WebPelaksanaanRencanaAksiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebPelaksanaanRencanaAksiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
