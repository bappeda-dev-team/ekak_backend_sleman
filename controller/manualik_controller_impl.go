package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/rencanakinerja"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ManualIKControllerImpl struct {
	ManualIKService service.ManualIKService
}

func NewManualIKControllerImpl(manualIKService service.ManualIKService) *ManualIKControllerImpl {
	return &ManualIKControllerImpl{
		ManualIKService: manualIKService,
	}
}

func (controller *ManualIKControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	indikatorId := params.ByName("indikatorId")

	manualIKCreateRequest := rencanakinerja.ManualIKCreateRequest{}
	helper.ReadFromRequestBody(request, &manualIKCreateRequest)

	manualIKResponse, err := controller.ManualIKService.Create(request.Context(), manualIKCreateRequest, indikatorId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   manualIKResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ManualIKControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	indikatorId := params.ByName("indikatorId")

	manualIKUpdateRequest := rencanakinerja.ManualIKUpdateRequest{}
	helper.ReadFromRequestBody(request, &manualIKUpdateRequest)

	manualIKResponse, err := controller.ManualIKService.Update(request.Context(), manualIKUpdateRequest, indikatorId)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   manualIKResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ManualIKControllerImpl) FindManualIKByIndikatorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	indikatorId := params.ByName("indikatorId")

	manualIKResponses, err := controller.ManualIKService.FindManualIKByIndikatorId(request.Context(), indikatorId)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   manualIKResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ManualIKControllerImpl) FindManualIKSasaranOpdByIndikatorId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	indikatorId := params.ByName("indikatorId")
	tahun := params.ByName("tahun")

	manualIKResponses, err := controller.ManualIKService.FindManualIKSasaranOpdByIndikatorId(request.Context(), indikatorId, tahun)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   manualIKResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
