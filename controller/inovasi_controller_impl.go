package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/inovasi"
	"ekak_kab_sleman/service"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type InovasiControllerImpl struct {
	InovasiService service.InovasiService
}

func NewInovasiControllerImpl(inovasiService service.InovasiService) *InovasiControllerImpl {
	return &InovasiControllerImpl{
		InovasiService: inovasiService,
	}
}

func (controller *InovasiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")
	if rekinId == "" {
		helper.WriteToResponseBody(writer, web.WebInovasiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "RekinId tidak boleh kosong",
		})
		return
	}

	// Decode request body
	inovasiCreateRequest := inovasi.InovasiCreateRequest{}
	helper.ReadFromRequestBody(request, &inovasiCreateRequest)

	// Set rekinId dari params ke request
	inovasiCreateRequest.RekinId = rekinId

	// Panggil service untuk membuat gambaran umum
	inovasiResponse, err := controller.InovasiService.Create(request.Context(), inovasiCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebInovasiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebInovasiResponse{
		Code:   http.StatusCreated,
		Status: "success create data inovasi",
		Data:   inovasiResponse,
	})
}

func (controller *InovasiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inovasiUpdateRequest := inovasi.InovasiUpdateRequest{}
	helper.ReadFromRequestBody(request, &inovasiUpdateRequest)

	inovasiId := params.ByName("id")
	inovasiUpdateRequest.Id = inovasiId

	inovasiResponse, err := controller.InovasiService.Update(request.Context(), inovasiUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebInovasiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebDasarHukumResponse{
		Code:   http.StatusOK,
		Status: "success update data inovasi",
		Data:   inovasiResponse,
	})
}

func (controller *InovasiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")

	inovasiResponses, err := controller.InovasiService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebInovasiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(writer, web.WebInovasiResponse{
		Code:   http.StatusOK,
		Status: "success get data inovasi",
		Data:   inovasiResponses,
	})
}

func (controller *InovasiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inovasiId := params.ByName("id")

	inovasiResponse, err := controller.InovasiService.FindById(request.Context(), inovasiId)
	if err != nil {
		// Periksa apakah error disebabkan oleh ID yang tidak ditemukan
		if err.Error() == "inovasi tidak ditemukan" {
			helper.WriteToResponseBody(writer, web.WebInovasiResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   "ID inovasi tidak ditemukan",
			})
		} else {
			helper.WriteToResponseBody(writer, web.WebInovasiResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   err.Error(),
			})
		}
		return
	}

	helper.WriteToResponseBody(writer, web.WebInovasiResponse{
		Code:   http.StatusOK,
		Status: "success get data inovasi",
		Data:   inovasiResponse,
	})

}

func (controller *InovasiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inovasiId := params.ByName("id")

	err := controller.InovasiService.Delete(request.Context(), inovasiId)
	if err != nil {
		// Periksa apakah error disebabkan oleh ID yang tidak ditemukan
		if err.Error() == "inovasi tidak ditemukan" {
			helper.WriteToResponseBody(writer, web.WebInovasiResponse{
				Code:   http.StatusNotFound,
				Status: "NOT FOUND",
				Data:   "ID inovasi tidak ditemukan",
			})
		} else {
			helper.WriteToResponseBody(writer, web.WebInovasiResponse{
				Code:   http.StatusInternalServerError,
				Status: "INTERNAL SERVER ERROR",
				Data:   err.Error(),
			})
		}
		return
	}

	helper.WriteToResponseBody(writer, web.WebInovasiResponse{
		Code:   http.StatusOK,
		Status: "success delete data inovasi",
		Data:   nil,
	})
}
func (controller *InovasiControllerImpl) FindAllByRekinId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekinId := params.ByName("rencana_kinerja_id")

	inovasiResponses, err := controller.InovasiService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebInovasiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   err.Error(),
		})
		return
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	buttonActions := []web.ActionButton{
		{
			NameAction: "Create Inovasi",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/inovasi/create/:rencana_kinerja_id", host, port),
		},
	}

	helper.WriteToResponseBody(writer, web.WebInovasiResponse{
		Code:   http.StatusOK,
		Status: "success get data inovasi",
		Action: buttonActions,
		Data:   inovasiResponses,
	})
}
