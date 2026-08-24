package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/usulan"
	"ekak_kab_sleman/service"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UsulanMusrebangControllerImpl struct {
	UsulanMusrebangService service.UsulanMusrebangService
}

func NewUsulanMusrebangControllerImpl(usulanMusrebangService service.UsulanMusrebangService) *UsulanMusrebangControllerImpl {
	return &UsulanMusrebangControllerImpl{
		UsulanMusrebangService: usulanMusrebangService,
	}
}

func (controller *UsulanMusrebangControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanMusrebangCreateRequest := usulan.UsulanMusrebangCreateRequest{}
	helper.ReadFromRequestBody(request, &usulanMusrebangCreateRequest)

	usulanMusrebangResponse, err := controller.UsulanMusrebangService.Create(request.Context(), usulanMusrebangCreateRequest)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success create usulan musrebang",
		Data:   usulanMusrebangResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanMusrebangUpdateRequest := usulan.UsulanMusrebangUpdateRequest{}
	helper.ReadFromRequestBody(request, &usulanMusrebangUpdateRequest)

	// Ambil id usulan dari params URL
	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanMusrebangUpdateRequest.Id = idUsulan

	// Lakukan update
	usulanMusrebangResponse, err := controller.UsulanMusrebangService.Update(request.Context(), usulanMusrebangUpdateRequest)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success update usulan musrebang",
		Data:   usulanMusrebangResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOpd := params.ByName("kode_opd")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")
	status := request.URL.Query().Get("status")

	var kodeOpdPtr *string
	if kodeOpd != "" {
		kodeOpdPtr = &kodeOpd
	}

	var rekinIDPtr *string
	if rekinID != "" {
		rekinIDPtr = &rekinID
	}

	var isActivePtr *bool
	if isActive != "" {
		isActiveBool, err := strconv.ParseBool(isActive)
		if err != nil {
			webResponse := web.WebUsulanMusrebangResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	usulanMusrebangResponses, err := controller.UsulanMusrebangService.FindAll(request.Context(), kodeOpdPtr, isActivePtr, rekinIDPtr, statusPtr)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success find all usulan musrebang",
		Data:   usulanMusrebangResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	usulanMusrebangResponse, err := controller.UsulanMusrebangService.FindById(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success find usulan musrebang by id",
		Data:   usulanMusrebangResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err := controller.UsulanMusrebangService.Delete(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan musrebang",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) FindAllRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pegawaiID := params.ByName("pegawai_id")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")
	status := request.URL.Query().Get("status")

	var pegawaiIDPtr *string
	if pegawaiID != "" {
		pegawaiIDPtr = &pegawaiID
	}

	var rekinIDPtr *string
	if rekinID != "" {
		rekinIDPtr = &rekinID
	}

	var isActivePtr *bool
	if isActive != "" {
		isActiveBool, err := strconv.ParseBool(isActive)
		if err != nil {
			webResponse := web.WebUsulanMusrebangResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD REQUEST",
				DataPilihan: "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	usulanMusrebangResponses, err := controller.UsulanMusrebangService.FindAll(request.Context(), pegawaiIDPtr, isActivePtr, rekinIDPtr, statusPtr)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:        http.StatusBadRequest,
			Status:      "BAD REQUEST",
			DataPilihan: err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	host := os.Getenv("host")
	port := os.Getenv("port")

	buttonActions := []web.ActionButton{
		{
			NameAction: "Create Usulan Musrebang",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/usulan_musrebang/create/", host, port),
		},
		{
			NameAction: "Update Usulan Musrebang",
			Method:     "PUT",
			Url:        fmt.Sprintf("%s:%s/usulan_musrebang/update/:id", host, port),
		},
		{
			NameAction: "Delete Usulan Musrebang",
			Method:     "DELETE",
			Url:        fmt.Sprintf("%s:%s/usulan_musrebang/delete/:id", host, port),
		},
		{
			NameAction: "Pilihan Usulan Musrebang",
			Method:     "GET",
			Url:        fmt.Sprintf("%s:%s/usulan_musrebang/pilihan", host, port),
		},
		{
			NameAction:  "Create Usulan Yang Dipilih",
			Method:      "POST",
			Url:         fmt.Sprintf("%s:%s/usulan_terpilih/create", host, port),
			JenisUsulan: "musrebang",
		},
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:        http.StatusOK,
		Status:      "success find all usulan musrebang",
		DataPilihan: usulanMusrebangResponses,
		Action:      buttonActions,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) CreateRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanMusrebangCreateRekinRequest := usulan.UsulanMusrebangCreateRekinRequest{}
	helper.ReadFromRequestBody(request, &usulanMusrebangCreateRekinRequest)

	idRekin := params.ByName("rencana_kinerja_id")
	usulanMusrebangCreateRekinRequest.RekinId = idRekin

	usulanMusrebangResponse, err := controller.UsulanMusrebangService.CreateRekin(request.Context(), usulanMusrebangCreateRekinRequest)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success create usulan musrebang",
		Data:   usulanMusrebangResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMusrebangControllerImpl) DeleteUsulanTerpilih(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	err := controller.UsulanMusrebangService.DeleteUsulanTerpilih(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanMusrebangResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMusrebangResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan musrebang terpilih",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
