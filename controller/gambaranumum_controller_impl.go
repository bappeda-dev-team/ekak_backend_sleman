package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/gambaranumum"
	"ekak_kab_sleman/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type GambaranUmumControllerImpl struct {
	GambaranUmumService service.GambaranUmumService
}

func NewGambaranUmumControllerImpl(gambaranUmumService service.GambaranUmumService) *GambaranUmumControllerImpl {
	return &GambaranUmumControllerImpl{
		GambaranUmumService: gambaranUmumService,
	}
}

func (controller *GambaranUmumControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil rekinId dari params URL
	rekinId := params.ByName("rencana_kinerja_id")
	if rekinId == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "RekinId tidak boleh kosong",
		})
		return
	}

	// Decode request body
	gambaranUmumCreateRequest := gambaranumum.GambaranUmumCreateRequest{}
	helper.ReadFromRequestBody(request, &gambaranUmumCreateRequest)

	// Set rekinId dari params ke request
	gambaranUmumCreateRequest.RekinId = rekinId

	// Panggil service untuk membuat gambaran umum
	gambaranUmumResponse, err := controller.GambaranUmumService.Create(request.Context(), gambaranUmumCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusCreated,
		Status: "success create data gambaran umum",
		Data:   gambaranUmumResponse,
	})
}

func (controller *GambaranUmumControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil id dari params URL
	id := params.ByName("id")
	if id == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID tidak boleh kosong",
		})
		return
	}

	// Decode request body
	gambaranUmumUpdateRequest := gambaranumum.GambaranUmumUpdateRequest{}
	err := json.NewDecoder(request.Body).Decode(&gambaranUmumUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Format JSON tidak valid",
		})
		return
	}

	// Set id dari params ke request
	gambaranUmumUpdateRequest.Id = id

	// Panggil service untuk update gambaran umum
	gambaranUmumResponse, err := controller.GambaranUmumService.Update(request.Context(), gambaranUmumUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusOK,
		Status: "success update data gambaran umum",
		Data:   gambaranUmumResponse,
	})
}

func (controller *GambaranUmumControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil id dari params URL
	id := params.ByName("id")
	if id == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID tidak boleh kosong",
		})
		return
	}

	// Panggil service untuk menghapus gambaran umum
	err := controller.GambaranUmumService.Delete(request.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusOK,
		Status: "success delete data gambaran umum",
		Data:   "Gambaran umum berhasil dihapus",
	})
}

func (controller *GambaranUmumControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil rekinId dari params URL
	rekinId := params.ByName("rencana_kinerja_id")
	if rekinId == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "RekinId tidak boleh kosong",
		})
		return
	}

	// Panggil service untuk mendapatkan semua gambaran umum
	gambaranUmumResponses, err := controller.GambaranUmumService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusOK,
		Status: "success get data gambaran umum",
		Data:   gambaranUmumResponses,
	})
}

func (controller *GambaranUmumControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil id dari params URL
	id := params.ByName("id")
	if id == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "ID tidak boleh kosong",
		})
		return
	}

	// Panggil service untuk mendapatkan gambaran umum berdasarkan ID
	gambaranUmumResponse, err := controller.GambaranUmumService.FindById(request.Context(), id)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	// Kirim response
	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusOK,
		Status: "success get data gambaran umum by id",
		Data:   gambaranUmumResponse,
	})
}

func (controller *GambaranUmumControllerImpl) FindAllByRekinId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil rekinId dari params URL
	rekinId := params.ByName("rencana_kinerja_id")
	if rekinId == "" {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "RekinId tidak boleh kosong",
		})
		return
	}

	// Panggil service untuk mendapatkan semua gambaran umum
	gambaranUmumResponses, err := controller.GambaranUmumService.FindAll(request.Context(), rekinId)
	if err != nil {
		helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
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
			NameAction: "Create Gambaran Umum",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/gambaran_umum/create/:rencana_kinerja_id", host, port),
		},
	}

	helper.WriteToResponseBody(writer, web.WebGambaranUmumResponse{
		Code:   http.StatusOK,
		Status: "success get data gambaran umum",
		Action: buttonActions,
		Data:   gambaranUmumResponses,
	})
}
