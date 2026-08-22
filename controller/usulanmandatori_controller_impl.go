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

type UsulanMandatoriControllerImpl struct {
	UsulanMandatoriService service.UsulanMandatoriService
}

func NewUsulanMandatoriControllerImpl(usulanMandatoriService service.UsulanMandatoriService) *UsulanMandatoriControllerImpl {
	return &UsulanMandatoriControllerImpl{
		UsulanMandatoriService: usulanMandatoriService,
	}
}

func (controller *UsulanMandatoriControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanMandatoriCreateRequest := usulan.UsulanMandatoriCreateRequest{}
	helper.ReadFromRequestBody(request, &usulanMandatoriCreateRequest)

	pegawaiID := params.ByName("pegawai_id")
	if pegawaiID == "" {
		pegawaiID = usulanMandatoriCreateRequest.PegawaiId
	}

	if pegawaiID == "" {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid pegawai_id: not found in URL params or request body",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanMandatoriCreateRequest.PegawaiId = pegawaiID

	usulanMandatoriResponse, err := controller.UsulanMandatoriService.Create(request.Context(), usulanMandatoriCreateRequest)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:   http.StatusOK,
		Status: "success create usulan mandatori",
		Data:   usulanMandatoriResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMandatoriControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanMandatoriUpdateRequest := usulan.UsulanMandatoriUpdateRequest{}
	helper.ReadFromRequestBody(request, &usulanMandatoriUpdateRequest)

	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanMandatoriUpdateRequest.Id = idUsulan

	pegawaiID := params.ByName("pegawai_id")
	if pegawaiID != "" {
		if usulanMandatoriUpdateRequest.PegawaiId != pegawaiID {
			webResponse := web.WebUsulanMandatoriResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Tidak dapat mengedit usulan pegawai lain",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	usulanMandatoriResponse, err := controller.UsulanMandatoriService.Update(request.Context(), usulanMandatoriUpdateRequest)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:   http.StatusOK,
		Status: "success update usulan mandatori",
		Data:   usulanMandatoriResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMandatoriControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pegawaiID := params.ByName("pegawai_id")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")
	kodeOpd := request.URL.Query().Get("kode_opd")

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
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	var kodeOpdPtr *string
	if kodeOpd != "" {
		kodeOpdPtr = &kodeOpd
	}

	usulanMandatoriResponses, err := controller.UsulanMandatoriService.FindAll(request.Context(), kodeOpdPtr, pegawaiIDPtr, isActivePtr, rekinIDPtr)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:   http.StatusOK,
		Status: "success find all usulan mandatori",
		Data:   usulanMandatoriResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMandatoriControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	usulanMandatoriResponse, err := controller.UsulanMandatoriService.FindById(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:   http.StatusOK,
		Status: "success find usulan mandatori by id",
		Data:   usulanMandatoriResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMandatoriControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err := controller.UsulanMandatoriService.Delete(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan mandatori",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanMandatoriControllerImpl) FindAllByRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pegawaiID := params.ByName("pegawai_id")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")
	kodeOpd := request.URL.Query().Get("kode_opd")

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
			webResponse := web.WebUsulanMandatoriResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD REQUEST",
				DataPilihan: "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	var kodeOpdPtr *string
	if kodeOpd != "" {
		kodeOpdPtr = &kodeOpd
	}

	usulanMandatoriResponses, err := controller.UsulanMandatoriService.FindAll(request.Context(), kodeOpdPtr, pegawaiIDPtr, isActivePtr, rekinIDPtr)
	if err != nil {
		webResponse := web.WebUsulanMandatoriResponse{
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
			NameAction: "Create Usulan Mandatori",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/usulan_mandatori/create", host, port),
		},
		{
			NameAction: "Update Usulan Mandatori",
			Method:     "PUT",
			Url:        fmt.Sprintf("%s:%s/usulan_mandatori/update/:id", host, port),
		},
		{
			NameAction: "Delete Usulan Mandatori",
			Method:     "DELETE",
			Url:        fmt.Sprintf("%s:%s/usulan_mandatori/delete/:id", host, port),
		},
		{
			NameAction: "Pilihan Usulan Mandatori",
			Method:     "GET",
			Url:        fmt.Sprintf("%s:%s/usulan_mandatori/pilihan", host, port),
		},
		{
			NameAction:  "Create Usulan Yang Dipilih",
			Method:      "POST",
			Url:         fmt.Sprintf("%s:%s/usulan_terpilih/create", host, port),
			JenisUsulan: "mandatori",
		},
	}

	webResponse := web.WebUsulanMandatoriResponse{
		Code:        http.StatusOK,
		Status:      "success find all usulan mandatori",
		DataPilihan: usulanMandatoriResponses,
		Action:      buttonActions,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
