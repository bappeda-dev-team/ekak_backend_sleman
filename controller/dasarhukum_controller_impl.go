package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/dasarhukum"
	"ekak_kab_sleman/service"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type DasarHukumControllerImpl struct {
	DasarHukumService service.DasarHukumService
}

func NewDasarHukumControllerImpl(dasarHukumService service.DasarHukumService) *DasarHukumControllerImpl {
	return &DasarHukumControllerImpl{
		DasarHukumService: dasarHukumService,
	}
}

func (controller *DasarHukumControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")
	if rekinId == "" {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "RekinId tidak boleh kosong",
		})
		return
	}

	// Decode request body
	dasarHukumCreateRequest := dasarhukum.DasarHukumCreateRequest{}
	helper.ReadFromRequestBody(request, &dasarHukumCreateRequest)

	// Set rekinId dari params ke request
	dasarHukumCreateRequest.RekinId = rekinId

	// Panggil service untuk membuat gambaran umum
	dasarHukumResponse, err := controller.DasarHukumService.Create(request.Context(), dasarHukumCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusCreated,
		Status: "success create data dasar hukum",
		Data:   dasarHukumResponse,
	})
}

func (controller *DasarHukumControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dasarHukumUpdateRequest := dasarhukum.DasarHukumUpdateRequest{}
	helper.ReadFromRequestBody(request, &dasarHukumUpdateRequest)

	dasarHukumId := params.ByName("id")
	dasarHukumUpdateRequest.Id = dasarHukumId

	dasarHukumResponse, err := controller.DasarHukumService.Update(request.Context(), dasarHukumUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success update data dasar hukum",
		Data:   dasarHukumResponse,
	})
}

func (controller *DasarHukumControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dasarHukumId := params.ByName("id")

	dasarHukumResponse, err := controller.DasarHukumService.FindById(request.Context(), dasarHukumId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success get data dasar hukum",
		Data:   dasarHukumResponse,
	})
}

func (controller *DasarHukumControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")

	dasarHukumResponses, err := controller.DasarHukumService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success get all data dasar hukum",
		Data:   dasarHukumResponses,
	})
}

func (controller *DasarHukumControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dasarHukumId := params.ByName("id")

	err := controller.DasarHukumService.Delete(request.Context(), dasarHukumId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success delete data dasar hukum",
		Data:   nil,
	})
}

func (controller *DasarHukumControllerImpl) FindAllByRekinId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")

	dasarHukumResponses, err := controller.DasarHukumService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	buttonActions := []web.ActionButton{
		{
			NameAction: "Create Dasar Hukum",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/dasar_hukum/create/:rencana_kinerja_id", host, port),
		},
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success get all data dasar hukum",
		Action: buttonActions,
		Data:   dasarHukumResponses,
	})
}
