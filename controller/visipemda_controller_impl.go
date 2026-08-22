package controller

import (
	"ekak_kab_sleman/helper"
	"ekak_kab_sleman/model/web"
	visimisipemda "ekak_kab_sleman/model/web/visimisi"
	"ekak_kab_sleman/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type VisiPemdaControllerImpl struct {
	VisiPemdaService service.VisiPemdaService
}

func NewVisiPemdaControllerImpl(visiPemdaService service.VisiPemdaService) *VisiPemdaControllerImpl {
	return &VisiPemdaControllerImpl{VisiPemdaService: visiPemdaService}
}

func (controller *VisiPemdaControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	visiPemdaCreateRequest := visimisipemda.VisiPemdaCreateRequest{}
	helper.ReadFromRequestBody(request, &visiPemdaCreateRequest)

	visiPemdaResponse, err := controller.VisiPemdaService.Create(request.Context(), visiPemdaCreateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "success create visi pemda",
		Data:   visiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *VisiPemdaControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	visiPemdaUpdateRequest := visimisipemda.VisiPemdaUpdateRequest{}
	helper.ReadFromRequestBody(request, &visiPemdaUpdateRequest)

	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	visiPemdaUpdateRequest.Id = idInt

	visiPemdaResponse, err := controller.VisiPemdaService.Update(request.Context(), visiPemdaUpdateRequest)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success update visi pemda",
		Data:   visiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *VisiPemdaControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	err = controller.VisiPemdaService.Delete(request.Context(), idInt)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success delete visi pemda",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *VisiPemdaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	visiPemdaResponse, err := controller.VisiPemdaService.FindById(request.Context(), idInt)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success find visi pemda",
		Data:   visiPemdaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *VisiPemdaControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	tahunAwal := params.ByName("tahun_awal")
	tahunAkhir := params.ByName("tahun_akhir")
	jenisPeriode := params.ByName("jenis_periode")

	visiPemdaResponses, err := controller.VisiPemdaService.FindAll(request.Context(), tahunAwal, tahunAkhir, jenisPeriode)
	if err != nil {
		helper.WriteToResponseBody(writer, err.Error())
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "success find all visi pemda",
		Data:   visiPemdaResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
