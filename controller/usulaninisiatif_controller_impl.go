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

type UsulanInisiatifControllerImpl struct {
	UsulanInisiatifService service.UsulanInisiatifService
}

func NewUsulanInisiatifControllerImpl(usulanInisiatifService service.UsulanInisiatifService) *UsulanInisiatifControllerImpl {
	return &UsulanInisiatifControllerImpl{
		UsulanInisiatifService: usulanInisiatifService,
	}
}

func (controller *UsulanInisiatifControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanInovasiCreateRequest := usulan.UsulanInisiatifCreateRequest{}
	helper.ReadFromRequestBody(request, &usulanInovasiCreateRequest)

	pegawaiID := params.ByName("pegawai_id")
	if pegawaiID == "" {
		pegawaiID = usulanInovasiCreateRequest.PegawaiId
	}

	if pegawaiID == "" {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid pegawai_id: not found in URL params or request body",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanInovasiCreateRequest.PegawaiId = pegawaiID

	usulanInovasiResponse, err := controller.UsulanInisiatifService.Create(request.Context(), usulanInovasiCreateRequest)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:   http.StatusOK,
		Status: "success create usulan inovasi",
		Data:   usulanInovasiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanInisiatifControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanInovasiUpdateRequest := usulan.UsulanInisiatifUpdateRequest{}
	helper.ReadFromRequestBody(request, &usulanInovasiUpdateRequest)

	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanInovasiUpdateRequest.Id = idUsulan

	pegawaiID := params.ByName("pegawai_id")
	if pegawaiID != "" {
		if usulanInovasiUpdateRequest.PegawaiId != pegawaiID {
			webResponse := web.WebUsulanInisiatifResponse{
				Code:   http.StatusForbidden,
				Status: "FORBIDDEN",
				Data:   "Tidak dapat mengedit usulan pegawai lain",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	usulanInovasiResponse, err := controller.UsulanInisiatifService.Update(request.Context(), usulanInovasiUpdateRequest)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:   http.StatusOK,
		Status: "success update usulan inovasi",
		Data:   usulanInovasiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanInisiatifControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pegawaiID := params.ByName("pegawai_id")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")

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
			webResponse := web.WebUsulanInisiatifResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	usulanInovasiResponses, err := controller.UsulanInisiatifService.FindAll(request.Context(), pegawaiIDPtr, isActivePtr, rekinIDPtr)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:   http.StatusOK,
		Status: "success find all usulan inisiatif",
		Data:   usulanInovasiResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanInisiatifControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	usulanInovasiResponse, err := controller.UsulanInisiatifService.FindById(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:   http.StatusOK,
		Status: "success find usulan inovasi by id",
		Data:   usulanInovasiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanInisiatifControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err := controller.UsulanInisiatifService.Delete(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan inovasi",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanInisiatifControllerImpl) FindAllByRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pegawaiID := params.ByName("pegawai_id")
	rekinID := params.ByName("rencana_kinerja_id")
	isActive := request.URL.Query().Get("is_active")

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
			webResponse := web.WebUsulanInisiatifResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "Parameter is_active harus berupa boolean",
			}
			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		isActivePtr = &isActiveBool
	}

	usulanInisiatifResponses, err := controller.UsulanInisiatifService.FindAll(request.Context(), pegawaiIDPtr, isActivePtr, rekinIDPtr)
	if err != nil {
		webResponse := web.WebUsulanInisiatifResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	host := os.Getenv("host")
	port := os.Getenv("port")

	buttonActions := []web.ActionButton{
		{
			NameAction: "Create Usulan Inisiatif",
			Method:     "POST",
			Url:        fmt.Sprintf("%s:%s/usulan_inisiatif/create", host, port),
		},
		{
			NameAction: "Update Usulan Inisiatif",
			Method:     "PUT",
			Url:        fmt.Sprintf("%s:%s/usulan_inisiatif/update/:id", host, port),
		},
		{
			NameAction: "Delete Usulan Inisiatif",
			Method:     "DELETE",
			Url:        fmt.Sprintf("%s:%s/usulan_inisiatif/delete/:id", host, port),
		},
		{
			NameAction: "Pilihan Usulan Inisiatif",
			Method:     "GET",
			Url:        fmt.Sprintf("%s:%s/usulan_inisiatif/pilihan", host, port),
		},
		{
			NameAction:  "Create Usulan Yang Dipilih",
			Method:      "POST",
			Url:         fmt.Sprintf("%s:%s/usulan_terpilih/create", host, port),
			JenisUsulan: "inisiatif",
		},
	}

	webResponse := web.WebUsulanInisiatifResponse{
		Code:        http.StatusOK,
		Status:      "success find all usulan inisiatif",
		DataPilihan: usulanInisiatifResponses,
		Action:      buttonActions,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
