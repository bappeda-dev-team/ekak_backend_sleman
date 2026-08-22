package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/usulan"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsulanTerpilihControllerImpl struct {
	UsulanTerpilihService service.UsulanTerpilihService
}

func NewUsulanTerpilihControllerImpl(usulanTerpilihService service.UsulanTerpilihService) *UsulanTerpilihControllerImpl {
	return &UsulanTerpilihControllerImpl{UsulanTerpilihService: usulanTerpilihService}
}

func (controller *UsulanTerpilihControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanTerpilihCreateRequest := usulan.UsulanTerpilihCreateRequest{}
	helper.ReadFromRequestBody(request, &usulanTerpilihCreateRequest)

	usulanTerpilihResponse, err := controller.UsulanTerpilihService.Create(request.Context(), usulanTerpilihCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "success create usulan terpilih",
		Data:   usulanTerpilihResponse,
	}
	writer.WriteHeader(http.StatusCreated)
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanTerpilihControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id_usulan")

	err := controller.UsulanTerpilihService.Delete(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan terpilih",
	}
	writer.WriteHeader(http.StatusOK)
	helper.WriteToResponseBody(writer, webResponse)
}
