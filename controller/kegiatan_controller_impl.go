package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	"ekak_kab_sleman/model/web/kegiatan"
	"ekak_kab_sleman/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type KegiatanControllerImpl struct {
	KegiatanService service.KegiatanService
}

func NewKegiatanControllerImpl(kegiatanService service.KegiatanService) *KegiatanControllerImpl {
	return &KegiatanControllerImpl{
		KegiatanService: kegiatanService,
	}
}

func (controller *KegiatanControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Decode request JSON ke struct kegiatan
	kegiatanCreateRequest := kegiatan.KegiatanCreateRequest{}
	helper.ReadFromRequestBody(request, &kegiatanCreateRequest)

	// Panggil service untuk membuat kegiatan baru
	kegiatanResponse, err := controller.KegiatanService.Create(request.Context(), kegiatanCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Kirim response sukses
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kegiatanResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *KegiatanControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Decode request JSON ke struct kegiatan
	kegiatanUpdateRequest := kegiatan.KegiatanUpdateRequest{}
	helper.ReadFromRequestBody(request, &kegiatanUpdateRequest)

	// Ambil ID dari parameter URL
	kegiatanUpdateRequest.Id = params.ByName("id")

	// Panggil service untuk update kegiatan
	kegiatanResponse, err := controller.KegiatanService.Update(request.Context(), kegiatanUpdateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Kirim response sukses
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kegiatanResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KegiatanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil ID dari parameter URL
	kegiatanId := params.ByName("id")

	// Panggil service untuk delete kegiatan
	err := controller.KegiatanService.Delete(request.Context(), kegiatanId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Kirim response sukses
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KegiatanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil ID dari parameter URL
	kegiatanId := params.ByName("id")

	// Panggil service untuk mencari kegiatan berdasarkan ID
	kegiatanResponse, err := controller.KegiatanService.FindById(request.Context(), kegiatanId)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Kirim response sukses
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kegiatanResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KegiatanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Panggil service untuk mendapatkan semua kegiatan
	kegiatanResponses, err := controller.KegiatanService.FindAll(request.Context())
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	// Kirim response sukses
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   kegiatanResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
