package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/lembaga"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LembagaControllerImpl struct {
	LembagaService service.LembagaService
}

func NewLembagaControllerImpl(lembagaService service.LembagaService) *LembagaControllerImpl {
	return &LembagaControllerImpl{
		LembagaService: lembagaService,
	}
}

func (controller *LembagaControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lembagaCreateRequest := lembaga.LembagaCreateRequest{}
	helper.ReadFromRequestBody(request, &lembagaCreateRequest)

	lembagaResponse, err := controller.LembagaService.Create(request.Context(), lembagaCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   lembagaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LembagaControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lembagaId := params.ByName("id")

	lembagaUpdateRequest := lembaga.LembagaUpdateRequest{}
	helper.ReadFromRequestBody(request, &lembagaUpdateRequest)

	lembagaUpdateRequest.Id = lembagaId

	lembagaResponse, err := controller.LembagaService.Update(request.Context(), lembagaUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   lembagaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LembagaControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lembagaId := params.ByName("id")

	controller.LembagaService.Delete(request.Context(), lembagaId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *LembagaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lembagaId := params.ByName("id")

	lembagaResponse, err := controller.LembagaService.FindById(request.Context(), lembagaId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   lembagaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *LembagaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	lembagaResponse, err := controller.LembagaService.FindAll(request.Context())
	if err != nil {
		webResponse := web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   lembagaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
