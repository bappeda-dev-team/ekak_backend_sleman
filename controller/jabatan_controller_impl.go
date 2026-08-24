package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/jabatan"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JabatanControllerImpl struct {
	jabatanService service.JabatanService
}

func NewJabatanControllerImpl(jabatanService service.JabatanService) *JabatanControllerImpl {
	return &JabatanControllerImpl{
		jabatanService: jabatanService,
	}
}

func (controller *JabatanControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Decode request body
	jabatanCreateRequest := jabatan.JabatanCreateRequest{}
	helper.ReadFromRequestBody(request, &jabatanCreateRequest)

	jabatanResponse, err := controller.jabatanService.Create(request.Context(), jabatanCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "failed create data jabatan",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success create data jabatan",
		Data:   jabatanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *JabatanControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jabatanUpdateRequest := jabatan.JabatanUpdateRequest{}
	helper.ReadFromRequestBody(request, &jabatanUpdateRequest)
	jabatanUpdateRequest.Id = params.ByName("id")

	jabatanResponse, err := controller.jabatanService.Update(request.Context(), jabatanUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "failed update data jabatan",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success update data jabatan",
		Data:   jabatanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *JabatanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jabatanId := params.ByName("id")
	controller.jabatanService.Delete(request.Context(), jabatanId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success delete data jabatan",
		Data:   nil,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *JabatanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jabatanId := params.ByName("id")
	jabatanResponse, err := controller.jabatanService.FindById(request.Context(), jabatanId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   404,
			Status: "not found",
			Data:   nil,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success get data jabatan by id",
		Data:   jabatanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *JabatanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	kodeOpd := params.ByName("kode_opd")
	tahun := params.ByName("tahun")

	// Validasi kode OPD tidak boleh kosong
	if kodeOpd == "" {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "failed",
			Data:   "kode OPD tidak boleh kosong",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	jabatanResponse, err := controller.jabatanService.FindAll(request.Context(), kodeOpd, tahun)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "failed get data jabatan",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "success get data jabatan",
		Data:   jabatanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
