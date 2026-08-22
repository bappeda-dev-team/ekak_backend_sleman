package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/usulan"
	"ekak_kab_sleman/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UsulanPokokPikiranControllerImpl struct {
	UsulanPokokPikiranService service.UsulanPokokPikiranService
}

func NewUsulanPokokPikiranControllerImpl(usulanPokokPikiranService service.UsulanPokokPikiranService) *UsulanPokokPikiranControllerImpl {
	return &UsulanPokokPikiranControllerImpl{
		UsulanPokokPikiranService: usulanPokokPikiranService,
	}
}

func (controller *UsulanPokokPikiranControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanPokokPikiranCreateRequest := usulan.UsulanPokokPikiranCreateRequest{}
	helper.ReadFromRequestBody(request, &usulanPokokPikiranCreateRequest)

	usulanPokokPikiranResponse, err := controller.UsulanPokokPikiranService.Create(request.Context(), usulanPokokPikiranCreateRequest)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "berhasil membuat usulan pokok pikiran",
		Data:   usulanPokokPikiranResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanPokokPikiranUpdateRequest := usulan.UsulanPokokPikiranUpdateRequest{}
	helper.ReadFromRequestBody(request, &usulanPokokPikiranUpdateRequest)

	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	usulanPokokPikiranUpdateRequest.Id = idUsulan

	usulanPokokPikiranResponse, err := controller.UsulanPokokPikiranService.Update(request.Context(), usulanPokokPikiranUpdateRequest)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "berhasil mengupdate usulan pokok pikiran",
		Data:   usulanPokokPikiranResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
			webResponse := web.WebUsulanPokokPikiranResponse{
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

	usulanPokokPikiranResponses, err := controller.UsulanPokokPikiranService.FindAll(request.Context(), kodeOpdPtr, isActivePtr, rekinIDPtr, statusPtr)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "berhasil mendapatkan semua usulan pokok pikiran",
		Data:   usulanPokokPikiranResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	usulanPokokPikiranResponse, err := controller.UsulanPokokPikiranService.FindById(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "berhasil mendapatkan usulan pokok pikiran berdasarkan id",
		Data:   usulanPokokPikiranResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")
	if idUsulan == "" {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "Invalid id usulan parameter",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err := controller.UsulanPokokPikiranService.Delete(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "berhasil menghapus usulan pokok pikiran",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) FindAllByRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
			webResponse := web.WebUsulanPokokPikiranResponse{
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

	usulanPokokPikiranResponses, err := controller.UsulanPokokPikiranService.FindAll(request.Context(), kodeOpdPtr, isActivePtr, rekinIDPtr, statusPtr)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:        http.StatusBadRequest,
			Status:      "BAD REQUEST",
			DataPilihan: err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:        http.StatusOK,
		Status:      "berhasil mendapatkan semua usulan pokok pikiran",
		DataPilihan: usulanPokokPikiranResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) CreateRekin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usulanPokokPikiranCreateRekinRequest := usulan.UsulanPokokPikiranCreateRekinRequest{}
	helper.ReadFromRequestBody(request, &usulanPokokPikiranCreateRekinRequest)

	idRekin := params.ByName("rencana_kinerja_id")
	usulanPokokPikiranCreateRekinRequest.RekinId = idRekin

	usulanPokokPikiranResponse, err := controller.UsulanPokokPikiranService.CreateRekin(request.Context(), usulanPokokPikiranCreateRekinRequest)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "success create usulan pokok pikiran",
		Data:   usulanPokokPikiranResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsulanPokokPikiranControllerImpl) DeleteUsulanTerpilih(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	idUsulan := params.ByName("id")

	err := controller.UsulanPokokPikiranService.DeleteUsulanTerpilih(request.Context(), idUsulan)
	if err != nil {
		webResponse := web.WebUsulanPokokPikiranResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebUsulanPokokPikiranResponse{
		Code:   http.StatusOK,
		Status: "success delete usulan pokok pikiran terpilih",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
